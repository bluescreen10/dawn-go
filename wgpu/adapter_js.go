//go:build js

package wgpu

import (
	"fmt"
	"syscall/js"
)

// Adapter represents a GPU adapter, which is a physical or virtual device that can be used to create WebGPU resources.
type Adapter struct {
	ref js.Value
}

// GetLimits returns the limits supported by the adapter.
// Returns the limits, panics if they cannot be retrieved.
func (a *Adapter) GetLimits() Limits {
	limits, err := a.TryGetLimits()
	if err != nil {
		panic(err)
	}
	return limits
}

// TryGetLimits returns the limits supported by the adapter, or an error if they cannot be retrieved.
func (a *Adapter) TryGetLimits() (Limits, error) {
	jsLimits := a.ref.Get("limits")
	if jsLimits.IsUndefined() || jsLimits.IsNull() {
		return Limits{}, fmt.Errorf("error getting adapter limits")
	}
	return limitsFromJS(jsLimits), nil
}

// GetInfo returns information about the adapter, such as vendor, device name, and backend type.
// Panics if the information cannot be retrieved.
func (a *Adapter) GetInfo() AdapterInfo {
	info, err := a.TryGetInfo()
	if err != nil {
		panic(err)
	}
	return info
}

// TryGetInfo returns information about the adapter, or an error if it cannot be retrieved.
func (a *Adapter) TryGetInfo() (AdapterInfo, error) {
	jsInfo := a.ref.Get("info")
	if jsInfo.IsUndefined() || jsInfo.IsNull() {
		return AdapterInfo{}, fmt.Errorf("error getting adapter info")
	}
	info := AdapterInfo{
		Vendor:       jsInfo.Get("vendor").String(),
		Architecture: jsInfo.Get("architecture").String(),
		Device:       jsInfo.Get("device").String(),
		Description:  jsInfo.Get("description").String(),
		BackendType:  BackendTypeWebGPU,
	}

	return info, nil
}

// HasFeature checks if the adapter supports the given feature.
// Returns true if the feature is supported, false otherwise.
func (a *Adapter) HasFeature(feature FeatureName) bool {
	name := feature.toJS()
	if name.IsUndefined() || name.IsNull() {
		return false
	}
	return a.ref.Get("features").Call("has", name).Bool()
}

// GetFeatures returns a list of all features supported by the adapter.
func (a *Adapter) GetFeatures() []FeatureName {
	return featuresFromJS(a.ref.Get("features"))
}

// RequestDevice requests a logical GPU device from the adapter with the given descriptor.
// Panics if the request fails.
func (a *Adapter) RequestDevice(descriptor *DeviceDescriptor) *Device {
	device, err := a.TryRequestDevice(descriptor)
	if err != nil {
		panic(err)
	}
	return device
}

// TryRequestDevice requests a logical GPU device from the adapter, returning the device and any error.
func (a *Adapter) TryRequestDevice(descriptor *DeviceDescriptor) (*Device, error) {
	desc := map[string]any{}

	if descriptor != nil {
		if descriptor.Label != "" {
			desc["label"] = descriptor.Label
		}
		if len(descriptor.RequiredFeatures) > 0 {
			features := make([]any, 0, len(descriptor.RequiredFeatures))
			for _, f := range descriptor.RequiredFeatures {
				if name := f.toJS(); !name.IsUndefined() {
					features = append(features, name)
				}
			}
			desc["requiredFeatures"] = features
		}
		if descriptor.RequiredLimits != nil {
			desc["requiredLimits"] = limitsToJS(*descriptor.RequiredLimits)
		}
		if descriptor.DefaultQueue.Label != "" {
			desc["defaultQueue"] = map[string]any{"label": descriptor.DefaultQueue.Label}
		}
	}

	promise := a.ref.Call("requestDevice", desc)
	val, err := awaitPromise(promise)
	if err != nil {
		return nil, fmt.Errorf("requestDevice: %w", err)
	}

	device := &Device{ref: val}

	if descriptor != nil {
		if descriptor.DeviceLostCallback != nil {
			cb := descriptor.DeviceLostCallback
			go func() {
				lostInfo, _ := awaitPromise(val.Get("lost"))
				reason := DeviceLostReasonUnknown
				msg := ""
				if !lostInfo.IsUndefined() && !lostInfo.IsNull() {
					msg = lostInfo.Get("message").String()
					switch lostInfo.Get("reason").String() {
					case "destroyed":
						reason = DeviceLostReasonDestroyed
					}
				}
				cb(device, reason, msg)
			}()
		}

		if descriptor.UncapturedErrorCallback != nil {
			cb := descriptor.UncapturedErrorCallback
			fn := js.FuncOf(func(this js.Value, args []js.Value) any {
				if len(args) > 0 {
					evt := args[0]
					errObj := evt.Get("error")
					typ, msg := ErrorTypeFromJS(errObj)
					go cb(device, typ, msg)
				}
				return nil
			})
			val.Call("addEventListener", "uncapturederror", fn)
		}
	}

	return device, nil
}

// Release is a no-op in JS (GC handles cleanup).
func (a *Adapter) Release() {}

// limitsFromJS reads a JS GPUSupportedLimits object into a Go Limits struct.
func limitsFromJS(jsLimits js.Value) Limits {
	get := func(name string) uint32 {
		v := jsLimits.Get(name)
		if v.IsUndefined() {
			return 0
		}
		return uint32(v.Int())
	}
	get64 := func(name string) uint64 {
		v := jsLimits.Get(name)
		if v.IsUndefined() {
			return 0
		}
		return uint64(v.Int())
	}
	return Limits{
		MaxTextureDimension1D:                     get("maxTextureDimension1D"),
		MaxTextureDimension2D:                     get("maxTextureDimension2D"),
		MaxTextureDimension3D:                     get("maxTextureDimension3D"),
		MaxTextureArrayLayers:                     get("maxTextureArrayLayers"),
		MaxBindGroups:                             get("maxBindGroups"),
		MaxBindGroupsPlusVertexBuffers:            get("maxBindGroupsPlusVertexBuffers"),
		MaxBindingsPerBindGroup:                   get("maxBindingsPerBindGroup"),
		MaxDynamicUniformBuffersPerPipelineLayout: get("maxDynamicUniformBuffersPerPipelineLayout"),
		MaxDynamicStorageBuffersPerPipelineLayout: get("maxDynamicStorageBuffersPerPipelineLayout"),
		MaxSampledTexturesPerShaderStage:          get("maxSampledTexturesPerShaderStage"),
		MaxSamplersPerShaderStage:                 get("maxSamplersPerShaderStage"),
		MaxStorageBuffersPerShaderStage:           get("maxStorageBuffersPerShaderStage"),
		MaxStorageTexturesPerShaderStage:          get("maxStorageTexturesPerShaderStage"),
		MaxUniformBuffersPerShaderStage:           get("maxUniformBuffersPerShaderStage"),
		MaxUniformBufferBindingSize:               get64("maxUniformBufferBindingSize"),
		MaxStorageBufferBindingSize:               get64("maxStorageBufferBindingSize"),
		MinUniformBufferOffsetAlignment:           get("minUniformBufferOffsetAlignment"),
		MinStorageBufferOffsetAlignment:           get("minStorageBufferOffsetAlignment"),
		MaxVertexBuffers:                          get("maxVertexBuffers"),
		MaxBufferSize:                             get64("maxBufferSize"),
		MaxVertexAttributes:                       get("maxVertexAttributes"),
		MaxVertexBufferArrayStride:                get("maxVertexBufferArrayStride"),
		MaxInterStageShaderVariables:              get("maxInterStageShaderVariables"),
		MaxColorAttachments:                       get("maxColorAttachments"),
		MaxColorAttachmentBytesPerSample:          get("maxColorAttachmentBytesPerSample"),
		MaxComputeWorkgroupStorageSize:            get("maxComputeWorkgroupStorageSize"),
		MaxComputeInvocationsPerWorkgroup:         get("maxComputeInvocationsPerWorkgroup"),
		MaxComputeWorkgroupSizeX:                  get("maxComputeWorkgroupSizeX"),
		MaxComputeWorkgroupSizeY:                  get("maxComputeWorkgroupSizeY"),
		MaxComputeWorkgroupSizeZ:                  get("maxComputeWorkgroupSizeZ"),
		MaxComputeWorkgroupsPerDimension:          get("maxComputeWorkgroupsPerDimension"),
	}
}

// limitsToJS converts a Go Limits struct into a JS-compatible map.
func limitsToJS(l Limits) map[string]any {
	return map[string]any{
		"maxTextureDimension1D":                     l.MaxTextureDimension1D,
		"maxTextureDimension2D":                     l.MaxTextureDimension2D,
		"maxTextureDimension3D":                     l.MaxTextureDimension3D,
		"maxTextureArrayLayers":                     l.MaxTextureArrayLayers,
		"maxBindGroups":                             l.MaxBindGroups,
		"maxBindGroupsPlusVertexBuffers":            l.MaxBindGroupsPlusVertexBuffers,
		"maxBindingsPerBindGroup":                   l.MaxBindingsPerBindGroup,
		"maxDynamicUniformBuffersPerPipelineLayout": l.MaxDynamicUniformBuffersPerPipelineLayout,
		"maxDynamicStorageBuffersPerPipelineLayout": l.MaxDynamicStorageBuffersPerPipelineLayout,
		"maxSampledTexturesPerShaderStage":          l.MaxSampledTexturesPerShaderStage,
		"maxSamplersPerShaderStage":                 l.MaxSamplersPerShaderStage,
		"maxStorageBuffersPerShaderStage":           l.MaxStorageBuffersPerShaderStage,
		"maxStorageTexturesPerShaderStage":          l.MaxStorageTexturesPerShaderStage,
		"maxUniformBuffersPerShaderStage":           l.MaxUniformBuffersPerShaderStage,
		"maxUniformBufferBindingSize":               l.MaxUniformBufferBindingSize,
		"maxStorageBufferBindingSize":               l.MaxStorageBufferBindingSize,
		"minUniformBufferOffsetAlignment":           l.MinUniformBufferOffsetAlignment,
		"minStorageBufferOffsetAlignment":           l.MinStorageBufferOffsetAlignment,
		"maxVertexBuffers":                          l.MaxVertexBuffers,
		"maxBufferSize":                             l.MaxBufferSize,
		"maxVertexAttributes":                       l.MaxVertexAttributes,
		"maxVertexBufferArrayStride":                l.MaxVertexBufferArrayStride,
		"maxInterStageShaderVariables":              l.MaxInterStageShaderVariables,
		"maxColorAttachments":                       l.MaxColorAttachments,
		"maxColorAttachmentBytesPerSample":          l.MaxColorAttachmentBytesPerSample,
		"maxComputeWorkgroupStorageSize":            l.MaxComputeWorkgroupStorageSize,
		"maxComputeInvocationsPerWorkgroup":         l.MaxComputeInvocationsPerWorkgroup,
		"maxComputeWorkgroupSizeX":                  l.MaxComputeWorkgroupSizeX,
		"maxComputeWorkgroupSizeY":                  l.MaxComputeWorkgroupSizeY,
		"maxComputeWorkgroupSizeZ":                  l.MaxComputeWorkgroupSizeZ,
		"maxComputeWorkgroupsPerDimension":          l.MaxComputeWorkgroupsPerDimension,
	}
}

// featuresFromJS reads a JS GPUSupportedFeatures Set into a Go []FeatureName.
func featuresFromJS(jsFeatures js.Value) []FeatureName {
	if jsFeatures.IsUndefined() || jsFeatures.IsNull() {
		return nil
	}
	var result []FeatureName
	forEach := js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) > 0 {
			if f, ok := featureNameFromJS(args[0].String()); ok {
				result = append(result, f)
			}
		}
		return nil
	})
	defer forEach.Release()
	jsFeatures.Call("forEach", forEach)
	return result
}
