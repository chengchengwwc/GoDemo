package taskrunner

import (
	"sync"
	"os"
	"go_dev/day11/video_server/api/dbops"
	"log"
	"errors"
)

func deleteVideo(vid string) error{
	err := os.Remove("./video/" + vid)
	if err != nil && !os.IsNotExist(err){
		log.Printf("BAD")
		return err
	}
	return nil
}





func VideoClearDispatcher(dc dataChan) error {
	res,err:= dbops.ReadVideoDelteionRecord(3)
	if err != nil{
		log.Print("Vedio bad")
		return err
	}
	if len(res) == 0{
		return errors.New("All task finished")
	}
	for _,id := range res{
		dc <- id
	}
	return nil
}

func VideoClearExecutor(dc dataChan) error{
	errMap := &sync.Map{}
	var err error
	forloop:
		for {
			select {
			case vid := <- dc:
				go func(id interface{}){
					if err := deleteVideo(id.(string)); err != nil{
						errMap.Store(id,err)
						return
					}

					err := dbops.DelVideoDeltetionRecord(id.(string))
					if err != nil{
						errMap.Store(id,err)
						return
					}
				}(vid)
				default:
					break forloop
			}
		}

	errMap.Range(func(k,v interface{}) bool {
		err = v.(error)
		if err != nil{
			return false
		}
		return true
	})
	return err
}