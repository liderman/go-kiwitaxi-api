package kiwitaxi

type Route struct {
	Id          int     `csv:"id" json:"id" bson:"id"`
	CountryId   int     `csv:"country_id" json:"country_id" bson:"country_id"`
	PlaceFromId int     `csv:"place_from_id" json:"place_from_id" bson:"place_from_id"`
	PlaceToId   int     `csv:"place_to_id" json:"place_to_id" bson:"place_to_id"`
	Distance    int     `csv:"distance" json:"distance" bson:"distance"`
	Timeinway   int     `csv:"timeinway" json:"timeinway" bson:"timeinway"`
	Weight      float64 `csv:"weight" json:"weight" bson:"weight"`
	Url         string  `csv:"url" json:"url" bson:"url"`
}

func (a *KiwitaxiApi) Routes() (routes []*Route, err error) {
	err = a.getCsv("services/data/csv/routes", &routes)
	return
}
