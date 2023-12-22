package main

import (
	"github.com/Vaixle/proof-of-work/internal/app"
	"github.com/Vaixle/proof-of-work/internal/config"
	"github.com/spf13/viper"
)

func main() {

	// Read config
	err := config.Init()
	if err != nil {
		return
	}

	//Create server
	server := app.NewServer(viper.GetInt("pow.difficulty_bites"))

	//Run tcp server
	server.Run(":" + viper.GetString("port.server"))
}
