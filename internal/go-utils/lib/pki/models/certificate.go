package models

type Certificate struct {
	ID          string `json:"id" sql:"primary key" groups:"rP,rX,lP"`
	Serial      int    `json:"serial_number" groups:"rP,rX,lP"`
	Host        string `json:"host" groups:"cP,rP,rX,lP"`
	Key         string `json:"key" groups:"rP,rX,lP"`
	Certificate string `json:"certificate" groups:"rP,rX,lP"`
	CACrt       string `json:"ca_certificate" sql:"column:ca_crt" groups:"rP,rX,lP"`
}
