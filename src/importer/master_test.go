package importer

import "testing"
import "time"
import "reflect"

// func Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location) Time

func TestParseLine(t *testing.T) {
	expected := MasterItem{
		StockCode:            "STL065506",
		FullTitle:            "OCT172989",
		Price:                12,
		Publisher:            "WIZKIDS/NECA",
		UPCNumber:            "634482708903",
		EANNumber:            "",
		ShippingDate:         time.Date(2018, time.February, 28, 0, 0, 0, 0, time.UTC),
		SuggestedRetailPrice: 0,
		Category:             5,
		Writer:               "",
		Artist:               "",
		CoverArtist:          "",
		FinalOrderCutoff:     time.Date(2018, time.November, 3, 0, 0, 0, 0, time.UTC),
	}

	source := "OCT172989	STL065506			DC HEROCLIX TRINITY MONTHLY OP KIT (Net) (C: 1-1-2)	DC HEROCLIX TRINITY MONTHLY OP KIT		0		0			12	WIZKIDS/NECA	634482708903			0	0	0	J	1	11/3/2017	N	2/28/2018	0	5	SH	WK	N	N	N	1	1	2	N	N	(NOTE: Advance solicited for February 2018 on-sale. Limit 3 per store location.)	591					WZK 70890	11/3/2017"
	result := parseMasterLine(source)

	if !reflect.DeepEqual(expected, result) {
		t.Error("For source", source, "expected", expected, "got", result)
	}
}
