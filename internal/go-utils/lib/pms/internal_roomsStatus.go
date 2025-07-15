package pms

import (
	"encoding/json"
	"time"

	"github.com/otelia/hotc/lib/auth"
	"github.com/otelia/hotc/lib/utils"
)

type RoomStatus struct {
	Data                  time.Time `json:"data"`
	TipoCamera            string    `json:"tipoCamera"`
	GiornoDiChiusura      bool      `json:"giornoDiChiusura"`
	NumeroCamere          int       `json:"numeroCamere"`
	NumeroCamereBloccate  int       `json:"numeroCamereBloccate"`
	NumeroCamereInagibili int       `json:"numeroCamereInagibili"`
	NumeroCamereVendibili int       `json:"numeroCamereVendibili"`
}

func (u *Unit) RoomStatus(from time.Time, to time.Time) ([]RoomStatus, error) {
	var items []RoomStatus
	token, err := auth.GetToken()
	if err != nil {
		return nil, err
	}
	url := utils.GetConfig("PMS_URL") + "/" + u.Pms.ID + "/api/dataextraction/roomtypestatus/v1?unit=" + u.ID + "from_date=" + from.Format("2006-01-02") + "&to_date=" + to.Format("2006-01-02")

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
