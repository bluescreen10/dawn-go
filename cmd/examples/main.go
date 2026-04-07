package main

import (
	"fmt"

	wgpu "github.com/bluescreen10/dawn-go"
)

func main() {
	instance := wgpu.CreateInstance(nil)
	future := instance.RequestAdapter(nil, wgpu.RequestAdapterCallbackInfo{
		Mode: wgpu.CallbackModeWaitAnyOnly,
		Callback: func(status wgpu.RequestAdapterStatus, adapter *wgpu.Adapter, message string) {
			fmt.Println(status)
		},
	})

	instance.WaitAny([]wgpu.FutureWaitInfo{{Future: future}}, 0)
	fmt.Println("bye")
	// var wg sync.WaitGroup
	// wg.Add(1)
	// wg.Wait()
}
