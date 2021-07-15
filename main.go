package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	config Config
	client http.Client
)

type Config struct {
	Url            string `json:"url"`
	Method         string `json:"method"`
	JsonBody       string `json:"json_body"`
	RequestCount   int    `json:"request_count"`
	RequestTimeout int    `json:"request_timeout"`
	TimeoutSeconds int    `json:"timeout_seconds"`
	AccessToken    string `json:"access_token"`
}

func main() {
	config.Read("config.json")

	for true {

		for i := 1; i <= config.RequestCount; i++ {
			fmt.Println("Request sending to " + config.Url)
			time.Sleep(time.Duration(config.RequestTimeout) * time.Millisecond)
			go sendRequest()

		}

		fmt.Printf("Waiting %d sec...\n", config.TimeoutSeconds)
		time.Sleep(time.Duration(config.TimeoutSeconds) * time.Second)
	}
}

func sendRequest() {
	request, err := http.NewRequest(config.Method, config.Url, bytes.NewBufferString(config.JsonBody))
	if err != nil {
		panic(err)
	}
	do, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer do.Body.Close()
}

func (c *Config) Read(filepath string) *Config {

	configFile, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil
	}

	err = json.Unmarshal(configFile, &c)
	if err != nil {
		panic(err)
	}

	return c
}
