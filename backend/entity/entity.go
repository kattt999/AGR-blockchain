package entity

type KlineData struct {
	Date string
	Data [4]float64//belh
}
type KlineTemp struct {
	Kline KlineData
	Btime string
	Etime string
}
type Article_old struct {
	//	Funda string     `json:"Funda"`
	PH [30]float64 `json:"PH"`
	L  [30]float64 `json:"L"`
	R  [30]float64 `json:"R"`
	SH [30]float64 `json:"SH"`
	ST [30]float64 `json:"ST"`
	AH [30]float64 `json:"AH"`
	AT [30]float64 `json:"AT"`
}
type Article struct {
	//	Funda string     `json:"Funda"`
	PH []KlineData `json:"PH"`
	L  []KlineData `json:"L"`
	R  []KlineData `json:"R"`
	SH []KlineData `json:"SH"`
	ST []KlineData `json:"ST"`
	AH []KlineData `json:"AH"`
	AT []KlineData `json:"AT"`
}
type Display_rate struct {
	Rate float64     `json:"rate"`
	PH []KlineData `json:"PH"`
	L  []KlineData `json:"L"`
	R  []KlineData `json:"R"`
	SH []KlineData `json:"SH"`
	ST []KlineData `json:"ST"`
	AH []KlineData `json:"AH"`
	AT []KlineData `json:"AT"`
}
type ID struct {
	ID int `json:"ID"`
}
