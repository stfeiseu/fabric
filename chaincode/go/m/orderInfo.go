package m

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
)

// OrderState string `json:"orderState"` // 订单状态，假设以assigned表示订单已被分配无人车进行运送
//订单数据包含字段：订单ID，订单下单时间,订单状态，订单私密程度，订单加急程度.
// Define the OrderInfo structure
type OrderInfo struct {
    OrderId int  `json:"orderId"`  //订单ID
	OrderTime string  `json:"orderTime"`  //订单下单时间
	PackageName string `json:"packageName"` //订单货物名字
	ReceiverName string `json:"receiverName"` //收货方姓名
	ReceiverAddress string `json:"receiverAddress"` //收货方地址
	ReceiverPhone string `json:"receiverPhone"` //收货方电话
	ReceiverPeput string `json:"receiverPeput"` //收货方信誉
	OrderState string `json:"orderState"`     // 订单状态，假设以assigned表示订单已被分配无人车进行运送
	Private int `json:"private"`   // 是否私密，数字越大，表示私密程度越高
	Urgent int `json:"urgent"`    // 加急程度，数字越大，表示加急程度越高
}

func (o *OrderInfo) ToBytes() []byte {
	b, err := json.Marshal(*o)//将数据转为json类型字符串。
	if err != nil {
		return nil
	}
	return b
}
// func (p *Policy) GetID() string {
// 	return fmt.Sprintf("%x", sha256.Sum256([]byte(p.AS.UserId+p.AO.DeviceId)))
// }

func (o *OrderInfo) GetID() string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(o.OrderId)))
}

func NewOrderInfo(b []byte) (OrderInfo, error) {
	r := OrderInfo{}
	err := json.Unmarshal(b, &r)
	return r, err
}



