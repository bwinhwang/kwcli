package cmd

import (
	"binhong/kwcli/common"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// CommandContext 命令执行上下文
type CommandContext struct {
	App *App
	Cmd *cobra.Command
}

// NewCommandContext 创建命令上下文
func NewCommandContext(cmd *cobra.Command) *CommandContext {
	return &CommandContext{
		App: GetApp(),
		Cmd: cmd,
	}
}

// FetchData 获取数据的通用方法
func (ctx *CommandContext) FetchData(action string, resultStruct interface{}) ([]interface{}, error) {
	client, err := ctx.App.GetClient()
	if err != nil {
		return nil, err
	}

	paramMap := ctx.CollectOptionsAsMap()
	paramMap["action"] = action

	lines, err := client.Execute(paramMap)
	if err != nil {
		return nil, err
	}

	return ctx.parseResults(lines, resultStruct)
}

// ExecuteAction 执行操作的通用方法
func (ctx *CommandContext) ExecuteAction(action string) error {
	client, err := ctx.App.GetClient()
	if err != nil {
		return err
	}

	paramMap := ctx.CollectOptionsAsMap()
	paramMap["action"] = action

	_, err = client.Execute(paramMap)
	return err
}

// CollectOptionsAsMap 收集命令行参数为 map
func (ctx *CommandContext) CollectOptionsAsMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	ctx.Cmd.Flags().VisitAll(func(flag *pflag.Flag) {
		if ctx.App.IsDefinedFlag(flag.Name) && ctx.Cmd.Flags().Changed(flag.Name) {
			switch flag.Value.Type() {
			case "string":
				paramMap[flag.Name] = flag.Value.String()
			case "stringSlice":
				stringValues, _ := ctx.Cmd.Flags().GetStringSlice(flag.Name)
				paramMap[flag.Name] = strings.Join(stringValues, ",")
			default:
				fmt.Printf("Unsupported type: %s, %s\n", flag.Value.Type(), flag.Name)
			}
		}
	})
	return paramMap
}

// parseResults 解析 API 返回的 JSON 行
func (ctx *CommandContext) parseResults(lines []string, resultStruct interface{}) ([]interface{}, error) {
	resultType := reflect.TypeOf(resultStruct)
	if resultType.Kind() != reflect.Ptr || resultType.Elem().Kind() != reflect.Struct {
		return nil, common.NewParseError("resultStruct must be a pointer to a struct", nil)
	}

	var results []interface{}
	for _, line := range lines {
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}

		resultValue := reflect.New(resultType.Elem())
		if err := json.Unmarshal([]byte(line), resultValue.Interface()); err != nil {
			return nil, common.NewParseError(fmt.Sprintf("error parsing JSON: %s", line), err)
		}
		results = append(results, resultValue.Elem().Interface())
	}

	return results, nil
}
