package kiwitaxi

type Country struct {
	Id         int    `csv:"id" json:"id" bson:"id"`
	Iata       string `csv:"iata" json:"iata" bson:"iata"`
	Published  int    `csv:"published" json:"published" bson:"published"`
	NameEn     string `csv:"name_en" json:"name_en" bson:"name_en"`
	NameRu     string `csv:"name_ru" json:"name_ru" bson:"name_ru"`
	CurrencyId int    `csv:"currency_id" json:"currency_id" bson:"currency_id"`
	TimeZone   string `csv:"time_zone" json:"time_zone" bson:"time_zone"`
}

func (a *KiwitaxiApi) Countries() (countries []*Country, err error) {
	err = a.getCsv("services/data/csv/countries", &countries)
	return
}
