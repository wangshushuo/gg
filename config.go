package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

//	// 获取用户主目录
//	homeDir, _ := os.UserHomeDir()
//
//	// 拼接文件路径
//	filePath := filepath.Join(homeDir, ".ggrc")
//
//	// 打开文件
//	file, err := os.Open(filePath)

type Config struct {
	Server   string `yaml:"server"`
	Port     int    `yaml:"port"`
	Database struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"database"`
}

func read() {
	// 读取配置文件
	configFile := getFilePath("~/.ggrc")
	configData, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	var config Config
	err = yaml.Unmarshal(configData, &config)
	if err != nil {
		log.Fatalf("Error unmarshalling config data: %v", err)
	}

	// 打印当前配置
	fmt.Printf("Current Config:\n%+v\n", config)

	// 修改配置
	config.Server = "new.example.com"
	config.Port = 8081
	config.Database.Username = "newuser"
	config.Database.Password = "newpass"

	// 将修改后的配置写入文件
	newConfigData, err := yaml.Marshal(&config)
	if err != nil {
		log.Fatalf("Error marshalling config data: %v", err)
	}

	err = ioutil.WriteFile(configFile, newConfigData, 0644)
	if err != nil {
		log.Fatalf("Error writing config file: %v", err)
	}

	fmt.Println("Config file updated successfully!")
}

func getFilePath(filename string) string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Error getting user's home directory: %v", err)
	}

	return filepath.Join(homeDir, filename)
}
