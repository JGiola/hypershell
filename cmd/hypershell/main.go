package main

import (
	"github.com/golang/glog"

	localapi "github.com/openshift-online/hypershell/pkg/api"
	pkgcmd "github.com/openshift-online/rh-trex-ai/pkg/cmd"

	_ "github.com/openshift-online/hypershell/cmd/hypershell/environments"
	_ "github.com/openshift-online/rh-trex-ai/plugins/events"
	_ "github.com/openshift-online/rh-trex-ai/plugins/generic"
	_ "github.com/openshift-online/hypershell/plugins/fleets"
	_ "github.com/openshift-online/hypershell/plugins/managedClusters"
	_ "github.com/openshift-online/hypershell/plugins/managedDatabases"
	_ "github.com/openshift-online/hypershell/plugins/gatewayReleases"
	_ "github.com/openshift-online/hypershell/plugins/gateways"
	_ "github.com/openshift-online/hypershell/plugins/gatewayNetworks"
)

func main() {
	rootCmd := pkgcmd.NewRootCommand("hypershell", "My service built with TRex library")
	rootCmd.AddCommand(
		pkgcmd.NewMigrateCommand("hypershell"),
		pkgcmd.NewServeCommand(localapi.GetOpenAPISpec),
	)

	if err := rootCmd.Execute(); err != nil {
		glog.Fatalf("error running command: %v", err)
	}
}