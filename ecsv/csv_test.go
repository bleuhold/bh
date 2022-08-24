package ecsv

import (
	"testing"
)

func TestCSV_ReadData(t *testing.T) {
	// this would be a similar output as when a file is read from the filesystem.
	xb := []byte("Account Name:Mr X Test  10012811234,\nTransaction Date,Posting Date,Description,Debits,Credits,Balance\n2021/12/27,2021/12/30,SHELL ULTRA CITY BUTTEBUTTERWORTHZA,769.65,,-25994.17,\n2021/12/27,2021/12/30,DEBONAIRSIBIKAZA,105.9,,-25224.52,\n2021/12/27,2021/12/30,FRESH LAUNDRYPORT ELIZABETZA,313.9,,-25118.62,\n2021/12/28,2021/12/29,APPLE.COM/BILLITUNES.COMIE,14.99,,-24804.72,\n2021/12/28,2021/12/29,GOOGLE STORAGELONDONGB,29.0,,-24789.73,\n2021/12/24,2021/12/29,FRANCIS PHARMACYST FRANCIS BAZA,164.8,,-24760.73,\n2021/12/27,2021/12/29,NANAGA EXPRESSALEXANDRIAZA,801.15,,-24595.93,\n2021/12/27,2021/12/28,HTTP://WWW.MYISTORE.COSANDTON  JOZA,9999.0,,-23794.78,\n2021/12/27,2021/12/28,SPAR EXPRESS BUTTERWORBUTTERWORTHZA,101.97,,-13795.78,\n2021/12/26,2021/12/27,REFINERY JEFFREYS BAYECZA,1348.0,,4306.19,\n2021/12/26,2021/12/27,WOOLWORTHS JEFFREYS BAJEFFREYS BAYZA,857.9,,5654.19,\n2021/12/24,2021/12/25,SUPERSPAR VILLAGE SQUEASTERN CAPEZA,108.47,,6512.09,\n2021/12/24,2021/12/25,NEVERMIND RESTAURANTCAPE ST FRANCZA,191.0,,6620.56,")

	c := CSV{StartOffset: 1}
	c.ReadData(xb)

	if len(c.Records) != 14 {
		t.Errorf("expected records to have length %d got %d", 14, len(c.Records))
	}
	// this should be the columns header
	rec1 := c.Records[0]
	if rec1[0] != "Transaction Date" {
		t.Errorf("expected record 1 element 1 to be '%s' got '%s'", "Transaction Date", rec1[0])
	}
}
