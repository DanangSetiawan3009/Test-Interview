package models

type Transaksi struct {
	ID              int    `json:"id"`
	LimitKonsumenId int    `json:"limit_konsumen_id"`
	NoKontrak       string `json:"no_kontrak"`
	Otr             int    `json:"otr"`
	AdminFee        int    `json:"admin_fee"`
	JumlahCicilan   int    `json:"jumlah_cicilan"`
	JumlahBunga     int    `json:"jumlah_bunga"`
	NamaAsset       string `json:"nama_asset"`
}
