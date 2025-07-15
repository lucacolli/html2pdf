package pms

import (
	"encoding/json"

	"github.com/otelia/hotc/lib/auth"
	"github.com/otelia/hotc/lib/utils"
)

type Unit struct {
	AlloggiatiWebType                        string   `json:"alloggiatiWebType"`
	BookingEmail                             string   `json:"bookingEmail"`
	CityLedgerEmail                          string   `json:"cityLedgerEmail"`
	CompanyAddress                           string   `json:"companyAddress"`
	CompanyCity                              string   `json:"companyCity"`
	CompanyCountry                           string   `json:"companyCounty"`
	CompanyName                              string   `json:"companyName"`
	CompanyPostalCode                        string   `json:"companyPostalCode"`
	CompanyTaxCode                           string   `json:"companyTaxCode"`
	CompanyVatNumber                         string   `json:"companyVatNumber"`
	DescriptionButtonManagementRelationship  string   `json:"descriptionButtonManagementRelationship"`
	DocumentTypeInvoiceManagementRelationhip string   `json:"documentTypeInvoiceManagementRelationhip"`
	DocumentTypeReceiptManagementRelationhip string   `json:"documentTypeReceiptManagementRelationhip"`
	HotelAddress                             string   `json:"hotelAddress"`
	HotelCity                                string   `json:"hotelCity"`
	HotelCountry                             string   `json:"hotelCounty"`
	HotelName                                string   `json:"hotelName"`
	HotelPostalCode                          string   `json:"hotelPostCode"`
	HotelRatings                             string   `json:"hotelRatings"`
	HotelTypology                            string   `json:"hotelTypology"`
	ID                                       string   `json:"id"`
	IstatSubmissionCode                      string   `json:"istatSubmissionCode"`
	LastNumberAccountingRecords              int      `json:"lastNumberAccountingRecords"`
	LastNumberArrivalsRegister               int      `json:"lastNumberArrivalsRegister"`
	LastNumberC59                            int      `json:"lastNumberC59"`
	LastNumberC59G                           int      `json:"lastNumberC59G"`
	LastNumberC60                            int      `json:"lastNumberC60"`
	LastNumberFatturaPA                      int      `json:"lastNumberFatturaPA"`
	MarketingEmail                           string   `json:"marketingEmail"`
	MasterUnitFiscal                         string   `json:"masterUnitFiscal"`
	MasterUnitIstat                          string   `json:"masterUnitIstat"`
	MasterUnitQuestura                       string   `json:"masterUnitQuestura"`
	Order                                    string   `json:"order"`
	RemarksOnPendingPaymentReceipt           []string `json:"remarksOnPendingPaymentReceipt"`
	RemarksOnReceipt                         []string `json:"remarksOnReceipt"`
	UnderManagement                          bool     `json:"underManagement"`
	Pms                                      *Pms
}

func (p *Pms) ListUnits() ([]Unit, error) {
	var items []Unit
	token, err := auth.GetToken()
	if err != nil {
		return nil, err
	}
	url := utils.GetConfig("PMS_URL") + "/" + p.ID + "/api/resources/units/v1"

	r, err := utils.Request("GET", url, []byte{}, token)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(r, &items)
	if err != nil {
		return nil, err
	}
	for k := range items {
		items[k].Pms = p
	}
	return items, nil
}

func (p *Pms) GetUnit(id string) (*Unit, error) {
	var item Unit
	token, err := auth.GetToken()
	if err != nil {
		return nil, err
	}
	url := utils.GetConfig("PMS_URL") + "/" + p.ID + "/api/resources/units/v1/" + id

	r, err := utils.Request("GET", url, []byte{}, token)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(r, &item)
	if err != nil {
		return nil, err
	}
	item.Pms = p
	return &item, nil
}
