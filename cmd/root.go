package cmd

import (
	"binhong/kwcli/common"
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// ... (Your other imports)
var definedFlags = make(map[string]bool)

var outputFile string
var cmdUser string
var cmdToken string
var cmdURL string

var GlobalKWClient *common.KWClient // Declare the global variable
func fetchDataCommand(cmd *cobra.Command, action string, resultStruct interface{}) ([]interface{}, error) {
	paramMap := make(map[string]interface{})
	cmd.Flags().VisitAll(func(flag *pflag.Flag) {
		if definedFlags[flag.Name] {
			paramMap[flag.Name] = flag.Value.String()
		}
	})
	paramMap["action"] = action

	client := getKWClientInstance()
	lines, err := client.Execute(paramMap)
	if err != nil {
		return nil, err
	}

	// Ensure resultStruct is a pointer to a struct
	resultType := reflect.TypeOf(resultStruct)
	if resultType.Kind() != reflect.Ptr || resultType.Elem().Kind() != reflect.Struct {
		return nil, fmt.Errorf("resultStruct must be a pointer to a struct")
	}
	var results []interface{}
	// Iterate through results and unmarshal
	for _, line := range lines {
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}

		// Create a new instance of the struct type
		resultValue := reflect.New(resultType.Elem())

		err := json.Unmarshal([]byte(line), resultValue.Interface())
		if err != nil {
			return nil, fmt.Errorf("error parsing JSON: %v - line: %s", err, line)
		}

		// Append the unmarshalled data (You'll need to adjust how the data is stored)
		results = append(results, resultValue.Elem().Interface())
	}

	return results, nil
}
func actionOrientedCommand(cmd *cobra.Command, action string) error {
	paramMap := make(map[string]interface{})
	cmd.Flags().VisitAll(func(flag *pflag.Flag) {
		if definedFlags[flag.Name] {
			paramMap[flag.Name] = flag.Value.String()
			if cmd.Flags().Changed(flag.Name) { // applied flag specfied by cmdline
				paramMap[flag.Name] = flag.Value.String()
			}
		}
	})
	paramMap["action"] = action

	client := getKWClientInstance()
	_, err := client.Execute(paramMap)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func writeJSONToFile(data interface{}, filename string) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("error marshalling JSON: %v", err)
	}

	err = os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("error writing to file: %v", err)
	}

	return nil
}

func parseURL(rawURL string) (string, string, string, error) {
	parts := strings.SplitN(rawURL, "://", 2)
	if len(parts) != 2 {
		return "", "", "", fmt.Errorf("invalid URL format: missing schema")
	}

	schema := parts[0]
	hostPortPath := parts[1]

	// Further split to separate host:port from the potential path
	hostPort, _, _ := strings.Cut(hostPortPath, "/")

	hostname, port, found := strings.Cut(hostPort, ":")
	if !found {
		port = ""
	}

	return schema, hostname, port, nil
}

func LoadKWauthInfo() (string, string, string, error) {

	var url, username, token string
	usr, err := user.Current()
	if err != nil {
		fmt.Println("Error getting home directory:", err)
		return url, username, token, err
	}
	homeDir := usr.HomeDir
	filePath := filepath.Join(homeDir, ".klocwork", "ltoken")
	//fmt.Println("klocwork auth file:", filePath)

	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error loading auth file:", err)
		return url, username, token, err
	}

	//fmt.Println(data)
	scanner := bufio.NewScanner(bytes.NewReader(data))
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ";")
		if len(parts) != 4 {
			fmt.Println("Invalid format in line:", line)
			continue
		}

		url = "http://" + parts[0] + ":" + parts[1] + "/review/api"
		//fmt.Println(url)
		username = parts[2]
		token = parts[3]
		break
	}

	if cmdUser != "" {
		username = cmdUser
	}
	if cmdToken != "" {
		token = cmdToken
	}
	if cmdURL != "" {

		schema, hostname, port, err := parseURL(cmdURL)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		} else {
			url = schema + "://" + hostname + ":" + port + "/review/api"
		}

	}
	//fmt.Printf("user: %s, token: %s, url: %s\n", username, token, url)
	return url, username, token, err
}

// The getKWClientInstance function
func getKWClientInstance() *common.KWClient {
	if GlobalKWClient == nil {

		url, user, token, err := LoadKWauthInfo()

		if err != nil {
			fmt.Println("Can't auth info:", err)
			os.Exit(1)
		}

		GlobalKWClient = common.NewKWClient(url, user, token)
	}
	return GlobalKWClient
}

var rootCmd = &cobra.Command{
	Use:   "kwcli",
	Short: "Interact with Klocwork servers using the Web API",
	Long: `kwcli provides a command-line interface for leveraging
          the Klocwork Web API. Manage projects, reports, and more.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Default action for the root command (could display help)
		fmt.Println("Welcome to kwcli! Use 'kwcli help' for usage.")
	},
}

func Execute() {

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {

	rootCmd.PersistentFlags().StringVar(&cmdUser, "user", "", "Klocwork user")
	rootCmd.PersistentFlags().StringVar(&cmdToken, "token", "", "Klocwork password")
	rootCmd.PersistentFlags().StringVar(&cmdURL, "url", "", "Klocwork base URL (e.g., http://server:8080)")
	rootCmd.PersistentFlags().StringVar(&outputFile, "output", "result.json", "The output file for results")
}
