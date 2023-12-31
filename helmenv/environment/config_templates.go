package environment

import (
	"path/filepath"
	"reflect"

	"github.com/rs/zerolog/log"
	"github.com/GoPlugin/pluginV2Lib/helmenv/tools"
)

const (
	mockServerConfigChartName = "mockserver-config"
	mockServerChartName       = "mockserver"
	chainlinkChartName        = "chainlink"
)

// NewChainlinkChart returns a default Chainlink Helm chart based on a set of override values
func NewChainlinkChart(index int, values map[string]interface{}) *HelmChart {
	return &HelmChart{Values: values, Index: index}
}

// NewChainlinkCCIPReorgConfig returns a Chainlink environment for the purpose of CCIP testing
func NewChainlinkCCIPReorgConfig(chainlinkValues map[string]interface{}, networkIDs []int) *Config {
	return &Config{
		NamespacePrefix: "chainlink-ccip",
		Charts: Charts{
			"geth-reorg": {
				Index:       1,
				ReleaseName: "geth-reorg",
				Path:        filepath.Join(tools.ChartsRoot, "geth-reorg"),
				Values: map[string]interface{}{
					"geth": map[string]interface{}{
						"genesis": map[string]interface{}{
							"networkId": networkIDs[0],
						},
					},
				},
			},
			"geth-reorg-2": {
				Index:       1,
				ReleaseName: "geth-reorg-2",
				Path:        filepath.Join(tools.ChartsRoot, "geth-reorg"),
				Values: map[string]interface{}{
					"geth": map[string]interface{}{
						"genesis": map[string]interface{}{
							"networkId": networkIDs[1],
						},
					},
				},
			},
			"chainlink": NewChainlinkChart(3, ChainlinkReplicas(5, chainlinkValues)),
		},
	}
}

// NewTerraChainlinkConfig returns a Chainlink environment designed for testing with a Terra relay
func NewTerraChainlinkConfig(chainlinkValues map[string]interface{}) *Config {
	return &Config{
		NamespacePrefix: "chainlink-terra",
		Charts: Charts{
			"localterra": {Index: 1},
			"geth-reorg": {Index: 2},
			"chainlink":  NewChainlinkChart(3, ChainlinkReplicas(2, chainlinkValues)),
		},
	}
}

// NewChainlinkReorgConfig returns a Chainlink environment designed for simulating re-orgs within testing
func NewChainlinkReorgConfig(chainlinkValues map[string]interface{}) *Config {
	return &Config{
		NamespacePrefix: "chainlink-reorg",
		Charts: Charts{
			"geth-reorg": {Index: 1},
			"chainlink":  NewChainlinkChart(2, chainlinkValues),
		},
	}
}

// Organizes passed in values for simulated network charts
type networkChart struct {
	Replicas int
	Values   map[string]interface{}
}

// NewChainlinkConfig returns a vanilla Chainlink environment used for generic functional testing. Geth networks can
// be passed in to launch differently configured simulated geth instances.
func NewChainlinkConfig(
	chainlinkValues map[string]interface{},
	optionalNamespacePrefix string,
	networks ...SimulatedNetwork,
) *Config {
	charts := Charts{
		mockServerConfigChartName: {Index: 1},
		mockServerChartName:       {Index: 2},
		chainlinkChartName:        NewChainlinkChart(2, chainlinkValues),
	}

	nameSpacePrefix := loadNetworkCharts(optionalNamespacePrefix, charts, networks)

	return &Config{
		NamespacePrefix: nameSpacePrefix,
		Charts:          charts,
	}
}

// NewPerformanceChainlinkConfig launches an environment with upgraded resources
// Mockserver launches with 1 CPU and 1 GB of RAM
// Chainlink DB launches with 1 CPU and 2GB of RAM
// Chainlink node launches with 2 CPU and 4GB RAM
func NewPerformanceChainlinkConfig(
	chainlinkValues map[string]interface{},
	optionalNamespacePrefix string,
	networks ...SimulatedNetwork,
) *Config {
	// Set the chainlink value resources to a higher level
	chainlinkValues["chainlink"] = map[string]interface{}{
		"resources": map[string]interface{}{
			"requests": map[string]interface{}{
				"cpu":    "2",
				"memory": "4096Mi",
			},
			"limits": map[string]interface{}{
				"cpu":    "2",
				"memory": "4096Mi",
			},
		},
	}
	chainlinkValues["db"] = map[string]interface{}{
		"resources": map[string]interface{}{
			"requests": map[string]interface{}{
				"cpu":    "1",
				"memory": "2048Mi",
			},
			"limits": map[string]interface{}{
				"cpu":    "1",
				"memory": "2048Mi",
			},
		},
	}
	mockServerValues := map[string]interface{}{
		"app": map[string]interface{}{
			"resources": map[string]interface{}{
				"requests": map[string]interface{}{
					"cpu":    "1",
					"memory": "1024Mi",
				},
				"limits": map[string]interface{}{
					"cpu":    "1",
					"memory": "1024Mi",
				},
			},
		},
	}
	charts := Charts{
		mockServerConfigChartName: {Index: 1},
		mockServerChartName:       {Index: 2, Values: mockServerValues},
		chainlinkChartName:        NewChainlinkChart(2, chainlinkValues),
	}

	nameSpacePrefix := loadNetworkCharts(optionalNamespacePrefix, charts, networks)

	return &Config{
		NamespacePrefix: nameSpacePrefix,
		Charts:          charts,
	}
}

// loads and properly configures the network charts, and builds a proper namespace config
func loadNetworkCharts(optionalNamespacePrefix string, charts Charts, networks []SimulatedNetwork) string {
	nameSpacePrefix := "chainlink"
	if optionalNamespacePrefix != "" {
		nameSpacePrefix = optionalNamespacePrefix
	}

	networkCharts := map[string]*networkChart{}
	for _, networkFunc := range networks {
		chartName, networkValues := networkFunc()
		if networkValues == nil {
			networkValues = map[string]interface{}{}
		}
		// TODO: If multiple networks with the same chart name are present, only use the values from the first one.
		// This means that we can't have mixed network values with the same type
		// (e.g. all geth deployments need to have the same values).
		// Enabling different behavior is a bit of a niche case.
		if _, present := networkCharts[chartName]; !present {
			networkCharts[chartName] = &networkChart{Replicas: 1, Values: networkValues}
		} else {
			if !reflect.DeepEqual(networkValues, networkCharts[chartName].Values) {
				log.Warn().Msg("If trying to launch multiple networks with different underlying values but the same type, " +
					"(e.g. 1 geth performance and 1 geth realistic), that behavior is not currently fully supported. " +
					"Only replicas of the first of that network type will be launched.")
			}
			networkCharts[chartName].Replicas++
		}
	}

	for chartName, networkChart := range networkCharts {
		networkChart.Values["replicas"] = networkChart.Replicas
		charts[chartName] = &HelmChart{Index: 1, Values: networkChart.Values}
	}
	return nameSpacePrefix
}

// SimulatedNetwork is a function that enables launching a simulated network with a returned chart name
// and corresponding values
type SimulatedNetwork func() (string, map[string]interface{})

// DefaultGeth sets up a basic, low-power simulated geth instance. Really just returns empty map to use default values
func DefaultGeth() (string, map[string]interface{}) {
	return "geth", map[string]interface{}{}
}

// PerformanceGeth sets up the simulated geth instance with more power, bigger blocks, and faster mining
func PerformanceGeth() (string, map[string]interface{}) {
	values := map[string]interface{}{}
	values["resources"] = map[string]interface{}{
		"requests": map[string]interface{}{
			"cpu":    "1",
			"memory": "1024Mi",
		},
		"limits": map[string]interface{}{
			"cpu":    "1",
			"memory": "1024Mi",
		},
	}
	values["config_args"] = map[string]interface{}{
		"--dev.period":      "1",
		"--miner.threads":   "1",
		"--miner.gasprice":  "10000000000",
		"--miner.gastarget": "30000000000",
		"--cache":           "4096",
	}
	return "geth", values
}

// RealisticGeth sets up the simulated geth instance to emulate the actual ethereum mainnet as close as possible
func RealisticGeth() (string, map[string]interface{}) {
	values := map[string]interface{}{}
	values["resources"] = map[string]interface{}{
		"requests": map[string]interface{}{
			"cpu":    "1",
			"memory": "1024Mi",
		},
		"limits": map[string]interface{}{
			"cpu":    "1",
			"memory": "1024Mi",
		},
	}
	values["config_args"] = map[string]interface{}{
		"--dev.period":      "14",
		"--miner.threads":   "1",
		"--miner.gasprice":  "10000000000",
		"--miner.gastarget": "15000000000",
		"--cache":           "4096",
	}

	return "geth", values
}

// ChainlinkVersion sets the version of the chainlink image to use
func ChainlinkVersion(version string, values map[string]interface{}) map[string]interface{} {
	if values == nil {
		values = map[string]interface{}{}
	}
	values["chainlink"] = map[string]interface{}{
		"image": map[string]interface{}{
			"version": version,
		},
	}
	return values
}

// ChainlinkReplicas sets the replica count of chainlink nodes to use
func ChainlinkReplicas(count int, values map[string]interface{}) map[string]interface{} {
	if values == nil {
		values = map[string]interface{}{}
	}
	values["replicas"] = count
	return values
}
