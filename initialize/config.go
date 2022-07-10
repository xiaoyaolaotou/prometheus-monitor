package initialize

import (
	"io/ioutil"
	"log"
	"monitor/global"

	"gopkg.in/yaml.v2"
)

func InitConfig() {
	ymlFile, err := ioutil.ReadFile("config/prometheus.yml")
	if err != nil {
		log.Fatal("config.yml read error: ", err)
		return
	}

	err = yaml.Unmarshal(ymlFile, &global.PromConfig)
	if err != nil {
		log.Fatal("config.yaml unmarshal error: ", err)
	}
}
