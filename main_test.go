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

func TestGetApiKey(t *testing.T) {
	correctKey := "YaX2vXeBA3W-w6sNNwpkebTheJjAVPE0tgCQ86NG"

	retVal := getApiKey( "apiKey" )
	if correctKey != retVal {
		t.Fail()
	}
}
