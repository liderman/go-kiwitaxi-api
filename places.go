package kiwitaxi

type Place struct {
	Id           int    `csv:"id" json:"id" bson:"id"`
	CountryId    int    `csv:"country_id" json:"country_id" bson:"country_id"`
	RegionId     int    `csv:"region_id" json:"region_id" bson:"region_id"`
	TypeId       int    `csv:"type_id" json:"type_id" bson:"type_id"`
	NameEn       string `csv:"name_en" json:"name_en" bson:"name_en"`
	NameRu       string `csv:"name_ru" json:"name_ru" bson:"name_ru"`
	Iata         string `csv:"iata" json:"iata" bson:"iata"`
	PlacePolygon string `csv:"place_polygon" json:"place_polygon" bson:"place_polygon"`
}

func (a *KiwitaxiApi) Places() (places []*Place, err error) {
	err = a.getCsv("services/data/csv/places", &places)
	return
}
