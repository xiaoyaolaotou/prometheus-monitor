package initialize

import (
	"fmt"
	"log"
	"monitor/global"

	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

func InitApis() {
	APIs := make(map[string]v1.API)

	for _, value := range global.PromConfig.Services {
		config := api.Config{
			Address: value.Address,
		}
		client, err := api.NewClient(config)
		if err != nil {
			log.Println("api.NewClient error : ", err)
			return
		}
		api := v1.NewAPI(client)
		APIs[value.Name] = api
	}
	global.APIs = APIs
}
