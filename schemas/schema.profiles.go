package schemas

type Profile struct {
	CustomerID uint `json:"customer_id"`
	StockPercent float32 `json:"stock_percent"`
	BondPercent float32 `json:"bond_percent"`
	MMPercent float32 `json:"mm_percent"`
}