package auth
import (
	"golang.org/x/oauth2"
	"crypto/rand"
	"encoding/base64"
	
)

func Randomtoken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}

func GetLoginURL(conf oauth2.Config, state string) string {

	return conf.AuthCodeURL(state) 
}

// Is Registered?
func CheckSession(sessionid) bool {
	
}

