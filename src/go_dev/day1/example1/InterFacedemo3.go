package example1




type Reader interface{
	Read()
}

type Writer interface{
	write()
}



type ReadWrite interface{
	Reader
	Writer
}


type File struct{

}


func (p *File)Read(){
	
}

func (p *File)write(){

}

func Test(rw ReadWrite){
	rw.Read()
	rw.write()
}

func main(){
	var f File
	Test(&f)
}

















