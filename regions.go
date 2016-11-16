package kiwitaxi

type Region struct {
	Id        int    `csv:"id"`
	CountryId int    `csv:"country_id"`
	NameEn    string `csv:"name_en"`
	NameRu    string `csv:"name_ru"`
}

func (a *KiwitaxiApi) Regions() (regions []*Region, err error) {
	err = a.getCsv("services/data/csv/regions", &regions)
	return
}
