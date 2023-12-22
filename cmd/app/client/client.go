package main

import (
	"fmt"
	"github.com/Vaixle/proof-of-work/internal/app"
	"github.com/Vaixle/proof-of-work/internal/config"
	"github.com/spf13/viper"
)

func main() {
	err := config.Init()
	if err != nil {
		return
	}

	//Create client
	client := app.NewClient()

	address := fmt.Sprintf("%s:%s", viper.GetString("host.server"), viper.GetString("port.server"))

	//Run request with connection
	client.Run(address)
}
