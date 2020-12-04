package configproviders

import (
	"log"
	"os"
)

type IEnvironmentVariableProvider interface {
	Get(name string) string
}

type EnvironmentVariableProvider struct{}

func (e EnvironmentVariableProvider) Get(name string) string {
	v := os.Getenv(name)
	if v == "" {
		log.Println("value not found : " + name)
	}
	return v
}
