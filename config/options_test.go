package config

import (
	"fmt"
	"os"
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

func TestArgs(t *testing.T) {

	fmt.Println(os.Args[1])

}
