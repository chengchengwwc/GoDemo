package engine

import (
	"go_dev/day9/fetcher"
	"log"
)



type SimpleEngine struct {}





func ( e SimpleEngine )Run(seeds ...Request){
	var requests []Request
	for _,r := range seeds{
		requests = append(requests,r)
	}
	for len(requests) >0{
		r := requests[0]
		requests = requests[1:]
		parseResult,err := worker(r)
		if err != nil{
			continue
		}
		for _,item := range parseResult.Items{
			log.Printf("got item %v",item)

		}
	}

}


func worker(r Request) (ParseResult,error){
	log.Printf("Fetching %s",r.Url)
	body,err := fetcher.Fetch(r.Url)
	if err != nil{
		log.Printf("Fetcher:error %s:%v",r.Url,err)
		return ParseResult{},err
	}
	return r.ParserFunc(body),nil

}











