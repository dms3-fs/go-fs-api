package shell

import (
	"fmt"
	"testing"
	"time"
)

var examplesHashForDMS3NS = "/dms3fs/Qmbu7x6gJbsKDcseQv66pSbUcAA3Au6f7MfTYVXwvBxN2K"
var testKey = "self" // feel free to change to whatever key you have locally

func TestPublishDetailsWithKey(t *testing.T) {
	t.Skip()
	shell := NewShell("localhost:5001")

	resp, err := shell.PublishWithDetails(examplesHashForDMS3NS, testKey, time.Second, time.Second, false)
	if err != nil {
		t.Fatal(err)
	}

	if resp.Value != examplesHashForDMS3NS {
		t.Fatalf(fmt.Sprintf("Expected to receive %s but got %s", examplesHash, resp.Value))
	}
}

func TestPublishDetailsWithoutKey(t *testing.T) {
	t.Skip()
	shell := NewShell("localhost:5001")

	resp, err := shell.PublishWithDetails(examplesHashForDMS3NS, "", time.Second, time.Second, false)
	if err != nil {
		t.Fatal(err)
	}

	if resp.Value != examplesHashForDMS3NS {
		t.Fatalf(fmt.Sprintf("Expected to receive %s but got %s", examplesHash, resp.Value))
	}
}
