package ee

import "testing"

func TestAddonName(t *testing.T) {
	if got := New().Name(); got != "ee" {
		t.Fatalf("Name() = %q, want ee", got)
	}
}
