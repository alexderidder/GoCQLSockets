package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configuration  struct {
	Server struct {
		IPAddress string `json:"ip-address"`
		Port      string `json:"port"`
	} `json:"server"`
	Database struct {
		Clusters []struct {
			Name      string `json:"name"`
			IPAddress string `json:"ip-address"`
		} `json:"clusters"`
		Keyspace string `json:"keyspace"`
	} `json:"database"`
	Client struct {
		IPAddress string `json:"ip-address"`
		Port      string `json:"port"`
	} `json:"client"`
}

var Config Configuration

func init() {
	file, _ := os.Open("conf.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	Config = Configuration{}
	err := decoder.Decode(&Config)
	if err != nil {
		fmt.Println("error:", err)
	}

}
