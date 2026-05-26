//go:build js

package wgpu

import (
	"fmt"
	"syscall/js"
	"time"
)

// Instance is a WebGPU instance, which serves as the entry point for all WebGPU operations.
// It manages the underlying GPU backend and is used to request adapters.
// In JS/WASM this wraps navigator.gpu.
type Instance struct {
	ref js.Value
}

// CreateSurface creates a surface for rendering to a window.
// In JS/WASM, the surface is created from the canvas element identified by CanvasID in the descriptor.
func (i *Instance) CreateSurface(descriptor SurfaceDescriptor) *Surface {
	canvas := js.Global().Get("document").Call("getElementById", descriptor.CanvasID)
	ctx := canvas.Call("getContext", "webgpu")
	return &Surface{ref: ctx, canvas: canvas}
}

// RequestAdapter requests a GPU adapter from the instance based on the provided options.
// Returns the adapter and an error if the request fails.
func (i *Instance) RequestAdapter(options *RequestAdapterOptions) (*Adapter, error) {
	opts := map[string]any{}
	if options != nil {
		if pp := options.PowerPreference.toJS(); pp != "" {
			opts["powerPreference"] = pp
		}
		if options.ForceFallbackAdapter {
			opts["forceFallbackAdapter"] = true
		}
	}

	promise := i.ref.Call("requestAdapter", opts)
	val, err := awaitPromise(promise)
	if err != nil {
		return nil, fmt.Errorf("requestAdapter: %w", err)
	}
	if val.IsNull() || val.IsUndefined() {
		return nil, fmt.Errorf("requestAdapter: no suitable adapter found")
	}
	return &Adapter{ref: val}, nil
}

// HasWGSLLanguageFeature checks if the instance supports the given WGSL language feature.
// Returns true if the feature is supported, false otherwise.
func (i *Instance) HasWGSLLanguageFeature(feature WGSLLanguageFeatureName) bool {
	features := i.ref.Get("wgslLanguageFeatures")
	if features.IsUndefined() || features.IsNull() {
		return false
	}
	// WGSLLanguageFeatureName is a uint32; the JS API uses string names.
	// We don't have a full mapping here; return false for now.
	_ = feature
	return false
}

// GetWGSLLanguageFeatures returns all WGSL language features supported by the instance.
func (i *Instance) GetWGSLLanguageFeatures() []WGSLLanguageFeatureName {
	return nil
}

// ProcessEvents processes any pending events in the instance, such as callbacks and device lost notifications.
// In JS, this is a no-op as the event loop handles this.
func (i *Instance) ProcessEvents() {}

// Release releases the instance and all associated resources.
// In JS, this is a no-op as GC handles cleanup.
func (i *Instance) Release() {}

// Wait blocks the current thread until the specified future completes or the timeout expires.
// Returns an error if the wait fails or times out.
func (i *Instance) Wait(future Future, timeout time.Duration) error {
	promise := js.Value(future)
	if promise.IsUndefined() || promise.IsNull() {
		return nil
	}
	done := make(chan error, 1)
	go func() {
		_, err := awaitPromise(promise)
		done <- err
	}()
	select {
	case err := <-done:
		return err
	case <-time.After(timeout):
		return fmt.Errorf("wgpu: Wait timed out after %v", timeout)
	}
}

// WaitAny blocks the current thread until any of the specified futures completes or the timeout expires.
// In JS, futures resolve through the JS event loop so this is a no-op.
func (i *Instance) WaitAny(futures []Future, timeout time.Duration) error {
	return nil
}

// CreateInstance creates a new WebGPU instance with optional descriptor.
// The instance is the entry point for all WebGPU operations.
// In JS/WASM this wraps navigator.gpu.
func CreateInstance(descriptor *InstanceDescriptor) *Instance {
	gpu := js.Global().Get("navigator").Get("gpu")
	return &Instance{ref: gpu}
}

// GetInstanceFeatures returns a list of all instance-level features supported by the WebGPU implementation.
func GetInstanceFeatures() []InstanceFeatureName {
	return nil
}

// GetInstanceLimits returns the instance-level limits supported by the WebGPU implementation.
// Returns the limits and an error if they cannot be retrieved.
func GetInstanceLimits() (InstanceLimits, error) {
	return InstanceLimits{}, nil
}

// HasInstanceFeature checks if the given instance feature is supported by the WebGPU implementation.
// Returns true if the feature is supported, false otherwise.
func HasInstanceFeature(feature InstanceFeatureName) bool {
	return false
}
