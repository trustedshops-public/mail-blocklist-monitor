package ip_provider

import "testing"

var sesProvider = SesIpProvider{}

func TestSesIpProvider_Name(t *testing.T) {
	testProviderName(sesProvider, t, "ses")
}

func TestSesIpProvider_ShouldFailSoft(t *testing.T) {
	testShouldFailSoft(sesProvider, t, true)
}
