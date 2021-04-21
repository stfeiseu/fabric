package m

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
)

// Define the Policy structure
type Policy struct {
	AS AS
	AO AO
	AP int //1 is allow , 0 is deney
	AE AE
}

type AS struct {
	UserId string `json:"userId"`
	Role   string `json:"role"`
	PK  string `json:"pk"`
	Cert string `json:"cert"` 
}

// type AO struct {
// 	DeviceId string `json:"DeviceId"`
// 	MAC      string `json:"MAC"`
// }

// PackageName string `json:"packageName"` //收货方姓名,感觉这个没必要
// PackageAddress string `json:"packageAddress"` //收货方地址
// PackagePhone string `json:"packagePhone"` //收货方电话

type AO struct {
    OrderId int  `json:"orderId"`  //订单ID
}

type AE struct {
	CreateTime int64  `json:"createTime"`
	EndTime int64  `json:"endTime"`
	Address string  `json:"address"`
	CurrentAccess int64  `json:"currentAccess"`  // 是否可访问
}

func (p *Policy) ToBytes() []byte {
	b, err := json.Marshal(*p)//将数据转为json类型字符串。
	if err != nil {
		return nil
	}
	return b
}

func (p *Policy) GetID() string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(p.AS.UserId+p.AO.OrderId)))
}
