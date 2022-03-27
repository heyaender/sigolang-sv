package table

type customer struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age int `json:"age"`
	Profile *Profile `gorm:"foreignkey:UserID" json:"profile,omitempty"` 
}