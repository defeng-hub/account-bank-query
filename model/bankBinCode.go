package model

// BankBinCode
// 银行卡号(KH),行别,简写,长度
// 622974,AEON信贷财务,AEON,6
// 620011,BC卡公司,BC卡公,6
// 622928,工商银行,工商银行,6
type BankBinCode struct {
	//银行卡号(KH),行别,简写,长度
	BankCode       string `json:"-"`
	BankName       string `json:"bank_name"`
	BankShortName  string `json:"bank_short_name"`
	BankCardLength int64  `json:"-"`
}
