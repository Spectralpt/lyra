package globals

import (
	_ "embed"
	"lyra/types"
)

//go:embed modules/base.yaml
var baseYaml string

//go:embed modules/iot-agent-json.yaml
var iotJson string

//go:embed modules/quantum-leap.yaml
var quantumLeap string

var (
	AvailableLanguages = [...]string{
		"C",
		"C++",
	}
	AvailableLicenses = [...]string{
		"apache",
		"mit",
		"gplv3",
	}
	AvailableModules = [...]types.Module{
		{
			Name:        "fiware-context-broker",
			Description: "A module for managing context information in Fiware.",
		},
		{
			Name:        "fiware-iot-agent",
			Description: "A module for managing IoT devices and data in Fiware.",
			Overlay:     iotJson,
		},
		{
			Name:        "fiware-orion",
			Description: "A module for the Orion Context Broker, a key component of Fiware.",
			Overlay:     baseYaml,
		},
		{
			Name:        "fiware-cygnus",
			Description: "A module for the Cygnus connector, used to persist context data in Fiware.",
		},
		{
			Name:        "fiware-perseo",
			Description: "A module for the Perseo complex event processing engine in Fiware.",
		},
		{
			Name:        "fiware-quantumleap",
			Overlay:     quantumLeap,
			Description: "A module for the QuantumLeap connector, used to store time series data in Fiware.",
		},
		{
			Name:        "fiware-graphql",
			Description: "A module for the GraphQL API in Fiware, allowing flexible data queries.",
		},
	}
)
