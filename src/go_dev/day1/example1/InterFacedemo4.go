package example1


import "fmt"



type Student struct{
	Name string
	Sex  string
}


func Test(a interface{}){
	b,ok := a.(int)
	if ok == false{
		fmt.Println("Bad")
		return
	}else{
		b += 3
		fmt.Println(b)
		return
	}
}

func justdemo(items ...interface{}){
	for index,v := range items{
		switch v.(type){
		case bool:
			fmt.Println("bool")
		case int,int16,int32,int64:
			fmt.Println("int")
		case float32:
			fmt.Println("float")
		case string:
			fmt.Println("String")
		default:
			fmt.Println("NO thing")
		
		}
	}
}




func main(){
	va b int
	Test(b)
}