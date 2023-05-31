package blocklist_contributor

import (
	"reflect"
	"testing"
)

func testPriority(c BlocklistContributor, t *testing.T, expected int) {
	actual := c.Priority()
	if actual != expected {
		t.Fatalf("Expected priority to be %d, but got %d", expected, actual)
	}
}

func testContribute(c BlocklistContributor, t *testing.T, before []string, expected []string) {
	actual := c.Contribute(before)
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Expected contributed list to be %v, but got %v", expected, actual)
	}
}

func testProviderName(p BlocklistContributor, t *testing.T, expectedName string) {
	actualName := p.Name()
	if actualName != expectedName {
		t.Fatalf("Expected provider name to be '%s' but got '%s'", expectedName, actualName)
	}
}

func TestDeduplicate(t *testing.T) {
	dedup := Deduplicate([]string{"a", "a", "b"})
	expected := []string{"a", "b"}

	if !reflect.DeepEqual(dedup, expected) {
		t.Fatalf("Expected %v but got %v - duplicates seem to still exist", expected, dedup)
	}
}

func TestAggregateBlocklist(t *testing.T) {
	res := AggregateBlocklist()
	if !reflect.DeepEqual(res, DefaultBlocklists) {
		t.Fatalf("Expected only default blocklists to be present")
	}
}
