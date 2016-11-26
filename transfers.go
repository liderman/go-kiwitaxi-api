package kiwitaxi

type Transfer struct {
	Id       int     `csv:"id" json:"id" bson:"id"`
	RouteId  int     `csv:"route_id" json:"route_id" bson:"route_id"`
	TypeId   int     `csv:"type_id" json:"type_id" bson:"type_id"`
	PriceRub float64 `csv:"price_rub" json:"price_rub" bson:"price_rub"`
	PriceEur float64 `csv:"price_eur" json:"price_eur" bson:"price_eur"`
	PriceUsd float64 `csv:"price_usd" json:"price_usd" bson:"price_usd"`
	Url      string  `csv:"url" json:"url" bson:"url"`
}

func (a *KiwitaxiApi) Transfers(paymentType string) (transfers []*Transfer, err error) {
	params := map[string]string{}
	if paymentType != "" {
		params["payment_type"] = paymentType
	}

	err = a.getCsvWithParams("services/data/csv/transfers", params, &transfers)
	return
}
