package dns

import (
	"context"
	"testing"
)

func TestResolverDial(t *testing.T) {
	dial, err := resolverDial(context.Background(), "ip4", "localhost")
	if err != nil {
		t.Fatal(err.Error())
	}
	_ = dial.Close()
}
