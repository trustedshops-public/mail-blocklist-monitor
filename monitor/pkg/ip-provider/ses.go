package ip_provider

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	log "github.com/sirupsen/logrus"
	"strings"
)

func init() {
	registerProvider(SesIpProvider{})
}

type SesIpProvider struct {
}

func (s SesIpProvider) Name() string {
	return "ses"
}

func (s SesIpProvider) ShouldFailSoft() bool {
	return true
}

func getDedicatedIps(ctx context.Context, region string) ([]string, error) {
	cfg, err := config.LoadDefaultConfig(ctx, config.WithDefaultRegion(region))
	if err != nil {
		return nil, err
	}

	client := sesv2.NewFromConfig(cfg)

	var ips []string
	var res = &sesv2.GetDedicatedIpsOutput{}

	for {
		res, err = client.GetDedicatedIps(ctx, &sesv2.GetDedicatedIpsInput{})
		if err != nil {
			return nil, err
		}

		for _, ip := range res.DedicatedIps {
			ips = append(ips, *ip.Ip)
		}

		if res.NextToken == nil {
			break
		}
	}

	return ips, nil
}

func (s SesIpProvider) fetchForRegion(ctx context.Context, region string) ([]string, error) {
	log.Info("Getting IPs for region ", region)
	localIps, err := getDedicatedIps(ctx, strings.TrimSpace(region))
	if err != nil {
		return nil, err
	}
	log.Debug("Got ", len(localIps), " IPs for region ", region)
	return localIps, nil
}

// GetAvailableIps from ses dedicated ip pools
func (s SesIpProvider) GetAvailableIps(ctx context.Context) ([]string, error) {
	regions := strings.Split(*getEnvConfig(s, "regions", "eu-central-1"), ",")

	sesIps := make([]string, 0)
	for _, region := range regions {
		localIps, err := s.fetchForRegion(ctx, region)
		if err != nil {
			return nil, err
		}
		sesIps = append(sesIps, localIps...)
	}

	return sesIps, nil
}
