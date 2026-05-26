//go:build js

package wgpu

import "syscall/js"

// Future is a JS Promise that resolves once the async operation and its
// callback have both completed.
type Future js.Value

// newCallbackPromise returns a Promise that resolves after fn() returns.
// The Promise executor is called synchronously by the JS engine, so resolve
// is captured before we spin up the goroutine; the FuncOf can be released
// immediately after construction.
func newCallbackPromise(fn func()) Future {
	var resolve js.Value
	executor := js.FuncOf(func(_ js.Value, args []js.Value) any {
		resolve = args[0]
		return nil
	})
	p := js.Global().Get("Promise").New(executor)
	executor.Release()
	go func() { fn(); resolve.Invoke(js.Undefined()) }()
	return Future(p)
}
