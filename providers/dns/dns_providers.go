// Factory for DNS providers
package dns

import (
	"fmt"

	"github.com/shankj3/lego/acme"
	"github.com/shankj3/lego/providers/dns/auroradns"
	"github.com/shankj3/lego/providers/dns/azure"
	"github.com/shankj3/lego/providers/dns/cloudflare"
	"github.com/shankj3/lego/providers/dns/digitalocean"
	"github.com/shankj3/lego/providers/dns/dnsimple"
	"github.com/shankj3/lego/providers/dns/dnsmadeeasy"
	"github.com/shankj3/lego/providers/dns/dnspod"
	"github.com/shankj3/lego/providers/dns/dyn"
	"github.com/shankj3/lego/providers/dns/exoscale"
	"github.com/shankj3/lego/providers/dns/gandi"
	"github.com/shankj3/lego/providers/dns/googlecloud"
	"github.com/shankj3/lego/providers/dns/linode"
	"github.com/shankj3/lego/providers/dns/namecheap"
	"github.com/shankj3/lego/providers/dns/ns1"
	"github.com/shankj3/lego/providers/dns/otc"
	"github.com/shankj3/lego/providers/dns/ovh"
	"github.com/shankj3/lego/providers/dns/pdns"
	"github.com/shankj3/lego/providers/dns/rackspace"
	"github.com/shankj3/lego/providers/dns/rfc2136"
	"github.com/shankj3/lego/providers/dns/route53"
	"github.com/shankj3/lego/providers/dns/vultr"
)

func NewDNSChallengeProviderByName(name string) (acme.ChallengeProvider, error) {
	var err error
	var provider acme.ChallengeProvider
	switch name {
	case "azure":
		provider, err = azure.NewDNSProvider()
	case "auroradns":
		provider, err = auroradns.NewDNSProvider()
	case "cloudflare":
		provider, err = cloudflare.NewDNSProvider()
	case "digitalocean":
		provider, err = digitalocean.NewDNSProvider()
	case "dnsimple":
		provider, err = dnsimple.NewDNSProvider()
	case "dnsmadeeasy":
		provider, err = dnsmadeeasy.NewDNSProvider()
	case "dnspod":
		provider, err = dnspod.NewDNSProvider()
	case "dyn":
		provider, err = dyn.NewDNSProvider()
	case "exoscale":
		provider, err = exoscale.NewDNSProvider()
	case "gandi":
		provider, err = gandi.NewDNSProvider()
	case "gcloud":
		provider, err = googlecloud.NewDNSProvider()
	case "linode":
		provider, err = linode.NewDNSProvider()
	case "manual":
		provider, err = acme.NewDNSProviderManual()
	case "namecheap":
		provider, err = namecheap.NewDNSProvider()
	case "rackspace":
		provider, err = rackspace.NewDNSProvider()
	case "route53":
		provider, err = route53.NewDNSProvider()
	case "rfc2136":
		provider, err = rfc2136.NewDNSProvider()
	case "vultr":
		provider, err = vultr.NewDNSProvider()
	case "ovh":
		provider, err = ovh.NewDNSProvider()
	case "pdns":
		provider, err = pdns.NewDNSProvider()
	case "ns1":
		provider, err = ns1.NewDNSProvider()
	case "otc":
		provider, err = otc.NewDNSProvider()
	default:
		err = fmt.Errorf("Unrecognised DNS provider: %s", name)
	}
	return provider, err
}
