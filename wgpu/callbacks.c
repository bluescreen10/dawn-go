#include "webgpu.h"
#include "_cgo_export.h"

void cgo_callback_BufferMapCallback(WGPUMapAsyncStatus status, WGPUStringView message, void *userData1, void *userData2)
{
	goBufferMapCallbackHandler(status, message, userData1, userData2);
}

void cgo_callback_QueueWorkDoneCallback(WGPUQueueWorkDoneStatus status, WGPUStringView message, void *userData1, void *userData2)
{
	goQueueWorkDoneCallbackHandler(status, message, userData1, userData2);
}

void cgo_callback_RequestDeviceCallback(WGPURequestDeviceStatus status, WGPUDevice device, WGPUStringView message, void *userData1, void *userData2)
{
	goRequestDeviceCallbackHandler(status, device, message, userData1, userData2);
}

void cgo_callback_CompilationInfoCallback(WGPUCompilationInfoRequestStatus status, WGPUCompilationInfo compilationInfo, void *userData1, void *userData2)
{
	goCompilationInfoCallbackHandler(status, compilationInfo, userData1, userData2);
}

void cgo_callback_CreateRenderPipelineAsyncCallback(WGPUCreatePipelineAsyncStatus status, WGPURenderPipeline pipeline, WGPUStringView message, void *userData1, void *userData2)
{
	goCreateRenderPipelineAsyncCallbackHandler(status, pipeline, message, userData1, userData2);
}

void cgo_callback_CreateComputePipelineAsyncCallback(WGPUCreatePipelineAsyncStatus status, WGPUComputePipeline pipeline, WGPUStringView message, void *userData1, void *userData2)
{
	goCreateComputePipelineAsyncCallbackHandler(status, pipeline, message, userData1, userData2);
}

void cgo_callback_DeviceLostCallback(WGPUDevice device, WGPUDeviceLostReason reason, WGPUStringView message, void *userData1, void *userData2)
{
	goDeviceLostCallbackHandler(device, reason, message, userData1, userData2);
}

void cgo_callback_PopErrorScopeCallback(WGPUPopErrorScopeStatus status, WGPUErrorType typ, WGPUStringView message, void *userData1, void *userData2)
{
	goPopErrorScopeCallbackHandler(status, typ, message, userData1, userData2);
}

void cgo_callback_RequestAdapterCallback(WGPURequestAdapterStatus status, WGPUAdapter adapter, WGPUStringView message, void *userData1, void *userData2)
{
	goRequestAdapterCallbackHandler(status, adapter, message, userData1, userData2);
}

void cgo_callback_UncapturedErrorCallback(WGPUDevice device, WGPUErrorType typ, WGPUStringView message, void *userData1, void *userData2)
{
	goUncapturedErrorCallbackHandler(device, typ, message, userData1, userData2);
}
