package m

import (
	"fmt"
	"testing"
)
//这是朱宇杰刚才写的，他说按照师兄来写就行。他写的应该有错因为我看写的很急
var p = ORDERequest{
	"fahuodizhi","shouhuodizhi","dianhua","xinyu"
}

func Test_ToBytes(t *testing.T) {
	fmt.Printf("%s\n", p.ToBytes())
}

func Test_GetID(t *testing.T) {
	println(p.GetID())
}