package main

import (
	"io/ioutil"
	"fmt"
	rcon "github.com/Tnze/go-mc/net"
)

func RconWhitelist(whitelist []Whitelist) {
	pass, err := ioutil.ReadFile("password.txt")
	passstr:= string(pass)

	if err != nil {
		panic(err.Error())
	}
	for _, w := range whitelist {
		l, err := rcon.DialRCON("172.31.18.164:25575", passstr)
		if err != nil {
			panic(err.Error())
		}
		user:=  w.Mcuser
		fmt.Println(w.Mcuser)
		err = l.Cmd("/whitelist add " + user)
		if err != nil {
			panic(err.Error())
		}
	}

}
