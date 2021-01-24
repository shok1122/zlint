package main

import "testing"

func TestRun(t *testing.T) {
	results, err := run("maintestdata")
	if err != nil {
		t.Error(err)
		return
	}
	if len(results) != 1 {
		t.Errorf("expected 1 error, got %d", len(results))
	}
}
