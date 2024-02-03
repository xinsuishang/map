package main

import (
	"fmt"
	"github.com/google/uuid"
	"msp/biz_server/gpt/internal/infra/thridpart/bailian"
	commoninit "msp/common/initialize"
	"msp/common/model/config"
	"strings"
	"time"
)

func main() {
	//svr := gpt.NewServer(new(ChatServiceImpl))
	//
	//err := svr.Run()
	//
	//if err != nil {
	//	log.Println(err.Error())
	//}
	// 本地配置文件初始化
	//initialize.InitLocalConfig()
	// 根据本地配置文件初始化日志
	// commoninit.InitLogger(&config.GlobalLocalConfig.LogConfig, config.GlobalServerConfig.Name, config.IsDev)
	commoninit.InitLogger(&config.LogConfig{
		LogDir:   "/home/work/log",
		LogDebug: true,
	}, "gpt", true)
	sessionId := strings.ReplaceAll(uuid.New().String(), "-", "")

	start := time.Now()
	for i := 0; i < 100; i++ {
		requestId := strings.ReplaceAll(uuid.New().String(), "-", "")
		prompt := fmt.Sprintf("帮我计算%d+%d", i, i)
		curStart := time.Now()
		completion, err := bailian.CreateCompletion(1, requestId, sessionId, prompt)
		if err != nil {
			print(err)
			break
		}
		now := time.Now()
		fmt.Printf("cast sum : %s, curCost %s, prompt: %s, result: %s\n", now.Sub(start), now.Sub(curStart), prompt, completion)
	}
}
