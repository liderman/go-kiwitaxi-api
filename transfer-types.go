package kiwitaxi

type TransferType struct {
	Id            int    `csv:"id"`
	NameEn        string `csv:"name_en"`
	NameRu        string `csv:"name_ru"`
	Pax           int    `csv:"pax"`
	Baggage       int    `csv:"baggage"`
	DescriptionEn string `csv:"description_en"`
	DescriptionRu string `csv:"description_ru"`
	Photo         string `csv:"photo"`
	Sortno        int    `csv:"sortno"`
}

func (a *KiwitaxiApi) TransferTypes() (transferTypes []*TransferType, err error) {
	err = a.getCsv("services/data/csv/transfer_types", &transferTypes)
	return
}
