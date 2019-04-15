package example1

import (
	"fmt"
)


type LineNode struct{
	data interface{}
	next *Line
}

type Line struct{
	head *LineNode
	tail *LineNode
}


func (p *Line)InsertNode(data interface{}){
	node := &Line{
		data:data,
		next:nil
	}
	node.next = p.head
	p.head = node
	
}

func (p *Line)InsertTail(data interface{}){
	node := &Line{
		data:data,
		next:nil
	}
	if(p.tail == nil && p.head == nil){
		p.tail = node
		p.head = node
	}
	p.tail.next = node
	p.tail = node
}

func (p *Line)Trans(){
    m := p.head
	for m != nil{
		fmt.Println(m.data)
		m = p.next
	}
}

func main(){

}