package kiwitaxi

type Route struct {
	Id          int     `csv:"id"`
	CountryId   int     `csv:"country_id"`
	PlaceFromId int     `csv:"place_from_id"`
	PlaceToId   int     `csv:"place_to_id"`
	Distance    int     `csv:"distance"`
	Timeinway   int     `csv:"timeinway"`
	Weight      float64 `csv:"weight"`
	Url         string  `csv:"url"`
}

func (a *KiwitaxiApi) Routes() (routes []*Route, err error) {
	err = a.getCsv("services/data/csv/routes", &routes)
	return
}
