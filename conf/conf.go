package conf

import (
	"flag"
	"fmt"
	"os"

	toml "github.com/pelletier/go-toml"
)

var (
	conf = flag.String("c", "./config.toml", "config file")
)

var (
	// for app config
	APP_Address string

	// for db config
	DB_Driver, DB_Connect string

	// for ss config
	SS_Image string

	// for host config
	HOST_Address       string
	HOST_From, HOST_To int
)

func InitConf() {
	//Check config file
	if _, err := os.Stat(*conf); os.IsNotExist(err) {
		fmt.Println("Cannot load config.toml, file doesn't exist...")
		os.Exit(1)
	}

	config, err := toml.LoadFile(*conf)

	if err != nil {
		fmt.Println("Failed to load config file:", *conf)
		fmt.Println(err.Error())
		os.Exit(1)
	}
	APP_Address = config.Get("app.address").(string)

	SS_Image = config.Get("ss.image").(string)

	HOST_Address = config.Get("host.address").(string)
	HOST_From = int(config.Get("host.from").(int64))
	HOST_To = int(config.Get("host.to").(int64))

	DB_Driver = config.Get("db.driver").(string)
	DB_Connect = config.Get("db.connect").(string)
}