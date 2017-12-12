package importer

import "time"

// MasterItem represents an Item in the master catalog
type MasterItem struct {
	StockCode            string
	MonthCode            string
	FullTitle            string
	Price                float64
	Publisher            string
	UPCNumber            string
	EANNumber            string
	ShippingDate         time.Time
	SuggestedRetailPrice float64
	Category             int32
	Writer               string
	Artist               string
	CoverArtist          string
	FinalOrderCutoff     time.Time
}

func parseMasterLine(s string) MasterItem {
	return MasterItem{}
}
