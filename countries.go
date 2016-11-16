package kiwitaxi

type Country struct {
	Id         int    `csv:"id"`
	Iata       string `csv:"iata"`
	Published  int    `csv:"published"`
	NameEn     string `csv:"name_en"`
	NameRu     string `csv:"name_ru"`
	CurrencyId int    `csv:"currency_id"`
	TimeZone   string `csv:"time_zone"`
}

func (a *KiwitaxiApi) Countries() (countries []*Country, err error) {
	err = a.getCsv("services/data/csv/countries", &countries)
	return
}
