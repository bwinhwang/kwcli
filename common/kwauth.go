package common

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

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
	return url, username, token, err
}
