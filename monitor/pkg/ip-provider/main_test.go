package ip_provider

import (
	"context"
	"testing"
)

func testProviderName(p IpProvider, t *testing.T, expectedName string) {
	actualName := p.Name()
	if actualName != expectedName {
		t.Fatalf("Expected provider name to be '%s' but got '%s'", expectedName, actualName)
	}
}

func testShouldFailSoft(p IpProvider, t *testing.T, expected bool) {
	actual := p.ShouldFailSoft()
	if actual != expected {
		t.Fatalf("Expected ShouldFailt to be %t, but got %t", expected, actual)
	}
}

func TestLoadAllIps(t *testing.T) {
	_ = LoadAllIps(context.TODO())
}
