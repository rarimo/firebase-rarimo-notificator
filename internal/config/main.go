package config

import (
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/copus"
	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/kit/kv"
	"log"
)

type Config interface {
	comfig.Logger
	types.Copuser
	comfig.Listenerer

	// Method to get admin_sdk paths
	AdminSDKPaths() map[string]string
}

type config struct {
	comfig.Logger
	types.Copuser
	comfig.Listenerer
	getter        kv.Getter
	adminSDKPaths map[string]string
}

func New(getter kv.Getter) Config {
	return &config{
		getter:        getter,
		Copuser:       copus.NewCopuser(getter),
		Listenerer:    comfig.NewListenerer(getter),
		Logger:        comfig.NewLogger(getter, comfig.LoggerOpts{}),
		adminSDKPaths: loadAdminSDKPaths(getter),
	}
}

func loadAdminSDKPaths(getter kv.Getter) map[string]string {
	result := make(map[string]string)

	// Retrieve the map from the configuration
	paths, err := getter.GetStringMap("admin_sdk")
	if err != nil {
		log.Printf("Failed to load admin_sdk paths: %v", err)
		return result
	}

	// Convert map[string]interface{} to map[string]string
	for key, value := range paths {
		if strVal, ok := value.(string); ok {
			result[key] = strVal
		} else {
			log.Printf("Invalid value type for admin_sdk.%s, expected string", key)
		}
	}

	return result
}

func (c *config) AdminSDKPaths() map[string]string {
	return c.adminSDKPaths
}
