package util

import (
	"io"
	"crypto/rand"
	"fmt"
)




func NewUUID()(string,error){
	uuid := make([]byte,16)
	n,err := io.ReadFull(rand.Reader,uuid)
	if n != len(uuid) || err != nil{
		return "",err
	}
	return fmt.Sprintf("%x-%x-%x-%x-%x",uuid[0:4],uuid[4:6],uuid[6:8],uuid[8:10],uuid[10:12]),nil

}