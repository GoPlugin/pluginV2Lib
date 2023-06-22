package main

import (
	"github.com/GoPlugin/pluginV2Lib/pluginV2-env/environment"
	"github.com/GoPlugin/pluginV2Lib/pluginV2-env/pkg/helm/chainlink"
	"github.com/GoPlugin/pluginV2Lib/pluginV2-env/pkg/helm/ethereum"
)

func main() {
	err := environment.New(&environment.Config{
		Labels:            []string{"type=construction-in-progress"},
		NamespacePrefix:   "new-environment",
		KeepConnection:    true,
		RemoveOnInterrupt: true,
	}).
		AddHelm(ethereum.New(nil)).
		AddHelm(chainlink.New(0, nil)).
		Run()
	if err != nil {
		panic(err)
	}
}
