package powerdns

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("PDNS_API_KEY", nil),
				Description: "REST API authentication key",
			},
			"server_url": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("PDNS_SERVER_URL", nil),
				Description: "Location of PowerDNS server",
			},
			"insecure_https": {
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("PDNS_INSECURE_HTTPS", false),
				Description: "Disable verification of the PowerDNS server's TLS certificate",
			},
			"ca_certificate": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("PDNS_CACERT", ""),
				Description: "Content or path of a Root CA to be used to verify PowerDNS's SSL certificate",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"powerdns_zone":   resourcePDNSZone(),
			"powerdns_record": resourcePDNSRecord(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(data *schema.ResourceData) (interface{}, error) {
	config := Config{
		ApiKey:        data.Get("api_key").(string),
		ServerUrl:     data.Get("server_url").(string),
		InsecureHTTPS: data.Get("insecure_https").(bool),
		CACertificate: data.Get("ca_certificate").(string),
	}

	return config.Client()
}
