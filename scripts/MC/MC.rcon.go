package mc

import (
	"io/ioutil"
	"fmt"
	rcon "github.com/bearbin/mcgorcon"

	model "github.com/win32prog/sagg.in/web/app/models"
)

func RconWhitelist(w model.Whitelist) {
	pass, err := ioutil.ReadFile("password.txt")
	passstr:= string(pass)

	if err != nil {
		panic(err.Error())
	}
	l, err := rcon.Dial("172.31.18.164", 25575, passstr)
	if err != nil {
		panic(err.Error())
	}
	user:=  w.Mcuser
	fmt.Println(w.Mcuser)
	output, err := l.SendCommand("/whitelist add " + user)
	if err == nil {
		fmt.Println(output)
	}
	output, err = l.SendCommand("/say Player " + user + " has been added!")
	if err == nil {
		fmt.Println(output)
	}
		
	

}
