package table

type Profile struct {
	UserID uint `json:"user_id"`
	StockPercent float32 `json:"stock_percent"`
	BondPercent float32 `json:"bond_percent"`
	MMPercent float32 `json:"mm_percent"`
}