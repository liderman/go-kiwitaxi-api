package kiwitaxi

type Place struct {
	Id           int    `csv:"id"`
	CountryId    int    `csv:"country_id"`
	RegionId     int    `csv:"region_id"`
	TypeId       int    `csv:"type_id"`
	NameEn       string `csv:"name_en"`
	NameRu       string `csv:"name_ru"`
	Iata         string `csv:"iata"`
	PlacePolygon string `csv:"place_polygon"`
}

func (a *KiwitaxiApi) Places() (places []*Place, err error) {
	err = a.getCsv("services/data/csv/places", &places)
	return
}
