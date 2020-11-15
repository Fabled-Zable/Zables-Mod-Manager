package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pelletier/go-toml"
)

//Config ... Will hold settings to make this app work
type Config struct {
	Kagpath string
}

//Load ... Reads in the toml file into config variable
func (c *Config) Load() {
	if _, err := os.Stat("./config.toml"); err == nil {
		fmt.Println("Found config, loading")
	} else if os.IsNotExist(err) {
		fmt.Println("Config not found, creating, then loading")

		if f, err := os.Create("./config.toml"); err != nil {
			fmt.Println("Error creating file, panicing")
			panic(err)
		} else {

			b, _ := toml.Marshal(Config{})

			_, err := f.WriteString(string(b))

			if err != nil {
				fmt.Println("Error writing to config file, panicing")
				panic(err)
			}
		}
	} else {
		fmt.Print("Error checking if config file exists")
		panic(err)
	}

	if data, err := ioutil.ReadFile("./config.toml"); err != nil {
		fmt.Println("Error reading confgig")
		panic(err)
	} else if err := toml.Unmarshal(data, c); err != nil {
		fmt.Println("Error parsing config")
		panic(err)
	}
}

func main() {
	var config Config
	config.Load()

	fmt.Println("Welcome to Zable's Mod Manager")
}
