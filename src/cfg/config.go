package cfg

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type log struct {
	MaxSize    int    `json:"maxsize" yaml:"maxsize"`
	MaxAge     int    `json:"maxage" yaml:"maxage"`
	MaxBackups int    `json:"maxbackups" yaml:"maxbackups"`
	LocalTime  bool   `json:"localtime" yaml:"localtime"`
	Compress   bool   `json:"compress" yaml:"compress"`
	Level      int8   `json:"level" yaml:"level"`
	Path       string `json:"path" yaml:"path"`
}

type gate struct {
	LogName   string `json:"logname" yaml:"logname"`
	Ip        string `json:"ip" yaml:"ip"`
	Port      string `json:"port" yaml:"port"`
	Id        int16  `json:"id" yaml:"id"`
	WorkNum   int    `json:"worknum" yaml:"worknum"`
	BufferMax int    `json:"buffermax" yaml:"buffermax"`
	DB        string `json:"db" yaml:"db"`
}

type gs struct {
	Id      int16  `json:"id" yaml:"id"`
	LogName string `json:"logname" yaml:"logname"`
	Ip      string `json:"ip" yaml:"ip"`
	Port    string `json:"port" yaml:"port"`
	DB      string `json:"db" yaml:"db"`
}

type config struct {
	GS   []gs   `json:"gs" yaml:"gs"`
	GATE []gate `json:"gate" yaml:"gate"`
	LOG  log    `json:"log" yaml:"log"`
}

var Config config

func init() {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		fmt.Println("gopath nil")
		os.Exit(-1)
	}
	readConfig(gopath + "/config.yaml")
}

func readConfig(file string) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	err = yaml.Unmarshal(content, &Config)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func GetGsConfig(id int16) *gs {
	for k := range Config.GS {
		if Config.GS[k].Id == id {
			return &Config.GS[k]
		}
	}
	return nil
}

func GetGateConfig(id int16) *gate {
	for k := range Config.GATE {
		if Config.GATE[k].Id == id {
			return &Config.GATE[k]
		}
	}
	return nil
}
