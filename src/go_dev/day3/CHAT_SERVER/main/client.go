package main

import "encoding/binary"
import "encoding/json"
import "errors"
import "fmt"
import "net"

type Client struct {
	conn net.Conn
	userId int
	buf []byte
}

func (p *Client)readPackage()(msg Message,err error ){
	n,err := p.conn.Read(p.buf[0:4])
	if n != 4{
		err = errors.New("read header failed")
		return
	}
	fmt.Println("read package:" , p.buf[0:4])
	var packLen uint32
	packLen = binary.BigEndian.Uint32(p.buf[0:4])
	fmt.Printf("receive len:%d",packLen)
	n,err = p.conn.Read(p.buf[0:packLen])
	if n!= int(packLen){
		err = errors.New("read body failed")
		return
	}
	fmt.Printf("receive data:%s\n",string(p.buf[0:packLen]))
	err = json.Unmarshal(p.buf[0:packLen],&msg)
	if err != nil{
		fmt.Println("Unmarsha1 failed,err",err)
		return 
	}
	return
}

func (p *Client)writePackage(data []byte)(err error){
	packLen := uint32(len(data))
	binary.BigEndian.PutUint32(p.buf[0:4],packLen)
	n,err := p.conn(p.buf[0:4])
	if err != nil{
		fmt.Println("write data failed")
		return 
	}
	n,err = p.conn.Write(data)
	if err != nil{
		fmt.Println("write data failed")
	}
	
	if n != int(packLen){
		fmt.Println("write data not finished")
		err =  errors.New("write data not finished")
		return
	}
	return
}


func (p *Client)Process()(err error){
	for{
		var msg Message
		msg,err = p.readPackage()
		if err != nil{
			clientMgr.DelClient(p.userId)
			return err
		}


	}

}


func (p *Client) processMsg(msg Message)(err error){
	switch msg.Cmd{
	case UserLogin:
		err = p.
	}
}



func (p *Clinet)LoginResp(err error){
	var respMsg Message
	respMsg.Cmd = UserLoginRes
	var loginRes LoginCmdRes
	loginRes.Code = 200
	userMap := clientMgr.GetAllUsers()
	for userId,_ range userMap{
		loginRes.User = append(loginRes.User,userId)
	}
	data,err := json.Marshal(loginRes)
	if err != nil{
		fmt.Println("marsha failed: ",err)
		return
	}
	respMsg.Data = string(data)
	data,err = json.Marshal(respMsg)
	if err != nil{
		fmt.Println("marshal failed: ",err)
		return 
	}
	err = p.writePackage(data)
	if err != nil{
		fmt.Println("send failed, ",err)
		return
	}
	return
}

func (p *Client) login(msg proto.Message)(err error){
	defer func(){
		p.loginresp(err)
	}()
	fmt.Println("recv user login request,data:%v",msg)
	var cmd.proto.LoginCmd
	err = json.Unmarshal([]byte(msg.Data),&cmd)
	if err != nil{
		fmt.Println("Unmarsha failed,err:",err)
		return
	}
	_,err = mgr.Login(cmd.Id,cmd.Passwd)
	if err != nil{
		return
	}
	ClientMgr.AddClient(cmd.Id,p)
	p.userId = cmd.Id
	p.NotifyOthersUserOnline(cmd.Id)

}




func (p *Client) NotifyOthersUserOnline(userId int){
	users := ClientMgr.GetAllUsers()
	for id,client := range users{
		if id == userId{
			continue
		}
		client.NotifyUserOnline(userId)
	}

}







func (p *Clinet)NotifyUserOnline(userId int){
	var resMsg Message
	resMsg.Cmd = UserStatusNotifyRes
	var noitfyRes UserStatusNotify
	noitfyRes.UserId = userId
	noitfyRes.Status = UserOnline
	data,err := json.Marshal(noitfyRes)
	if err != nil{
		fmt.Println("marshal failed", err)
		return
	}
	resMsg.Data = string(data)
	data,err = json.Marshal(respMsg)
	if err != nil{
		fmt.Println("marshal failed", err)
		return
	}
	err = p.writePackage(data)
	if err != nil{
		fmt.Println("send failed,",err)
		return
	}

}

func (p *Clinet)register(msg proto.Message)(err error){
	var cmd RegisterCmd
	err := json.Unmarshal([]byte(msg.Data),&cmd)
	if err != nil{
		return
	}
	err = mgr.Register(&cmd.User)
	if err != nil{
		return
	}
	return

}