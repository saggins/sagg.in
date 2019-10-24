package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/minotar/minecraft"
)

type whitelist struct {
	mcuuid string `json:"uuid"`
	mcuser string `json:"name"`
}

func MCwhitelist(username postmsg, wJSON string) {

	minecrafts := minecraft.NewMinecraft()
	uuid, err := minecrafts.GetUUID(username.MCuser)
	if err != nil {
		fmt.Println("Error when getting uuid")
	}
	msg := whitelist{
		mcuuid: uuid,
		mcuser: username.MCuser,
	}

	postNames(msg)
	exsistingList := whitelistScan()

	b, err := json.Marshal(exsistingList)
	if err != nil {
		fmt.Println("Error when marshaling")
	}
	ioutil.WriteFile("whitelist.json", b, os.ModePerm)

}
