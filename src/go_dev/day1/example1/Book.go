package example1


import (
	"errors"
)

var (
	ErrStockNotEnough = errors.New("Book is not enough")
)



type Book struct{
	Name string
	Total int
	Author string
	CreateTime string
}

func CreateBook(name string,total int,author,createtime string)(b *Book){
	b := &Book{
		Name:name,
		Total:total,
		Author:author,
		CreateTime:createtime
	}
	return b
	
}

func (b *Book) canBorrow(c int) bool{
	return b.Total >= c
}

func (b *Book) Borrow(c int)(err error){
	if b.canBorrow(c) == false{
		err = ErrStockNotEnough
		return err
	}
	b.Total -= c
}

func (b *Book) Back(c int)(err error){
	b.Total += c
	return
}







