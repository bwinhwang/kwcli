package common

// KWExecutor 定义 Klocwork API 客户端接口
// 便于测试时使用 mock 实现
type KWExecutor interface {
	Execute(data map[string]interface{}) ([]string, error)
}

// KWClientConfig 客户端配置
type KWClientConfig struct {
	BaseURL string
	User    string
	Token   string
}

// 编译时检查：确保 KWClient 实现了 KWExecutor 接口
var _ KWExecutor = (*KWClient)(nil)
