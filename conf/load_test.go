package conf

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"reflect"
	"testing"
)

func TestConfig(t *testing.T) {
	LoadConfigFromToml("../dist/config.toml")
	fmt.Printf("%+v\n", global.App.Name)
}

func TestJsonConfig(t *testing.T) {
	jsonFile, err := os.Open("../config.json")
	if err != nil {
		return
	}
	jsonData, err := io.ReadAll(jsonFile)
	if err != nil {
		return
	}
	cfg := newConfig()
	err = json.Unmarshal(jsonData, &cfg)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(reflect.TypeOf(cfg.TelnetCmds))
	fmt.Printf("%+v\n", cfg.TelnetCmds)
}
