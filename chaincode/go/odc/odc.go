package main
import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	 sc "github.com/hyperledger/fabric/protos/peer"
	"github.com/zmj/fabric-iot/chaincode/go/m"
)
//订单分配智能合约Order Distribution Contract
//是无人车在获取到订单配送资格后按照订单的下单时间对订单进行合理分配。

//CheckDistribute():首先需要验证无人车是否获得配送资格；
//OrderDistribute():已经确定无人车具有订单配送资格后，物流配送中心按照订单的下单时间，配送距离，等信息对订单进行合理分配。

type OrderDistributeContract interface {
	Init(shim.ChaincodeStubInterface) sc.Response
	Invoke(shim.ChaincodeStubInterface) sc.Response
	Synchro() sc.Response

	Auth(string) (m.ABACRequest, error)
	CheckDistribute(shim.ChaincodeStubInterface)sc.Response
    OrderDistribute(shim.ChaincodeStubInterface)sc.Response
}

type ChainCode struct {
	OrderDistributeContract
}

func NewOrderDistributeContract() OrderDistributeContract {
	return new(ChainCode)
}

func (cc *ChainCode) CheckDistribute(APIstub shim.ChaincodeStubInterface,args []) sc.Response {
	return true;
}

//首先从区块链账本中取出订单数据，然后进行订单分配，订单分配之后更改订单状态为已分配allocate。
func (cc *ChainCode)  OrderDistribute(APIstub shim.ChaincodeStubInterface, args [] string) sc.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of argumentcc. Expecting 1")
	}	
	
	if cc.CheckDistribute(APIstub,args) {
		return shim.Error("The driverless car is qualified for distribution.")
	}
	// 前边是不是应该加个方法将其参数转为字节流形式。 答：觉得没必要
    orderAsBytes, _ := APIstub.GetState(args[0])

	// 将获取到的订单信息转化成OrderInfo结构类型
	order := new(OrderInfo)
	err = json.Unmarshal(orderAsBytes, &order)//应该加地址符
	if err != nil{
		return  shim.Error(err.Error())
	}

	// 将订单的OrderState设置为已分配(assigned)
	order.OrderState("assigned")    
	// order_bytes, err := json.Marshal(order)
	// if err != nil {
	// 	return shim.Error(err.Error())
	// }
	//args[0]是传入的OrderID
    err = APIstub.PutState(args[0], order.ToBytes(order_bytes))
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(m.OK)

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
	if function == "CheckDistribute" {
		return cc.CheckDistribute(APIstub, args)
	} else if function == "OrderDistribute" {
		return cc.OrderDistribute(APIstub, args)
	}
	return shim.Error("Invalid Smart Contract function name.")

// base function
func (cc *ChainCode) Synchro() sc.Response {
	return shim.Success(m.OK)
}

func main() {
	// Create a new Smart Contract
	err := shim.Start(NewOrderDistributeContract())
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
