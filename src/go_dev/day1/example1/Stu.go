package example1

type Student struct{
	Name string
	Crade string
	Id  string
	Sex string
	books []*BorrowItem
}

type BorrowItem struct{
	book *Book
	num  int

}


func CreateStudent(name,grade,id,sex string) *Student{
	stu := &Student{
		Name:name,
		Crade:grade,
		Id:id,
		Sex:sex
	}
	return stu
}

func (s *Student)AddBook(b *BorrowItem){
	s.books = append(s.books,b)

}

func (s *Student)DelBook(b *BorrowItem) (err error){
	found := false
	
	for i :=0;i<len(s.books);i++{
		if (s.books[i].book.Name == b.book.Name){
			if(b.num == s.books[i].num){
				front:= s.books[0:i]
				left := s.books[i+1:]
				front = append(front,left)
				s.books = front
				return 
			}else{
				s.books[i].num -= b.num
				return
			}

		}
	}
	err = ErrStockNotEnough
	return
}


func (s *Student)GetBookList()([]*BorrowItem){
	return s.books

}







