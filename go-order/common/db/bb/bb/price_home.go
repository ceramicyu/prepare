package bb

import "fmt"

/**
`id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `city_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '城市id',
 */
type PriceHomeModel struct {
	ID        int32          `json:"id" db:"id"`
	CityId    int32          `json:"city_id" db:"city_id"`
}

func (this *BbBbModel)GetPriceHomeInfo(){
	city:=[]PriceHomeModel{}
	err:=this.Db.Select(&city,"select id,city_id from price_home")
	fmt.Println("PriceHome信息",err,city,this.Db.Ping())
}
