package scheduler

import "go_dev/day9/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}


func (s *SimpleScheduler)ConfigureMasterWorkerChan(c chan engine.Request){
	s.workerChan = c

}







func (s *SimpleScheduler)Subimt(r engine.Request){
	// send request
	//解决死循环
	go func() {
		s.workerChan <- r
	}()
}