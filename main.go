package main
//1111111
import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/terraform-providers/terraform-provider-fortios/fortios"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: fortios.Provider})
}
