package MetroAPI

type items struct {
	Items []item `json:"items"`
}

type item struct {
	Seconds     float32 `json:"seconds"`
	BlockID     string  `json:"block_id"`
	RouteID     string  `json:"route_id"`
	IsDeparting bool    `json:"is_departing"`
	RunID       string  `json:"run_id"`
	Minutes     float32 `json:"minutes"`
}

type byMinutes []item

func (a byMinutes) Len() int           { return len(a) }
func (a byMinutes) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byMinutes) Less(i, j int) bool { return a[i].Minutes < a[j].Minutes }

type input struct {
	StopID      string
	StationName string
}
