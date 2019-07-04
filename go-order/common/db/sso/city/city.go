package city

import (
	"fmt"
)

type CityModel struct {
	ID   int32  `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

func (this *SsoCityModel) GetCityInfo() {
	city := []CityModel{}
	err := this.Db.Select(&city, "select id,name from city")
	fmt.Println("城市信息", err, city, this.Db.Ping())
}
