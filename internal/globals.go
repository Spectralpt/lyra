package utils

type Module struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

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
	AvailableModules = [...]Module{
		{
			Name:        "fiware-context-broker",
			Description: "A module for managing context information in Fiware.",
		},
		{
			Name:        "fiware-iot-agent",
			Description: "A module for managing IoT devices and data in Fiware.",
		},
		{
			Name:        "fiware-orion",
			Description: "A module for the Orion Context Broker, a key component of Fiware.",
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
			Description: "A module for the QuantumLeap connector, used to store time series data in Fiware.",
		},
		{
			Name:        "fiware-graphql",
			Description: "A module for the GraphQL API in Fiware, allowing flexible data queries.",
		},
	}
)
