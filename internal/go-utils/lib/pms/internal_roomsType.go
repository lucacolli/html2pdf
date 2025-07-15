package pms

import (
	"encoding/json"

	"github.com/otelia/hotc/lib/auth"
	"github.com/otelia/hotc/lib/utils"
)

type RoomType struct {
	AllestimentoDefault string `json:"allestimentoDefault"`
	Codice              string `json:"codice"`
	Descrizione         string `json:"descrizione"`
	EqAllotment         string `json:"eqAllotment"`
	EqDisponibilita     string `json:"eqDisponibilita"`
	EqStatistica        string `json:"eqStatistica"`
	EqTariffa           string `json:"eqTariffa"`
	NumeroLetti         int    `json:"numeroLetti"`
	NumeroPersone       int    `json:"numeroPersone"`
	Ordine              int    `json:"ordine"`
	QuanteCamere        int    `json:"quanteCamere"`
	TipoStruttura       string `json:"tipoStruttura"`
}

func (p *Pms) ListRoomTypes() ([]RoomType, error) {
	var items []RoomType
	token, err := auth.GetToken()
	if err != nil {
		return nil, err
	}
	url := utils.GetConfig("PMS_URL") + "/" + p.ID + "/api/resources/roomtypes/v1"

	r, err := utils.Request("GET", url, []byte{}, token)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(r, &items)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (p *Pms) GetRoomType(id string) (*RoomType, error) {
	var item RoomType
	token, err := auth.GetToken()
	if err != nil {
		return nil, err
	}
	url := utils.GetConfig("PMS_URL") + "/" + p.ID + "/api/resources/roomtypes/v1/" + id

	r, err := utils.Request("GET", url, []byte{}, token)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(r, &item)
	if err != nil {
		return nil, err
	}
	return &item, nil
}
