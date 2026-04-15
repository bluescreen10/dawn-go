module example

go 1.25.0

replace github.com/bluescreen10/dawn-go => ../../

replace github.com/bluescreen10/dawn-go/examples/ => ../

require github.com/bluescreen10/dawn-go v0.0.4

require github.com/go-gl/glfw/v3.4/glfw v0.1.0-pre.1

require (
	bluescreen10.com/wgpu-android v0.0.5 // indirect
	bluescreen10.com/wgpu-darwin v0.0.5 // indirect
	bluescreen10.com/wgpu-linux v0.0.5 // indirect
	bluescreen10.com/wgpu-windows v0.0.5 // indirect
)
