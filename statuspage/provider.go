package statuspage

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	sp "github.com/yannh/statuspage-go-sdk"
)

func Provider() *schema.Provider {
	return statusPageProvider()
}

func statusPageProvider() *schema.Provider {
	p := &schema.Provider{
		Schema: newSchema(),
		ResourcesMap: map[string]*schema.Resource{
		"statuspage_component":        resourceComponent(),
		"statuspage_component_group":  resourceComponentGroup(),
		"statuspage_metric":           resourceMetric(),
		"statuspage_metrics_provider": resourceMetricsProvider(),
	}}
	p.ConfigureContextFunc = providerConfigure(p)

	return p
}

func newSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"token": {
			Type:        schema.TypeString,
			Required:    true,
			DefaultFunc: schema.EnvDefaultFunc("STATUSPAGE_TOKEN", ""),
		},
	}
}

//func providerConfigure(d *schema.ResourceData) (interface{}, error) {
//	return sp.NewClient(d.Get("token").(string)), nil
//}

func providerConfigure(p *schema.Provider) schema.ConfigureContextFunc  {
	return func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		return sp.NewClient(d.Get("token").(string)), nil
	}
}