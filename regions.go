package kiwitaxi

type Region struct {
	Id        int    `csv:"id" json:"id" bson:"id"`
	CountryId int    `csv:"country_id" json:"country_id" bson:"country_id"`
	NameEn    string `csv:"name_en" json:"name_en" bson:"name_en"`
	NameRu    string `csv:"name_ru" json:"name_ru" bson:"name_ru"`
}

func (a *KiwitaxiApi) Regions() (regions []*Region, err error) {
	err = a.getCsv("services/data/csv/regions", &regions)
	return
}
