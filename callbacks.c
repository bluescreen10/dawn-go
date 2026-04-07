#include <webgpu/webgpu.h>

// This is the Go function we are calling. 
// Go generates this header automatically when you use //export
#include "_cgo_export.h"

void c_callback_trampoline(WGPURequestAdapterStatus status, 
                           WGPUAdapter adapter, 
                           WGPUStringView message, 
                           void * userdata1,
                           void * userdata2) { // This 'userdata' is the value from 'userdata1'
                           
    // Forward the value exactly as received
    goCallbackHandler(status, adapter, message, userdata1, userdata2);
}