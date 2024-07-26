package models

type Product struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Quantity  int    `json:"quantity"`
	UnitCost  int64  `json:"unit_cost"`
	MeasureID int    `json:"measure_id"`
}
