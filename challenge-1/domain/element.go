package domain

type Element struct {
	Water  int    `json:"water"`
	Wind   int    `json:"wind"`
	WaterStatus string 
	WindStatus string 
}