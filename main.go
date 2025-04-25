package main

import (
	"log"

	"github.com/eclipse-xfsc/cloud-wallet-plugin-discovery/services/plugins"
	"github.com/eclipse-xfsc/cloud-wallet-plugin-discovery/types"
	serviceCore "gitlab.eclipse.org/eclipse/xfsc/libraries/microservice/core"
	serviceTypes "gitlab.eclipse.org/eclipse/xfsc/libraries/microservice/core/types"
)

func main() {
	env := types.GetEnvironment()
	err := serviceCore.InitializeService("api", env)
	services := make([]serviceTypes.Service, 0)
	services = append(services, new(plugins.PluginsService))
	serviceCore.RegisterServices(services)
	if err == nil {
		err = serviceCore.StartService()
		if err != nil {
			log.Fatalf("unexpected error, %v", err)
		}
	} else {
		log.Fatalf("unexpected error, %v", err)
	}
}
