package main

import (
	rcon "github.com/Tnze/go-mc/net"
)

func RconWhitelist(whitelist []Whitelist) {
	l, err := rcon.DialRCON("0.0.0.0:25575", "1234")

	if err != nil {
		panic(err.Error())
	}
	for _, w := range whitelist {
		err = l.Cmd("/whitelist add " + w.Mcuser)
	}

}
