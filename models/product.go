package models

type Product struct {
	Id        int     `json:"Id"`
	Name      string  `json:"name"`
	Quantity  int     `json:"quantity"`
	UnitCoast float64 `json:"unit_coast"`
	Measure   string  `json:"measure"`
}
