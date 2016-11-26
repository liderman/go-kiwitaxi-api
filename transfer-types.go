package kiwitaxi

type TransferType struct {
	Id            int    `csv:"id" json:"id" bson:"id"`
	NameEn        string `csv:"name_en" json:"name_en" bson:"name_en"`
	NameRu        string `csv:"name_ru" json:"name_ru" bson:"name_ru"`
	Pax           int    `csv:"pax" json:"pax" bson:"pax"`
	Baggage       int    `csv:"baggage" json:"baggage" bson:"baggage"`
	DescriptionEn string `csv:"description_en" json:"description_en" bson:"description_en"`
	DescriptionRu string `csv:"description_ru" json:"description_ru" bson:"description_ru"`
	Photo         string `csv:"photo" json:"photo" bson:"photo"`
	Sortno        int    `csv:"sortno" json:"sortno" bson:"sortno"`
}

func (a *KiwitaxiApi) TransferTypes() (transferTypes []*TransferType, err error) {
	err = a.getCsv("services/data/csv/transfer_types", &transferTypes)
	return
}
