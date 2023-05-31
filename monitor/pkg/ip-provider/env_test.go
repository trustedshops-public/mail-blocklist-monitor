package ip_provider

import (
	"context"
	"os"
	"reflect"
	"testing"
)

var envIpProvider = EnvIpProvider{}

func TestEnvIpProvider_Name(t *testing.T) {
	testProviderName(envIpProvider, t, "env")
}

func TestEnvIpProvider_ShouldFailSoft(t *testing.T) {
	testShouldFailSoft(envIpProvider, t, false)
}

func TestEnvIpProvider_GetAvailableIps(t *testing.T) {
	testCases := []struct {
		envVars map[string]string
		ips     []string
	}{
		{
			map[string]string{
				"IP_PROVIDER_ENV_IPS": "",
			},
			[]string{},
		},
		{
			map[string]string{
				"IP_PROVIDER_ENV_IPS": "127.0.0.1",
			},
			[]string{
				"127.0.0.1",
			},
		},
		{
			map[string]string{
				"IP_PROVIDER_ENV_IPS": "127.0.0.1,127.0.0.2",
			},
			[]string{
				"127.0.0.1",
				"127.0.0.2",
			},
		},
		{
			map[string]string{
				"IP_PROVIDER_ENV_IPS": "127.0.0.1,127.0.0.2 ",
			},
			[]string{
				"127.0.0.1",
				"127.0.0.2",
			},
		},
		{
			map[string]string{
				"IP_PROVIDER_ENV_IPS": "invalid",
			},
			nil,
		},
	}

	for _, tc := range testCases {
		for k, v := range tc.envVars {
			_ = os.Setenv(k, v)
		}
		ips, _ := envIpProvider.GetAvailableIps(context.TODO())

		if !reflect.DeepEqual(ips, tc.ips) {
			t.Fatalf("Expected ips to be %v, but got %v", tc.ips, ips)
		}
	}
}
