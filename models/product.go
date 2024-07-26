package models

type Product struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Quantity  int    `json:"quantity"`
	UnitCoast int64  `json:"unit_coast"`
	MeasureID int    `json:"measure_id"`
}
