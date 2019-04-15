package balance

import (
	"fmt"
)

type BalanceMer struct{
	allBalancer map[string]Balance
}

var (
	//初始化 
	mgr = BalanceMer{
		allBalancer:make(map[string]Balance),
	}
)


func(p *BalanceMer) regiserBalance(name string,b Balance){
	p.allBalancer[name] =b

}

func RegisterBalancer(name string,b Balance){
	mgr.regiserBalance(name,b)
}


func DoBalance(name string,insts []*Instance) (inst *Instance,err error){
	balancer,ok := mgr.allBalancer[name]
	if !ok {
		err =fmt.Errorf("Not Fount balance",name)
		return
	}
	inst,err = balancer.DoBalance(insts)
	return 
}





