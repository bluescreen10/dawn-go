// CODE GENERATED. DO NOT EDIT
//go:generate go run ./cmd/wrapper/. 
package wgpu

type RenderPassEncoder struct{
	ref uintptr
}

func (r *RenderPassEncoder) SetPipeline(pipeline RenderPipeline) 
func (r *RenderPassEncoder) SetBindGroup(groupIndex uint32, group BindGroup, dynamicOffsetCount int, dynamicOffsets uint32) 
func (r *RenderPassEncoder) Draw(vertexCount uint32, instanceCount uint32, firstVertex uint32, firstInstance uint32) 
func (r *RenderPassEncoder) DrawIndexed(indexCount uint32, instanceCount uint32, firstIndex uint32, baseVertex int32, firstInstance uint32) 
func (r *RenderPassEncoder) DrawIndirect(indirectBuffer Buffer, indirectOffset uint64) 
func (r *RenderPassEncoder) DrawIndexedIndirect(indirectBuffer Buffer, indirectOffset uint64) 
func (r *RenderPassEncoder) ExecuteBundles(bundleCount int, bundles RenderBundle) 
func (r *RenderPassEncoder) InsertDebugMarker(markerLabel string) 
func (r *RenderPassEncoder) PopDebugGroup() 
func (r *RenderPassEncoder) PushDebugGroup(groupLabel string) 
func (r *RenderPassEncoder) SetStencilReference(reference uint32) 
func (r *RenderPassEncoder) SetBlendConstant(color Color) 
func (r *RenderPassEncoder) SetViewport(x float32, y float32, width float32, height float32, minDepth float32, maxDepth float32) 
func (r *RenderPassEncoder) SetScissorRect(x uint32, y uint32, width uint32, height uint32) 
func (r *RenderPassEncoder) SetVertexBuffer(slot uint32, buffer Buffer, offset uint64, size uint64) 
func (r *RenderPassEncoder) SetIndexBuffer(buffer Buffer, format IndexFormat, offset uint64, size uint64) 
func (r *RenderPassEncoder) BeginOcclusionQuery(queryIndex uint32) 
func (r *RenderPassEncoder) EndOcclusionQuery() 
func (r *RenderPassEncoder) End() 
func (r *RenderPassEncoder) SetLabel(label string) 


type CommandEncoder struct{
	ref uintptr
}

func (c *CommandEncoder) Finish(descriptor CommandBufferDescriptor) CommandBuffer
func (c *CommandEncoder) BeginComputePass(descriptor ComputePassDescriptor) ComputePassEncoder
func (c *CommandEncoder) BeginRenderPass(descriptor RenderPassDescriptor) RenderPassEncoder
func (c *CommandEncoder) CopyBufferToBuffer(source Buffer, sourceOffset uint64, destination Buffer, destinationOffset uint64, size uint64) 
func (c *CommandEncoder) CopyBufferToTexture(source TexelCopyBufferInfo, destination TexelCopyTextureInfo, copySize Extent3D) 
func (c *CommandEncoder) CopyTextureToBuffer(source TexelCopyTextureInfo, destination TexelCopyBufferInfo, copySize Extent3D) 
func (c *CommandEncoder) CopyTextureToTexture(source TexelCopyTextureInfo, destination TexelCopyTextureInfo, copySize Extent3D) 
func (c *CommandEncoder) ClearBuffer(buffer Buffer, offset uint64, size uint64) 
func (c *CommandEncoder) InsertDebugMarker(markerLabel string) 
func (c *CommandEncoder) PopDebugGroup() 
func (c *CommandEncoder) PushDebugGroup(groupLabel string) 
func (c *CommandEncoder) ResolveQuerySet(querySet QuerySet, firstQuery uint32, queryCount uint32, destination Buffer, destinationOffset uint64) 
func (c *CommandEncoder) WriteTimestamp(querySet QuerySet, queryIndex uint32) 
func (c *CommandEncoder) SetLabel(label string) 


type ShaderModule struct{
	ref uintptr
}

func (s *ShaderModule) GetCompilationInfo(callbackInfo CompilationInfoCallbackInfo) Future
func (s *ShaderModule) SetLabel(label string) 


type TextureView struct{
	ref uintptr
}

func (t *TextureView) SetLabel(label string) 


type ExternalTexture struct{
	ref uintptr
}

func (e *ExternalTexture) SetLabel(label string) 


type Texture struct{
	ref uintptr
}

func (t *Texture) CreateView(descriptor TextureViewDescriptor) TextureView
func (t *Texture) SetLabel(label string) 
func (t *Texture) GetWidth() uint32
func (t *Texture) GetHeight() uint32
func (t *Texture) GetDepthOrArrayLayers() uint32
func (t *Texture) GetMipLevelCount() uint32
func (t *Texture) GetSampleCount() uint32
func (t *Texture) GetDimension() TextureDimension
func (t *Texture) GetFormat() TextureFormat
func (t *Texture) GetUsage() TextureUsage
func (t *Texture) GetTextureBindingViewDimension() TextureViewDimension
func (t *Texture) Destroy() 


type BindGroupLayout struct{
	ref uintptr
}

func (b *BindGroupLayout) SetLabel(label string) 


type Device struct{
	ref uintptr
}

func (d *Device) CreateBindGroup(descriptor BindGroupDescriptor) BindGroup
func (d *Device) CreateBindGroupLayout(descriptor BindGroupLayoutDescriptor) BindGroupLayout
func (d *Device) CreateBuffer(descriptor BufferDescriptor) (*Buffer, error)
func (d *Device) CreateCommandEncoder(descriptor CommandEncoderDescriptor) CommandEncoder
func (d *Device) CreateComputePipeline(descriptor ComputePipelineDescriptor) ComputePipeline
func (d *Device) CreateComputePipelineAsync(descriptor ComputePipelineDescriptor, callbackInfo CreateComputePipelineAsyncCallbackInfo) Future
func (d *Device) CreatePipelineLayout(descriptor PipelineLayoutDescriptor) PipelineLayout
func (d *Device) CreateQuerySet(descriptor QuerySetDescriptor) QuerySet
func (d *Device) CreateRenderPipelineAsync(descriptor RenderPipelineDescriptor, callbackInfo CreateRenderPipelineAsyncCallbackInfo) Future
func (d *Device) CreateRenderBundleEncoder(descriptor RenderBundleEncoderDescriptor) RenderBundleEncoder
func (d *Device) CreateRenderPipeline(descriptor RenderPipelineDescriptor) RenderPipeline
func (d *Device) CreateSampler(descriptor SamplerDescriptor) Sampler
func (d *Device) CreateShaderModule(descriptor ShaderModuleDescriptor) ShaderModule
func (d *Device) CreateTexture(descriptor TextureDescriptor) Texture
func (d *Device) Destroy() 
func (d *Device) GetLimits(limits Limits) Status
func (d *Device) GetLostFuture() Future
func (d *Device) HasFeature(feature FeatureName) bool
func (d *Device) GetFeatures(features SupportedFeatures) 
func (d *Device) GetAdapterInfo(adapterInfo AdapterInfo) Status
func (d *Device) GetQueue() Queue
func (d *Device) PushErrorScope(filter ErrorFilter) 
func (d *Device) PopErrorScope(callbackInfo PopErrorScopeCallbackInfo) Future
func (d *Device) SetLabel(label string) 


type ComputePipeline struct{
	ref uintptr
}

func (c *ComputePipeline) GetBindGroupLayout(groupIndex uint32) BindGroupLayout
func (c *ComputePipeline) SetLabel(label string) 


type ComputePassEncoder struct{
	ref uintptr
}

func (c *ComputePassEncoder) InsertDebugMarker(markerLabel string) 
func (c *ComputePassEncoder) PopDebugGroup() 
func (c *ComputePassEncoder) PushDebugGroup(groupLabel string) 
func (c *ComputePassEncoder) SetPipeline(pipeline ComputePipeline) 
func (c *ComputePassEncoder) SetBindGroup(groupIndex uint32, group BindGroup, dynamicOffsetCount int, dynamicOffsets uint32) 
func (c *ComputePassEncoder) DispatchWorkgroups(workgroupCountX uint32, workgroupCountY uint32, workgroupCountZ uint32) 
func (c *ComputePassEncoder) DispatchWorkgroupsIndirect(indirectBuffer Buffer, indirectOffset uint64) 
func (c *ComputePassEncoder) End() 
func (c *ComputePassEncoder) SetLabel(label string) 


type Surface struct{
	ref uintptr
}

func (s *Surface) Configure(config SurfaceConfiguration) 
func (s *Surface) GetCapabilities(adapter Adapter, capabilities SurfaceCapabilities) Status
func (s *Surface) GetCurrentTexture(surfaceTexture SurfaceTexture) 
func (s *Surface) Present() Status
func (s *Surface) Unconfigure() 
func (s *Surface) SetLabel(label string) 


type Instance struct{
	ref uintptr
}

func (i *Instance) CreateSurface(descriptor SurfaceDescriptor) Surface
func (i *Instance) ProcessEvents() 
func (i *Instance) WaitAny(futureCount int, futures FutureWaitInfo, timeoutNS uint64) WaitStatus
func (i *Instance) RequestAdapter(options RequestAdapterOptions, callbackInfo RequestAdapterCallbackInfo) Future
func (i *Instance) HasWGSLLanguageFeature(feature WGSLLanguageFeatureName) bool
func (i *Instance) GetWGSLLanguageFeatures(features SupportedWGSLLanguageFeatures) 


type Queue struct{
	ref uintptr
}

func (q *Queue) Submit(commandCount int, commands CommandBuffer) 
func (q *Queue) OnSubmittedWorkDone(callbackInfo QueueWorkDoneCallbackInfo) Future
func (q *Queue) WriteBuffer(buffer Buffer, bufferOffset uint64, data []byte, size int) 
func (q *Queue) WriteTexture(destination TexelCopyTextureInfo, data []byte, dataSize int, dataLayout TexelCopyBufferLayout, writeSize Extent3D) 
func (q *Queue) SetLabel(label string) 


type RenderPipeline struct{
	ref uintptr
}

func (r *RenderPipeline) GetBindGroupLayout(groupIndex uint32) BindGroupLayout
func (r *RenderPipeline) SetLabel(label string) 


type Buffer struct{
	ref uintptr
}

func (b *Buffer) MapAsync(mode MapMode, offset int, size int, callbackInfo BufferMapCallbackInfo) Future
func (b *Buffer) GetMappedRange(offset int, size int) uintptr
func (b *Buffer) GetConstMappedRange(offset int, size int) []byte
func (b *Buffer) WriteMappedRange(offset int, data []byte, size int) Status
func (b *Buffer) ReadMappedRange(offset int, data []byte, size int) Status
func (b *Buffer) SetLabel(label string) 
func (b *Buffer) GetUsage() BufferUsage
func (b *Buffer) GetSize() uint64
func (b *Buffer) GetMapState() BufferMapState
func (b *Buffer) Unmap() 
func (b *Buffer) Destroy() 


type Adapter struct{
	ref uintptr
}

func (a *Adapter) GetLimits(limits Limits) Status
func (a *Adapter) GetInfo(info AdapterInfo) Status
func (a *Adapter) HasFeature(feature FeatureName) bool
func (a *Adapter) GetFeatures(features SupportedFeatures) 
func (a *Adapter) RequestDevice(descriptor DeviceDescriptor, callbackInfo RequestDeviceCallbackInfo) Future


type RenderBundleEncoder struct{
	ref uintptr
}

func (r *RenderBundleEncoder) SetPipeline(pipeline RenderPipeline) 
func (r *RenderBundleEncoder) SetBindGroup(groupIndex uint32, group BindGroup, dynamicOffsetCount int, dynamicOffsets uint32) 
func (r *RenderBundleEncoder) Draw(vertexCount uint32, instanceCount uint32, firstVertex uint32, firstInstance uint32) 
func (r *RenderBundleEncoder) DrawIndexed(indexCount uint32, instanceCount uint32, firstIndex uint32, baseVertex int32, firstInstance uint32) 
func (r *RenderBundleEncoder) DrawIndirect(indirectBuffer Buffer, indirectOffset uint64) 
func (r *RenderBundleEncoder) DrawIndexedIndirect(indirectBuffer Buffer, indirectOffset uint64) 
func (r *RenderBundleEncoder) InsertDebugMarker(markerLabel string) 
func (r *RenderBundleEncoder) PopDebugGroup() 
func (r *RenderBundleEncoder) PushDebugGroup(groupLabel string) 
func (r *RenderBundleEncoder) SetVertexBuffer(slot uint32, buffer Buffer, offset uint64, size uint64) 
func (r *RenderBundleEncoder) SetIndexBuffer(buffer Buffer, format IndexFormat, offset uint64, size uint64) 
func (r *RenderBundleEncoder) Finish(descriptor RenderBundleDescriptor) RenderBundle
func (r *RenderBundleEncoder) SetLabel(label string) 


type RenderBundle struct{
	ref uintptr
}

func (r *RenderBundle) SetLabel(label string) 


type BindGroup struct{
	ref uintptr
}

func (b *BindGroup) SetLabel(label string) 


type QuerySet struct{
	ref uintptr
}

func (q *QuerySet) SetLabel(label string) 
func (q *QuerySet) GetType() QueryType
func (q *QuerySet) GetCount() uint32
func (q *QuerySet) Destroy() 


type Sampler struct{
	ref uintptr
}

func (s *Sampler) SetLabel(label string) 


type PipelineLayout struct{
	ref uintptr
}

func (p *PipelineLayout) SetLabel(label string) 


type CommandBuffer struct{
	ref uintptr
}

func (c *CommandBuffer) SetLabel(label string) 


