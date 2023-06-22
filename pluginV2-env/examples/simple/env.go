package main

import (
	"github.com/GoPlugin/pluginV2Lib/pluginV2-env/environment"
	"github.com/GoPlugin/pluginV2Lib/pluginV2-env/pkg/helm/chainlink"
	"github.com/GoPlugin/pluginV2Lib/pluginV2-env/pkg/helm/ethereum"
	"github.com/GoPlugin/pluginV2Lib/pluginV2-env/pkg/helm/mockserver"
	mockservercfg "github.com/GoPlugin/pluginV2Lib/pluginV2-env/pkg/helm/mockserver-cfg"
)

func main() {
	err := environment.New(&environment.Config{
		NamespacePrefix:   "ztest",
		KeepConnection:    true,
		RemoveOnInterrupt: true,
	}).
		AddHelm(mockservercfg.New(nil)).
		AddHelm(mockserver.New(nil)).
		AddHelm(ethereum.New(nil)).
		AddHelm(chainlink.New(0, map[string]interface{}{
			"replicas": 1,
		})).
		Run()
	if err != nil {
		panic(err)
	}
}
