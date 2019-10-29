package abi

import (
	"encoding/hex"
	"fmt"
	"log"
	"testing"
)

func TestParseInput(t *testing.T) {

	Init()
	textData := "0xba377a26da7ab8c89d4b0fe5008c3a6d474cc345f9a0292a7f36bc01b28be3856be5cb6d00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000060000000000000000000000000000000000000000000000000000000000000000d424b323031393038333030303100000000000000000000000000000000000000"
	// decode txInput method signature
	byteData, err := hex.DecodeString(textData[2:10])
	if err != nil {
		t.Fail()
	}

	m, err := customerJourneyABI.MethodById(byteData)
	if err != nil {
		// (probably) non-notary contract
		t.Fail()
	} else {
		fmt.Println(m.Name)
		// decode txInput Payload
		decodedData, err := hex.DecodeString(textData[10:])
		if err != nil {
			log.Fatal(err)
		}
		var data CustomerJourneyStart
		err = m.Inputs.Unpack(&data, decodedData)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(data.Hash.Hex())

	}

	textData = "0xd16e024eda7ab8c89d4b0fe5008c3a6d474cc345f9a0292a7f36bc01b28be3856be5cb6d8964b08ab2c7a94cf4b2d23948eb6d8e4a290da7bb5f5dbab6c4ed6d04c5b80100000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000000a3630303339373035393500000000000000000000000000000000000000000000"
	// decode txInput method signature
	byteData, err = hex.DecodeString(textData[2:10])
	if err != nil {
		t.Fail()
	}
	m, err = customerJourneyABI.MethodById(byteData)
	if err != nil {
		// (probably) non-notary contract
		t.Fail()
	} else {
		fmt.Println(m)
		// decode txInput Payload
		decodedData, err := hex.DecodeString(textData[10:])
		if err != nil {
			log.Fatal(err)
		}
		var data CustomerJourneyLink
		err = m.Inputs.Unpack(&data, decodedData)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(data.Hash.Hex())
		fmt.Println(data.Parent.Hex())
		fmt.Println(data.Typ)

	}

}

func TestDetermineType(t *testing.T) {
	Init()
	textData := "0xba377a26da7ab8c89d4b0fe5008c3a6d474cc345f9a0292a7f36bc01b28be3856be5cb6d00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000060000000000000000000000000000000000000000000000000000000000000000d424b323031393038333030303100000000000000000000000000000000000000"
	typ, _ := DetermineType(textData)
	fmt.Println(typ)
}
