package main
import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	 sc "github.com/hyperledger/fabric/protos/peer"
	"github.com/zmj/fabric-iot/chaincode/go/m"
)

//订单管理 智能合约order management contract
//是对用户订单进行管理，比如用户增加订单，撤销订单，改动订单收货信息
//以及订单丢失如何赔付情况，订单配送超时情况，订单配送收费情况,当无人车遇到故障时订单交接情况的管理。

type OrderManageContract interface {
	Init(shim.ChaincodeStubInterface) sc.Response
	Invoke(shim.ChaincodeStubInterface) sc.Response
	Synchro() sc.Response

	parseOrder(shim.ChaincodeStubInterface,args []string)sc.Response
	GetOrder(shim.ChaincodeStubInterface,args []string)sc.Response
    AddOrder(shim.ChaincodeStubInterface,args []string)sc.Response
	UpdateOrder(shim.ChaincodeStubInterface,args []string)sc.Response
    DeleteOrder(shim.ChaincodeStubInterface,args []string)sc.Response
	OrderCharge(shim.ChaincodeStubInterface)sc.Response
}

type ChainCode struct {
	OrderManageContract
}

/*
 * The Init method is called when the Smart Contract "fabcar" is instantiated by the blockchain network
 * Best practice is to have any Ledger initialization in separate function -- see initLedger()
*/
func (cc *ChainCode) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(m.OK)
}

/*
 * The Invoke method is called as a result of an application request to run the Smart Contract "fabcar"
 * The calling application program has also specified the particular smart contract function to be called, with arguments
 */

 func (cc *ChainCode) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {
	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "GetOrder" {
		return cc.GetOrder(APIstub, args)
	} else if function == "AddOrder" {
		return cc.AddOrder(APIstub, args)
	}else if function == "DeleteOrder" {
		return cc.DeleteOrder(APIstub, args)
	}else if function == "UpdateOrder" {
		return cc.UpdateOrder(APIstub, args)
	}else if function == "OrderCharge" {
		return cc.OrderCharge(APIstub)
	}
	return shim.Error("Invalid Smart Contract function name.")
}

func NewOrderManageContract() OrderManageContract {
	return new(ChainCode)
}

// func (cc *ChainCode) parsePolicy(arg string) (m.Policy, error) {
// 	policyAsBytes := []byte(arg)
// 	policy := m.Policy{}
// 	err := json.Unmarshal(policyAsBytes, &policy)
// 	return policy, err
// }

func (cc *ChainCode) parseOrder(arg string) (m.Policy, error) {
	orderAsBytes := []byte(arg)
	order := m.OrderInfo{}
	err := json.Unmarshal(orderAsBytes, &order)//将json字符串解码到相应数据结构。
	return order, err
}

func (cc *ChainCode) GetOrder(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	b, _ := APIstub.GetState(args[0])
	return shim.Success(b)
}

func (cc *ChainCode)  AddOrder(APIstub shim.ChaincodeStubInterface，args []string) sc.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of ordercc. Expecting 1")
	}
	// parse order
	order, err := cc.parseOrder(args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	// check Order 验证用户发出订单请求的合理性
	if cc.CheckOrder(order) {
		return shim.Error("bad order Request")
	}
	// put k-v to DB
	err = APIstub.PutState(order.GetID(), order.ToBytes())
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(m.OK)
}

func (cc *ChainCode) UpdateOrder(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of argumentcc. Expecting 1")
	}
	// parse request 
	order, err := cc.parseOrder(args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
   // check Order 验证用户发出订单请求的合理性
	if cc.CheckOrder(order) {
		return shim.Error("bad order Request")
	}
	r := cc.QueryRequest(APIstub, []string{order.GetID()})
	if r.GetStatus() != 200 {
		return shim.Error("order not exist")
	}
	return cc.AddOrder(APIstub, args)
}

func (cc *ChainCode)  DeleteOrder(APIstub shim.ChaincodeStubInterface,args []string) sc.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of ordercc. Expecting 1")
	}
	// parse request 
	order, err := cc.parseOrder(args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
   // check Order 验证用户发出订单请求的合理性
	if cc.CheckOrder(order) {
		return shim.Error("bad order Request")
	}
	err := APIstub.DelState(args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(m.OK)
}
//当货物丢失时应由无人车方进行一定金额赔偿，当订单配送超时应按照 事先约定的赔偿准则 进行赔付以及订单送达时触发 智能合约 调用该方法完成付款操作。
// 这个不知道该怎么写，因为准则没制定，考虑用空方法。
func (cc *ChainCode)  OrderCharge(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(m.OK)
}

// base function
func (cc *ChainCode) Synchro() sc.Response {
	return shim.Success(m.OK)
}

func main() {
	// Create a new Smart Contract
	err := shim.Start(NewOrderManageContract())
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
