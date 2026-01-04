package cmd

import (
	"binhong/kwcli/common"
	"io"
	"os"
)

// App 应用上下文，封装所有依赖
type App struct {
	Client       common.KWExecutor // API 客户端接口
	Config       *AppConfig        // 应用配置
	Output       io.Writer         // 输出目标，便于测试
	DefinedFlags map[string]bool   // 标志注册表
}

// AppConfig 应用配置
type AppConfig struct {
	User       string
	Token      string
	URL        string
	OutputFile string
}

// 默认 App 单例
var defaultApp *App

// GetApp 获取默认 App 实例
func GetApp() *App {
	if defaultApp == nil {
		defaultApp = NewApp()
	}
	return defaultApp
}

// NewApp 创建新的 App 实例
func NewApp() *App {
	return &App{
		Config:       &AppConfig{},
		DefinedFlags: make(map[string]bool),
		Output:       os.Stdout,
	}
}

// NewAppWithClient 创建带有指定客户端的 App 实例（用于测试）
func NewAppWithClient(client common.KWExecutor) *App {
	return &App{
		Client:       client,
		Config:       &AppConfig{},
		DefinedFlags: make(map[string]bool),
		Output:       os.Stdout,
	}
}

// InitClient 延迟初始化客户端
func (a *App) InitClient() error {
	if a.Client != nil {
		return nil
	}

	// 使用命令行参数覆盖配置
	a.Config.User = cmdUser
	a.Config.Token = cmdToken
	a.Config.URL = cmdURL
	a.Config.OutputFile = outputFile

	url, user, token, err := LoadKWauthInfo()
	if err != nil {
		return common.NewConfigError("failed to load auth info", err)
	}

	a.Client = common.NewKWClient(url, user, token)
	return nil
}

// GetClient 获取客户端，如果未初始化则先初始化
func (a *App) GetClient() (common.KWExecutor, error) {
	if err := a.InitClient(); err != nil {
		return nil, err
	}
	return a.Client, nil
}

// RegisterFlag 注册一个命令标志
func (a *App) RegisterFlag(name string) {
	a.DefinedFlags[name] = true
}

// IsDefinedFlag 检查标志是否已注册
func (a *App) IsDefinedFlag(name string) bool {
	return a.DefinedFlags[name]
}
