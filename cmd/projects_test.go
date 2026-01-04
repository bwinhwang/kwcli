package cmd

import (
	"binhong/kwcli/common"
	"io"
	"os"
	"strings"
	"testing"
)

func Test_projectsCommand(t *testing.T) {
	// 设置 mock 客户端
	mockClient := &common.MockKWClient{
		ExecuteFunc: func(data map[string]interface{}) ([]string, error) {
			return []string{
				`{"id": "proj1", "name": "TestProject"}`,
			}, nil
		},
	}
	SetKWClientInstance(mockClient)
	defer ResetKWClientInstance()

	// 保存原始 outputFile 并设置为 stdout
	origOutputFile := outputFile
	outputFile = "-"
	defer func() { outputFile = origOutputFile }()

	testCases := []struct {
		name           string
		args           []string
		expectedError  bool
		expectedOutput string // For basic output checks
	}{
		{"no project flag", nil, false, "TestProject"},
		{"include streams flag", []string{"--include_streams"}, false, "TestProject"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Capture the output
			old := os.Stdout // Keep the original stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			// 通过 rootCmd 执行 projects 子命令
			args := append([]string{"projects"}, tc.args...)
			rootCmd.SetArgs(args)
			err := rootCmd.Execute()

			w.Close()
			out, _ := io.ReadAll(r)
			os.Stdout = old // Restore stdout

			if (err != nil) != tc.expectedError {
				t.Errorf("Expected error: %v, got: %v", tc.expectedError, err)
			}

			// Basic output check
			if !strings.Contains(string(out), tc.expectedOutput) {
				t.Errorf("Expected output to contain: %s, got: %s", tc.expectedOutput, out)
			}
		})
	}
}
