package kiwitaxi

type PlaceType struct {
	Id     int    `csv:"id"`
	NameEn string `csv:"name_en"`
	NameRu string `csv:"name_ru"`
}

func (a *KiwitaxiApi) PlaceTypes() (placeTypes []*PlaceType, err error) {
	err = a.getCsv("services/data/csv/place_types", &placeTypes)
	return
}
