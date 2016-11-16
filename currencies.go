package kiwitaxi

type Currency struct {
	Id   int    `csv:"id"`
	Code string `csv:"code"`
	Iso  int    `csv:"iso"`
}

func (a *KiwitaxiApi) Currencies() (currencies []*Currency, err error) {
	err = a.getCsv("services/data/csv/currencies", &currencies)
	return
}
