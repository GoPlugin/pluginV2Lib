package main

import (
	"github.com/GoPlugin/pluginV2Lib/pluginV2-env/environment"
	"github.com/GoPlugin/pluginV2Lib/pluginV2-env/examples/deployment_part_cdk8s"
	"github.com/GoPlugin/pluginV2Lib/pluginV2-env/pkg/helm/chainlink"
	"github.com/GoPlugin/pluginV2Lib/pluginV2-env/pkg/helm/ethereum"
)

func main() {
	e := environment.New(nil).
		AddChart(deployment_part_cdk8s.New(&deployment_part_cdk8s.Props{})).
		AddHelm(ethereum.New(nil)).
		AddHelm(chainlink.New(0, map[string]interface{}{
			"replicas": 2,
		}))
	if err := e.Run(); err != nil {
		panic(err)
	}
	e.Shutdown()
}
