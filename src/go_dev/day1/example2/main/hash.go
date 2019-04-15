package main

import (
	"fmt"
	"go_dev/day1/example2/balance"
	"hash/crc32"
)


type HashBalance struct{
	

}

func init(){
	balance.RegisterBalancer("hash",&HashBalance{})
}




func (p *HashBalance) DoBalance(insts []*balance.Instance,key ...string)(inst *balance.Instance,err error){
	if len(key) == 0{
		err = fmt.Errorf("Hash balacne must pass balance hash key")
		return
	}
	lens := len(insts)
	if (lens == 0){
		err = fmt.Errorf("No backend instance")
		return
	}
	// hash 
	crcTable := crc32.MakeTable(crc32.IEEE)
	hashVal := crc32.Checksum([]byte(key[0]),crcTable)
	index := int(hashVal)%lens
	inst = insts[index]
	return 



}