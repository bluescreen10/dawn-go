//go:build js

package wgpu

import (
	"fmt"
	"syscall/js"
)

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
