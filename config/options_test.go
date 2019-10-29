package config

import (
	"testing"
)

func TestDefault(t *testing.T) {
	if ApiKey != "MDPT8U7JZ2DB54PXV59TPVQQZ94QJSB5PN" {
		t.Fail()
	}
	if DefaultKey.Password != "rinkeby" {
		t.Fail()
	}

}
