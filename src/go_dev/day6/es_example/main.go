package main

import "fmt"
import elastic  "gopkg.in/olivere/elastic.v2"

type Tweet struct{
	User    string
	Message string
}

func main(){
	client,err := elastic.NewClient(elastic.SetSniff(false),elastic.SetURL("http:127.0.0.1:8000"))
	if err != nil{
		fmt.Println("connet es err",err)
		return
	}

	fmt.Println("conn es succ")
	for i :=0;i<1000;i++{
		tweet := Tweet{User:"olive",Message:"Take Five"}
		_,err = client.Index().
		   Index("twitter").
		   Type("tweet").
		   Id(fmt.Sprintf("%d", i)).
		   BodyJson(tweet).
		   Do()
		if err != nil{
			panic(err)
			return
		}

	}
	fmt.Println("insert succ")
}