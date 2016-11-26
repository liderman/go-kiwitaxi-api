package kiwitaxi

type Transfer struct {
	Id       int     `csv:"id"`
	RouteId  int     `csv:"route_id"`
	TypeId   int     `csv:"type_id"`
	PriceRub float64 `csv:"price_rub"`
	PriceEur float64 `csv:"price_eur"`
	PriceUsd float64 `csv:"price_usd"`
	Url      string  `csv:"url"`
}

func (a *KiwitaxiApi) Transfers(paymentType string) (transfers []*Transfer, err error) {
	params := map[string]string{}
	if paymentType != "" {
		params["payment_type"] = paymentType
	}

	err = a.getCsvWithParams("services/data/csv/transfers", params, &transfers)
	return
}
