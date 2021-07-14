package main

//import (
//	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
//	"github.com/yannh/terraform-provider-statuspage/statuspage"
//)
//
//func main() {
//	plugin.Serve(&plugin.ServeOpts{
//		ProviderFunc: statuspage.Provider,
//	})
//}

import (
	"context"
	"flag"
	"github.com/evanfreed/terraform-provider-statuspage/statuspage"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"log"
)

func main() {
	var debugMode bool

	flag.BoolVar(&debugMode, "debuggable", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	if debugMode {
		err := plugin.Debug(context.Background(), "registry.terraform.io/evanfreed/terraform-provider-statuspage",
			&plugin.ServeOpts{
				ProviderFunc: statuspage.Provider,
			})
		if err != nil {
			log.Println(err.Error())
		}
	} else {
		plugin.Serve(&plugin.ServeOpts{
			ProviderFunc: statuspage.Provider})
	}
}