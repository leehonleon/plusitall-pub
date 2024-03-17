package interfaces

import "fmt"

type Plugin interface {
	Init(config map[string]interface{}) error // 初始化方法，接收配置信息
	Run() error                               // 执行插件功能的方法
	GetName() string                          // 获取插件名称
}

// RegisterPlugin 注册插件到全局注册表
func RegisterPlugin(plugin Plugin) {
	// 在这里将插件添加到全局插件列表
	fmt.Printf("Registering plugin: %s\n", plugin.GetName())
}
