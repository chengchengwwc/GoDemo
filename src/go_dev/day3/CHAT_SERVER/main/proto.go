package main

import "model"




type Message struct{
	Cmd  string `json:"cmd"`
	Data string `json:"data"`
}

type LoginiCmd struct{
	Id     int    `json:"user_id"`
	Passwd string `json:"passwd"`
}


type RegisterCmd struct {
	User model.User `json:"user"`
}

type LoginCmdRes strcut{
	Code int `json:"code"`
	User []int `json:"users"`
	Error string `json:"error"`
}

type UserStatusNotify struct {
	UserId int `json:"user_id"`
	Status int `json:"user_status"` 
}

type UserSendMessageReq struct {
	UserId int   `json:"user_id"`
	Data   string `json:"data"` 
}

type UserRecvMessageReq struct {
	UserId int `json:"user_id"`
	Data   string `json:"data"`
}
