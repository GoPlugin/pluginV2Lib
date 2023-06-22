package main

import (
	"github.com/GoPlugin/pluginV2Lib/pluginV2-env/environment"
	"github.com/GoPlugin/pluginV2Lib/pluginV2-env/pkg/helm/chainlink"
	"github.com/GoPlugin/pluginV2Lib/pluginV2-env/pkg/helm/ethereum"
)

func main() {
	e := environment.New(nil).
		AddHelm(ethereum.New(nil)).
		AddHelm(chainlink.New(0, nil))
	if err := e.Run(); err != nil {
		panic(err)
	}
	if err := e.DumpLogs("logs/mytest"); err != nil {
		panic(err)
	}
}
