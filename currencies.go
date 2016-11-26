package kiwitaxi

type Currency struct {
	Id   int    `csv:"id" json:"id" bson:"id"`
	Code string `csv:"code" json:"code" bson:"code"`
	Iso  int    `csv:"iso" json:"iso" bson:"iso"`
}

func (a *KiwitaxiApi) Currencies() (currencies []*Currency, err error) {
	err = a.getCsv("services/data/csv/currencies", &currencies)
	return
}
