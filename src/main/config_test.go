package main

import (
	"confdReWrite/src/log"
	"testing"
)

func TestInitConfig(t *testing.T){

	log.SetLevel("warn")

	if err := initConfig(); err != nil {
		t.Errorf(err.Error())
	}


}