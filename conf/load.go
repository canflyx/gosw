package conf

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/caarlos0/env/v6"
	"gopkg.in/yaml.v3"
)

var (
	global *Config
)

// C 全局配置对象
func C() *Config {
	if global == nil {
		panic("Load Config first")
	}
	return global
}

// LoadConfigFromToml 从toml中添加配置文件, 并初始化全局对象
func LoadConfigFromToml(filePath string) error {
	cfg := newConfig()
	if _, err := toml.DecodeFile(filePath, cfg); err != nil {
		return err
	}
	// 加载全局配置单例
	global = cfg
	return nil
}
func LoadConfigFromJson(filePath string) error {
	cfg := newConfig()
	defaultConfig := "config.json"
	if filePath == "" {
		filePath = defaultConfig
	}
	jsonFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println(CreateConfig(defaultConfig, cfg))
	}
	jsonData, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
	}
	if err := json.Unmarshal(jsonData, &cfg); err != nil {
		fmt.Println(err)
	}
	// 加载全局配置单例
	global = cfg
	return nil

}
func LoadConfigFromYaml(filePath string) error {
	config := newConfig()
	configFile := "config.yaml"

	file, err := os.ReadFile(configFile)
	if err != nil {
		yamlData, _ := yaml.Marshal(&config)
		err = os.WriteFile(configFile, yamlData, 0644)
		if err != nil {
			fmt.Println("Unable to write data into the file")
		}
	}
	err2 := yaml.Unmarshal(file, &config)
	if err2 != nil {
		fmt.Println(err2)
	}
	global = config
	return nil
}

// 如果没有默认文件，直接写个默认的
func CreateConfig(filename string, config *Config) error {
	filePtr, err := os.Create(filename)
	if err != nil {
		fmt.Println("文件创建失败", err.Error())
		return err
	}
	defer filePtr.Close()
	// 创建Json编码器
	encoder := json.NewEncoder(filePtr)
	err = encoder.Encode(config)
	if err != nil {
		fmt.Println("编码错误", err.Error())
	} else {
		fmt.Println("编码成功")
	}
	return nil
}

// LoadConfigFromEnv 从环境变量中加载配置
func LoadConfigFromEnv() error {
	cfg := newConfig()
	if err := env.Parse(cfg); err != nil {
		return err
	}
	// 加载全局配置单例
	global = cfg
	return nil
}
