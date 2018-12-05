package main

import (
	"testing"
)

func TestGetBridgeUrl(t *testing.T) {
	correctUrl := "http://192.168.0.17/api/"

	retVal := getBridgeUrl()
	if correctUrl != retVal {
		t.Fail()
	}
}
