package main

import (
	"context"
	"fmt"
	"github.com/grafana-tools/sdk"
	"github.com/spf13/viper"
	"log"
)

func main() {
	viper.SetConfigName("config")         // name of config file (without extension)
	viper.SetConfigType("yaml")           // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("/etc/appname/")  // path to look for the config file in
	viper.AddConfigPath("$HOME/.termana") // call multiple times to add many search paths
	viper.AddConfigPath(".")              // optionally look for config in the working directory
	err := viper.ReadInConfig()           // Find and read the config file
	if err != nil {                       // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	url := viper.GetString("url")
	key := viper.GetString("apikey")
	client,_ := sdk.NewClient(url, key, sdk.DefaultHTTPClient)
	ctx := context.Background()
	boards, _ := client.Search(ctx)
	for _, b := range boards {
		board,_,_ := client.GetDashboardByUID(ctx, b.UID)
		if board.Title == "SimpleTest" {
			log.Print(board.UID)
		}
	}
}
