package models

type LimitKonsumen struct {
	ID         int `json:"id"`
	KonsumenId int `json:"konsumen_id"`
	Tenor1     int `json:"tenor_1"`
	Tenor2     int `json:"tenor_2"`
	Tenor3     int `json:"tenor_3"`
	Tenor4     int `json:"tenor_4"`
}
