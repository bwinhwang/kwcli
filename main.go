package main

import (
	"binhong/kwcli/cmd"
	"binhong/kwcli/common"
	"fmt"
	"os"
)

func main() {

	url, user, token, err := common.LoadKWauthInfo()
	if err != nil {
		fmt.Println("Can't kwauth info:", err)
		os.Exit(1)
	}

	cmd.GlobalKWClient = common.NewKWClient(url, user, token)
	cmd.Execute()

}
