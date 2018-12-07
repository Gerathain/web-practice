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
func TestGetApiKeyMissingFile(t *testing.T) {
	correctKey := "YaX2vXeBA3W-w6sNNwpkebTheJjAVPE0tgCQ86NG"

	retVal := getApiKey( "fileDoesntExist" )
	if correctKey != retVal {
		t.Fail()
	}
}

func TestReadConfig(t *testing.T) {
	controllerOn := true
	startHour := 21
	startMin := 0
	warmRate := 50
	correctConfig := Config{ &controllerOn, &startHour, &startMin, &warmRate }

	var retConfig Config
	err := readConfig( "testConfig", &retConfig )
	if err != nil {
		t.Fail()
		return
	}
	
	// This is needed to cover all the values as the members are all pointers
	if *(retConfig.ControllerOn) != *(correctConfig.ControllerOn) {
		t.Fail()
	}
	if *(retConfig.StartHour) != *(correctConfig.StartHour) {
		t.Fail()
	}
	if *(retConfig.StartMinute) != *(correctConfig.StartMinute) {
		t.Fail()
	}
	if *(retConfig.WarmingRate) != *(correctConfig.WarmingRate) {
		t.Fail()
	}
}
func TestReadConfigMissingFile(t *testing.T) {
	err := readConfig( "fileDoesntExist", nil )
	if err != nil {
		t.Fail()
	}
}
func TestReadConfigNilConfig(t *testing.T) {
	err := readConfig( "testConfig", nil )
	if err != nil {
		t.Fail()
	}
}
