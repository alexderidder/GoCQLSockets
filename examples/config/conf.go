package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configuration struct {
	Client struct {
		IPAddress     string `json:"ip-address"`
		Port          string `json:"port"`
		ReconnectTime int    `json:"reconnect_time"`
		Certs         struct {
			Directory string `json:"directory"`
			Pem       string `json:"pem"`
			Key       string `json:"key"`
		} `json:"certs"`
	} `json:"client"`
}

var Config Configuration

func init() {
	file, _ := os.Open("examples/config/conf.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	Config = Configuration{}
	err := decoder.Decode(&Config)
	if err != nil {
		fmt.Println("error:", err)
	}

}
