package main

import (
	"io/ioutil"

	rcon "github.com/Tnze/go-mc/net"
)

func RconWhitelist(whitelist []Whitelist) {
	pass, err := ioutil.ReadFile("password.txt")
	l, err := rcon.DialRCON("172.31.18.164:25575", string(pass))

	if err != nil {
		panic(err.Error())
	}
	for _, w := range whitelist {
		err = l.Cmd("/whitelist add " + w.Mcuser)
	}

}
