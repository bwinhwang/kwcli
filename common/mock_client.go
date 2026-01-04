package common

// MockKWClient 用于测试的 mock 客户端
type MockKWClient struct {
	ExecuteFunc func(data map[string]interface{}) ([]string, error)
}

// Execute 实现 KWExecutor 接口
func (m *MockKWClient) Execute(data map[string]interface{}) ([]string, error) {
	if m.ExecuteFunc != nil {
		return m.ExecuteFunc(data)
	}
	return nil, nil
}

// 编译时检查：确保 MockKWClient 实现了 KWExecutor 接口
var _ KWExecutor = (*MockKWClient)(nil)
