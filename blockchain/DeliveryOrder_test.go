package blockchain_test

import (
	"encoding/xml"
	"os"
	"testing"

	"github.blockchaindltlab.nl/deliver-ng/deliver-oracle-geofencing/blockchain"
)

func TestDeliveryOrderXMLParse(t *testing.T) {

	raw, err := os.Open("deliver_order_test.xml")
	if err != nil {
		t.Fatalf("could not open file: %v", err)
	}
	defer raw.Close()

	var (
		dec = xml.NewDecoder(raw)
		doc blockchain.DeliveryOrder
	)

	if err := dec.Decode(&doc); err != nil {
		t.Fatalf("could not parse delivery order from xml (err=%s)", err)
	}

	if len(doc.Data.Containers) != 2 {
		t.Errorf("got %d containers, expected %d", len(doc.Data.Containers), 2)
	}
}
