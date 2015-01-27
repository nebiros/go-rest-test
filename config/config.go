package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
)

const (
	ApplicationEnvProduction  = "production"
	ApplicationEnvDevelopment = "development"
)

var (
	ApplicationEnv string
	Data           config
)

type config struct {
	Ws struct {
		Auth struct {
			Password string `json:"password"`
			Username string `json:"username"`
		} `json:"auth"`
		Endpoint struct {
			FacturacionLookupByClientID string `json:"facturacion_lookup_by_client_id"`
		} `json:"endpoint"`
		Host string `json:"host"`
	} `json:"ws"`
}

func init() {
	ApplicationEnv = applicationEnv()
	Data = readConfigFile()
}

func applicationEnv() string {
	env := os.Getenv("APPLICATION_ENV")
	if len(strings.TrimSpace(env)) <= 0 {
		env = ApplicationEnvProduction
	}

	return env
}

func readConfigFile() config {
	file, err := ioutil.ReadFile("config/config.json")
	if err != nil {
		panic(err)
	}

	// data := map[string]interface{}{}
	data := map[string]config{}
	err = json.Unmarshal(file, &data)
	if err != nil {
		panic(err)
	}

	envData := data[ApplicationEnv]

	return envData
}
