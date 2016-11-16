package kiwitaxi

import (
	"fmt"
	"testing"
)

func TestUrlDomain(t *testing.T) {
	api := NewKiwitaxiApi("24e8a1c890fb0d16ecdce4926dc4b2d6")
	data, err := api.Places()
	if err != nil {
		t.Errorf("%s", err.Error())
	}

	fmt.Printf("Count %d\n", len(data))
	for _, v := range data {
		fmt.Printf("%d:%s:%s:%s:%d:%d:%d\n", v.Id, v.NameEn, v.NameRu, v.Iata, v.CountryId, v.RegionId, v.TypeId)
	}

}
