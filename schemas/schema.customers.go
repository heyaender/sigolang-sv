package schemas

type Customer struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age int `json:"age"`
	Profile *Profile `gorm:"foreignkey:CustomerID" json:"profile,omitempty"` 
}