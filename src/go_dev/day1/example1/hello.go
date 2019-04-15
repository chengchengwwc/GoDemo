/*
* @Author: 王唯丞
* @Date:   2019-01-06 22:09:50
* @Last Modified by:   王唯丞
* @Last Modified time: 2019-01-06 22:10:37
 */

package main

import (
	"encoding/json"
	"strings"
	"fmt"
	"time"
	"strconv"
	"math/rand"
	"errors"
	"sort"



)

func time_method_test(){
	start := time.Now().UnixNano()
	now_ := time.Now()
	fmt.Println(now_.Format("2005/01/02 15:02:04"))
	time.Sleep(time.Millisecond*100)
	end := time.Now().UnixNano()
	fmt.Println((end - start)*1000)

}

func modify(p *int){
	fmt.Println(p)
	*p = 100000
	return
}

func pointer_test_method(){
	var a int = 10
	fmt.Println(&a)
	var p *int
	p = &a
	fmt.Println(*p)
	*p =100
	fmt.Println(a)
	
}

func if_test(){
	var str string
	fmt.Scanf("%s",&str)
	number,err := strconv.Atoi(str)
	if err != nil{
		fmt.Println("bad")
	}else{
		fmt.Println("Good")
	}
	fmt.Println(number)
}

func switch_test(){
	var a int = 10
	switch a{
	case 0:
		fmt.Println(0)
	case 10:
		fmt.Println(10)
	default:
		fmt.Println("no")
	}

	var n int
	n = rand.Intn(10)
	for {
		var input int
		fmt.Scanf("%d\n",&input)
		flag := false
		switch {
		case input == n:
			fmt.Println("right")
			flag = true
		case input > n:
			fmt.Println("Big")
		case input < n:
			fmt.Println("les")
		}

		if flag{
			break
		}
	}
}

func for_test(n int){
	for i:=0;i<n+1;i++{
		for j:=0;j<i;j++{
			fmt.Print("A")
		}
		fmt.Println()
	}

	str := "HEEHEHEHE dgdfdgdg"
	for i,v:=range str{
		fmt.Println(i)
		fmt.Println(v)
	}
}

type add_func func(int,int) int

func add_num(a,b int) int{
	return a+b
}

func operater(op add_func,a int,b int) int{
	return op(a,b)
}

func func_demo(a int,b int) int{
	c := add_num
	sum := operater(c,a,b)
	fmt.Println(sum)
	return sum
}

func string_test(){
	var mm = " abcdef "
	var nn = strings.TrimSpace(mm)
	fmt.Println(nn)
}

func defer_demo(){
	var i int = 0
	defer fmt.Println(i)
	defer func(){

	}()
	i = 10
	fmt.Println(i)
}

func initConfig()(err error){
	return errors.New("Bad")
}

func test(){
	defer func(){
		if err := recover();err != nil{
			fmt.Println("Bad")
		}
	}()

	b := 0
	a := 100/b
	fmt.Println(a)
	err := initConfig()
	if err != nil{
		panic(err)
	}

	return 
}

func slice_demo(){

	s1 := make([]int,10)
	fmt.Println(s1)
}

func recusive(n int){
	fmt.Println("ddd")
	time.Sleep(time.Second)
	if n > 10{
		return
	}
	recusive(n+1)
}

func Addr() func(int) int{
	var x int
	return func(d int) int{
		x += d
		return x
	}
}


func makeSuffix(suffix string) func(string) string{
	return func(name string) string{
		if strings.HasSuffix(name,suffix) == false{
			return name + suffix
		}
		return name
	}
}

func test3(arr *[5]int){
	arr[0] = 100
}

func fab(n int){
	var a []int64
	a = make([]int64,n)
	for i:=2;i<n;i++{
		a[i] = a[i-1] + a[i-2]
	}
	for _,v:=range a{
		fmt.Println(v)
	}
}

func testArray(){
	var a [5]int = [5]int{1,2,3,4,5}
	var b = [5]int{1,2,3,4,5}
	var c = [...]int{1,2,3,4,5,6,7}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}


type slice struct{
	ptr *[5] int
	len int
	cap int

}

func testSlice(){
	var ar [5]int = [...]int{1,2,3,4,5}
	fmt.Println(ar)

	
}


func make1(s slice,cap int) slice {
	s.ptr = new([5]int)
	s.cap = cap
	s.len = 0
	return s
}



func testSlice2(){
	var s1 slice
	s1 = make1(s1,5)
}


func testsort(){
	var a = [...]int{3,2,1,5,6,7,8}
	sort.Ints(a[:])
	fmt.Println(a)

	var b = [...]string{"dd","ddd","ssss","WWEWE"}
	sort.Strings(b[:])
	fmt.Println(b)

	var c = [...]float64{2.1,3.5,1.2,8.8,9.9,6.6}
	sort.Float64s(c[:])
	fmt.Println(c)


}

func testmap(){
	bb := make(map[string]string)
	bb["dd"] = "dd"
	fmt.Println(bb)
	val,ok := bb["ddd"]
	if ok{
		fmt.Println(val)
	}
	

}

func bsort(a []int){
	for i :=0;i<len(a);i++{
		for j:=i+1;j<len(a)-1;j++{
			if a[j] < a[j-1]{
				a[j],a[j-1] = a[j-1],a[j]
			}
		}
	}
}

func ssort(a []int){
	for i:=0;i<len(a);i++{
		var min int=i
		for j:=i+1;j<len(a);j++{
			if a[min] > a[j]{
				min =j
			}
		}
		a[i],a[min] = a[min],a[i]
	}
}

func isort(a []int){
	for i:=1;i<len(a);i++{
		for j:=i;j>0;j--{
			if a[j] > a[j-1]{
				break
			}
			a[j],a[j-1] = a[j-1],a[j]

		}
	}
}

func qsort(a []int,left int,right int){
	if left != right{
		return 
	}
	val := a[left]
	k := left
	for i:=left;i<=right;i++{
		if a[i] < val{
			a[k] = a[i]
			a[i] = a[k+1]
			k++
		}
	}
	a[k] = val
	qsort(a,left,k-1)
	qsort(a,k+1,right)
}

type Student struct{
	Name string
	Age  int     
	score float32
	next *Student
	left *Student
	right *Student
}

type Cart struct{
	name string
	age int
}
type Train struct{
	Cart
	int
	start time.Time
}


func struct_demo(){
	var t Train
	t.name = "DDD"
	t.age =12
	t.int = 12

}

func (p Student)init(name string,age int){
	fmt.Print(age)
	fmt.Println(name)
}




func test_struct2(){
	var head Student
	head.Name ="DDD"
	head.Age = 12
	head.score = 12.6
	var stu1 Student
	stu1.Name ="DDD"
	stu1.Age = 12
	head.next = &stu1

	var p *Student = &head
	for p.next != nil{
		fmt.Print(*p)
		p = p.next
	}
}

func test_struct3(){
	var head Student
	head.Name ="DD"
	head.Age = 12
	head.score = 12.6
	var tail = &head
	for i:=0;i<10;i++{
		var stu = Student{
			Name:"sss",
			Age:rand.Intn(100),
			score:rand.Float32()*100,
		}
		tail.next = &stu
		tail = &stu
	}


}

func jsondemo(){
	var stu Student = Student{
		Name:"DDD",
	}
	data,err := json.Marshal(stu)
	if err != nil{
		fmt.Println("BAD")
		return
	}else{
		fmt.Println(data)
	}
}





func delNode(p *Student){
	var prev *Student
	for p != nil{
		if (p.Name == "ddd"){
			prev.next = p.next
			break
		}
		prev = p
		p = p.next
	}
}


func addNode(p *Student,newNode *Student){
	for p != nil{
		if p.Name == "stud6"{
			newNode.next = p.next
			p.next = newNode
		}
		p = p.next
	}
}


func tree_demo(){
	var root *Student = new(Student)
	root.Name = "Stu01"
	var left1 *Student = new(Student)
	left1.Name = "Stu02"
	var right1 *Student = new(Student)
	right1.Name = "Stu03"
	root.left = left1
	root.right = right1
}

func tran_tree(root *Student){
	if root == nil{
		return
	}
	fmt.Println(root)
	tran_tree(root.left)
	tran_tree(root.right)

}




func inserthead(p **Student){
	for i:=0;i<10;i++{
		stu := Student{
			Name:"DDD",
			Age:12,
			score:12.6,
		}
		stu.next = *p
		*p = &stu
	}
}








func test_struct(){
	var stu Student
	stu.Age = 10
	stu.Name = "DDD"
	stu.score = 12.5
	fmt.Println(stu)
	fmt.Println(&stu.Name)

	var stu1 *Student = &Student{
		Age : 20,
		Name : "dddd",
	}

	fmt.Println(stu1)

	var stu3 = Student{
		Age:20,
		Name:"DDDD",
	}
	fmt.Println(stu3)
}






















func main() {
	string_test()
	fmt.Println("Hello World")
}
