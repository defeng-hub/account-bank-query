package dto

type BankCode string

func (b BankCode) Prefix(k int) string {
	if k >= len(b) {
		return string(b)
	}
	return string(b[0:k])
}

type BankCodeDto struct {
	BankCode  BankCode   `json:"bank_code"`
	BankCodes []BankCode `json:"bank_codes"`
}
