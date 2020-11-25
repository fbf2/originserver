package simple_rpc

import (
	"fmt"
	"github.com/duanhf2012/origin/node"
	"github.com/duanhf2012/origin/rpc"
	"github.com/duanhf2012/origin/service"
	"time"
)

func init(){
	node.Setup(&TestService6{})
}

type TestService6 struct {
	service.Service
}

func (slf *TestService6) OnInit() error {
	return nil
}

type InputData struct {
	A int
	B int
}

func (slf *TestService6) RPC_Sum(input *InputData,output *int) error{
	*output = input.A+input.B
	return nil
}

func (slf *TestService6) RPC_RawTest(input interface{}) error{
	b := input.([]byte)
	fmt.Println(string(b))

	return nil
}

func (slf *TestService6) RPC_SyncTest(resp rpc.Responder,input *int,out *int) error{
	go func(){
		time.Sleep(3*time.Second)
		var output int = *input
		resp(&output,rpc.NilError)
	}()

	return nil
}
