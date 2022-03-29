package databases

import (
	"fmt"
	"sigolang-sv/schemas"
)

func DBMigrate() {
	DB.AutoMigrate(&schemas.Customer{}, &schemas.Profile{})
	fmt.Println("DB Migrate Success")
}
