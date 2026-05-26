//go:build js

package wgpu

import (
	"fmt"
	"syscall/js"
)

// sizeOrUndefined returns js.Undefined when size is WholeSize (letting the browser use the rest of the buffer),
// or a JS number otherwise.
func sizeOrUndefined(size uint64) js.Value {
	if size == WholeSize {
		return js.Undefined()
	}
	return js.ValueOf(size)
}

// awaitPromise blocks until the JS promise resolves or rejects.
func awaitPromise(promise js.Value) (js.Value, error) {
	res := make(chan js.Value, 1)
	rej := make(chan string, 1)

	thenFn := js.FuncOf(func(this js.Value, args []js.Value) any {
		var v js.Value
		if len(args) > 0 {
			v = args[0]
		}
		res <- v
		return nil
	})
	rejectFn := js.FuncOf(func(this js.Value, args []js.Value) any {
		msg := "unknown error"
		if len(args) > 0 {
			msg = args[0].String()
		}
		rej <- msg
		return nil
	})
	defer thenFn.Release()
	defer rejectFn.Release()

	promise.Call("then", thenFn).Call("catch", rejectFn)

	select {
	case v := <-res:
		return v, nil
	case msg := <-rej:
		return js.Undefined(), fmt.Errorf("%s", msg)
	}
}
