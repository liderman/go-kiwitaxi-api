package kiwitaxi

type UrlDomain struct {
	Id     int    `csv:"id" json:"id" bson:"id"`
	Locale string `csv:"locale" json:"locale" bson:"locale"`
	Value  string `csv:"value" json:"value" bson:"value"`
}

func (a *KiwitaxiApi) UrlDomains() (urlDomains []*UrlDomain, err error) {
	err = a.getCsv("services/data/csv/url_domains", &urlDomains)
	return
}
