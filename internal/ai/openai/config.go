package openai

import "github.com/mylxsw/aidea-server/config"

type Config struct {
	Enable             bool
	OpenAIAzure        bool
	OpenAIAPIVersion   string
	OpenAIOrganization string
	OpenAIServers      []string
	OpenAIKeys         []string
}

func parseMainConfig(conf *config.Config) *Config {
	return &Config{
		Enable:             conf.EnableOpenAI,
		OpenAIAzure:        conf.OpenAIAzure,
		OpenAIAPIVersion:   conf.OpenAIAPIVersion,
		OpenAIOrganization: conf.OpenAIOrganization,
		OpenAIServers:      conf.OpenAIServers,
		OpenAIKeys:         conf.OpenAIKeys,
	}
}

func parseBackupConfig(conf *config.Config) *Config {
	return &Config{
		Enable:             conf.EnableFallbackOpenAI,
		OpenAIAzure:        conf.FallbackOpenAIAzure,
		OpenAIAPIVersion:   conf.FallbackOpenAIAPIVersion,
		OpenAIOrganization: conf.FallbackOpenAIOrganization,
		OpenAIServers:      conf.FallbackOpenAIServers,
		OpenAIKeys:         conf.FallbackOpenAIKeys,
	}
}

func parseDalleConfig(conf *config.Config) *Config {
	if conf.DalleUsingOpenAISetting {
		return &Config{
			Enable:             conf.EnableOpenAI && conf.EnableOpenAIDalle,
			OpenAIAzure:        conf.OpenAIAzure,
			OpenAIAPIVersion:   conf.OpenAIAPIVersion,
			OpenAIOrganization: conf.OpenAIOrganization,
			OpenAIServers:      conf.OpenAIServers,
			OpenAIKeys:         conf.OpenAIKeys,
		}
	}

	return &Config{
		Enable:             conf.EnableOpenAIDalle,
		OpenAIAzure:        conf.OpenAIDalleAzure,
		OpenAIAPIVersion:   conf.OpenAIDalleAPIVersion,
		OpenAIOrganization: conf.OpenAIDalleOrganization,
		OpenAIServers:      conf.OpenAIDalleServers,
		OpenAIKeys:         conf.OpenAIDalleKeys,
	}
}
