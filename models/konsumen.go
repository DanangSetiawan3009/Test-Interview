package models

import "time"

type Konsumen struct {
	ID           int        `json:"id"`
	Nik          string     `json:"nik"`
	Fullname     string     `json:"fullname"`
	Legalname    string     `json:"legalname"`
	TempatLahir  string     `json:"tempat_lahir"`
	TanggalLahir *time.Time `json:"tanggal_lahir"`
	Gaji         int        `json:"gaji"`
	FotoKtp      string     `json:"foto_ktp"`
	FotoSelfie   string     `json:"foto_selfie"`
}
