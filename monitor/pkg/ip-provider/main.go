package ip_provider

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/trustedshops-public/mail-blocklist-monitor/pkg/util"
	"strings"
)

// IpProvider gets our ips from some source to process
type IpProvider interface {
	// Name of the provider, also used for env variable prefix
	Name() string
	// GetAvailableIps from provider
	GetAvailableIps(ctx context.Context) ([]string, error)
	// ShouldFailSoft specifies if the monitor should crash on errors (hard==true) or continue with a warning (soft==false)
	ShouldFailSoft() bool
}

var _provider []IpProvider

func registerProvider(i IpProvider) {
	_provider = append(_provider, i)
}

func GetProvider() []IpProvider {
	return _provider
}

// getEnvConfig loads the env variable in format IP_PROVIDER_<<PROVIDERNAME>>_<<KEY>>
func getEnvConfig(i IpProvider, key string, defaultVal string) *string {
	val := util.GetEnvWithDefault(strings.ToUpper(fmt.Sprintf("IP_PROVIDER_%s_%s", i.Name(), key)), defaultVal)
	return &val
}

// LoadAllIps from all providers
func LoadAllIps(ctx context.Context) []string {
	ips := make([]string, 0)

	for _, i := range GetProvider() {
		providerIps, err := i.GetAvailableIps(ctx)
		if err != nil {
			if i.ShouldFailSoft() {
				log.Warn("Error while executing provider ", i.Name(), " ignoring, because soft fail is set")
			} else {
				log.Fatal("Error while calling provider ", i.Name(), " ", err.Error())
			}
		}
		ips = append(ips, providerIps...)
	}

	return ips
}
