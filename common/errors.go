package common

import "fmt"

// ErrorCode 错误码
type ErrorCode int

const (
	ErrCodeUnknown ErrorCode = iota
	ErrCodeAuth              // 认证错误
	ErrCodeNetwork           // 网络错误
	ErrCodeAPI               // API 返回错误
	ErrCodeParse             // 解析错误
	ErrCodeConfig            // 配置错误
)

// KWError 结构化错误
type KWError struct {
	Code    ErrorCode
	Message string
	Cause   error // 原始错误
}

func (e *KWError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("[%d] %s: %v", e.Code, e.Message, e.Cause)
	}
	return fmt.Sprintf("[%d] %s", e.Code, e.Message)
}

func (e *KWError) Unwrap() error {
	return e.Cause
}

// NewAuthError 创建认证错误
func NewAuthError(msg string, cause error) *KWError {
	return &KWError{Code: ErrCodeAuth, Message: msg, Cause: cause}
}

// NewNetworkError 创建网络错误
func NewNetworkError(msg string, cause error) *KWError {
	return &KWError{Code: ErrCodeNetwork, Message: msg, Cause: cause}
}

// NewAPIError 创建 API 错误
func NewAPIError(status int, message string) *KWError {
	return &KWError{
		Code:    ErrCodeAPI,
		Message: fmt.Sprintf("status: %d, message: %s", status, message),
	}
}

// NewParseError 创建解析错误
func NewParseError(msg string, cause error) *KWError {
	return &KWError{Code: ErrCodeParse, Message: msg, Cause: cause}
}

// NewConfigError 创建配置错误
func NewConfigError(msg string, cause error) *KWError {
	return &KWError{Code: ErrCodeConfig, Message: msg, Cause: cause}
}
