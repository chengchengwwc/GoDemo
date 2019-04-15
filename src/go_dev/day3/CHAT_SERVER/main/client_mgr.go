package main

import "fmt"


type ClientMgr struct{
	onlineUsers may[int]*Client
}

var (
	clientMgr *ClientMgr
)

func init(){
	clientMgr = &ClientMgr{
		onlineUsers:make(map[int]*Client,1024)
	}
}

func (p *ClientMgr)AddClient(userId int,client *Client){
	p.onlineUsers[userId] = client
}

func (p *ClientMgr)GetClient(userId int)(err error){
	client,ok := p.onlineUsers[userId]
	if !ok {
		err = fmt.Errorf("user %d not exists",userId)
		return
	}
	real
}

func(p *ClientMgr)GetAllUsers() map[int]*Client{
	return p.onlineUsers
}

func (p *ClientMgr)DelClient(userId int){
	delete(p.onlineUsers,userId)
}









