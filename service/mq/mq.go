package main

import (
	"flag"
	"fmt"
	"forum/common/utils"
	"forum/service/mq/internal/config"
	"forum/service/mq/internal/listen"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
)

var configFile = flag.String("f", "etc/mq.yaml", "Specify the config file")

func main() {
	flag.Parse()
	var c config.Config

	// 使用我们自己相对路径
	configPath := utils.GetDefaultConfigPath(*configFile)
	conf.MustLoad(configPath, &c)

	// log、prometheus、trace、metricsUrl.
	if err := c.SetUp(); err != nil {
		panic(err)
	}

	serviceGroup := service.NewServiceGroup()
	defer serviceGroup.Stop()

	for _, mq := range listen.Mqs(c) {
		serviceGroup.Add(mq)
	}

	fmt.Println("starting mq service...")
	serviceGroup.Start()
}
