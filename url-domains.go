package kiwitaxi

type UrlDomain struct {
	Id     int    `csv:"id"`
	Locale string `csv:"locale"`
	Value  string `csv:"value"`
}

func (a *KiwitaxiApi) UrlDomains() (urlDomains []*UrlDomain, err error) {
	err = a.getCsv("services/data/csv/url_domains", &urlDomains)
	return
}
