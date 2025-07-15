package billing

import (
	"encoding/json"
	"time"

	"github.com/otelia/hotc/lib/utils"
)

type Invoice struct {
	ID               string    `json:"id"`
	CreatedOn        time.Time `json:"created_on"`
	InvoiceNumber    int       `json:"invoice_number"`
	InvoiceNumbering string    `json:"invoice_numbering"`
	CreditNote       bool      `json:"credit_note"`
	InvoiceDate      time.Time `json:"invoice_date"`
	PeriodFrom       time.Time `json:"period_from"`
	PeriodTo         time.Time `json:"period_to"`
	Currency         string    `json:"currency"`
	Status           int       `json:"status"`
	Exported         bool      `json:"exported"`
	DDT              string    `json:"ddt"`
	Notes            string    `json:"notes"`
	Payments         []struct {
		AmountDec string    `json:"amount"`
		Date      time.Time `json:"date"`
	} `json:"payments"`
	Instalments []struct {
		AmountDec string    `json:"amount"`
		DueDate   time.Time `json:"due_date"`
	} `json:"instalments"`
	PaymentMethod  string `json:"payment_method"`
	PaymentDetail1 string `json:"payment_detail_1"`
	PaymentDetail2 string `json:"payment_detail_2"`
	Recipient      struct {
		ID            string  `json:"id"`
		Name          string  `json:"name"`
		CompanyNumber string  `json:"company_number"`
		VATCode       string  `json:"vat_code"`
		Address       Address `json:"address"`
	} `json:"recipient"`
	RecipientID string `json:"recipient_id"`
	VATFree     string `json:"vat_free"`
	Sender      struct {
		ID            string  `json:"id"`
		Name          string  `json:"name"`
		CompanyNumber string  `json:"company_number"`
		VATCode       string  `json:"vat_code"`
		Address       Address `json:"address"`
	} `json:"sender"`
	SenderID string `json:"sender_id"`
	Lines    []struct {
		RelatedOrderID    string  `json:"related_order_id"`
		ReasonDescription string  `json:"reason_description"`
		OrderDescription  string  `json:"order_description"`
		SkuDescription    string  `json:"sku_description"`
		Address           Address `json:"address"`
		Product           string  `json:"product"`
		SKU               string  `json:"sku"`
		Quantity          int     `json:"quantity"`
		ListedPriceDec    string  `json:"listed_price"`
		UnitPriceDec      string  `json:"unit_price"`
		SubTotalDec       string  `json:"subtotal"`
		VatPercentage     int     `json:"vat_percentage"`
		VatAmountDec      string  `json:"vat_amount"`
		TotalDec          string  `json:"total"`
	} `json:"lines"`
	SubTotalDec string `json:"subtotal"`
	Vats        []struct {
		Percentage  int    `json:"percentage"`
		SubTotalDec string `json:"subtotal"`
		AmountDec   string `json:"amount"`
	} `json:"vats"`
	TotalDec string `json:"total"`
}

func GetInvoices(token string) ([]Invoice, error) {
	url := utils.GetConfig("API_URL") + "/billing/v0/invoices?limit=1000"

	r, err := utils.Request("GET", url, []byte{}, token)

	var out []Invoice
	if err != nil {
		return out, err
	}
	err = json.Unmarshal(r, &out)
	return out, err
}
