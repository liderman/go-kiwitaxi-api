package kiwitaxi

type PlaceType struct {
	Id     int    `csv:"id" json:"id" bson:"id"`
	NameEn string `csv:"name_en" json:"name_en" bson:"name_en"`
	NameRu string `csv:"name_ru" json:"name_ru" bson:"name_ru"`
}

func (a *KiwitaxiApi) PlaceTypes() (placeTypes []*PlaceType, err error) {
	err = a.getCsv("services/data/csv/place_types", &placeTypes)
	return
}
