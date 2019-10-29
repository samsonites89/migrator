package blockchain

// DeliveryOrder represents a delivery order document as provided by cello.
// TODO: This will also not be part of the geofence model.
type DeliveryOrder struct {
	Data struct {
		Containers []struct {
			Info struct {
				Number string `xml:"Number"`
			} `xml:"ContainerInformation"`
		} `xml:"CntrItem"`
	} `xml:"CntrList"`
}
