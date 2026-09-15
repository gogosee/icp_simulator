package main

import (
	"flag"
	"os"
	"strings"

	"github.com/gogosee/icp_simulator/log"
	sms "github.com/gogosee/icp_simulator/smc_client"

	"github.com/gogosee/icp_simulator/auth"
	// "github.com/gogosee/icp_simulator/msc_client"
)

func main() {
	// 启动记录数据库的程序
	if sms.ConfigYml.GetString("Mongo.URI") != "" {
		sms.PersistenceSmsJournal()
	} else {
		sms.StartCacheExpireTicker(nil)
	}

	auth.Cache = auth.New(sms.ConfigYml)

	phone := flag.String("p", "13800001111,13300001111,18600001111", "phone")
	message := flag.String("m", "hello world, 你好世界！", "message")
	iterates := flag.Int("i", 1, "iterates")
	flag.Parse()

	go func() {
		i := *iterates
		arr := strings.Split(*phone, ",")
		for i > 0 {
			for _, s := range arr {
				err := sms.AsyncPool().Submit(func() {
					sms.Send(*message, s)
				})
				if err != nil {
					log.Errorf("AsyncPool Error: %v", err)
					return
				}
			}
			i--
		}
	}()

	<-sms.StatChan()

	log.Warn("main goroutine exit.")
	log.Sync()
	os.Exit(0)
}
