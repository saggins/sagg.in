package models


//discord
type Player struct {
	Userid	string `json:"userid"`
	Name    string `json:"name"`
	Ip		string `json:"ip"`
}
type DiscordUser struct {
	Name    string `form:"name"`
}

//saggweb

type Item struct { //DB
	Blobs      []string `json:"blobs"`
	Title      string   `json:"title"`
	ID         string   `json:"id"`
	BlobsTitle []string `json:"blobstitle"`
}

//oauth

type OauthCreds struct {
	Cid string `json:"cid"`
    Csecret string `json:"csecret"`
}

type GoogleProfile struct {
	Sub string `json:"sub"`
    Name string `json:"name"`
    GivenName string `json:"given_name"`
    FamilyName string `json:"family_name"`
    Profile string `json:"profile"`
    Picture string `json:"picture"`
    Email string `json:"email"`
    EmailVerified string `json:"email_verified"`
    Gender string `json:"gender"`
}

type Profile struct {
	Uuid string `json:"sub"`
	FirstName string `json:"given_name"`
	LastName string `json:"family_name"`
	ProfilePic string `json:"picture"`
	Email string `json:"email"`
	Locale string `json:"locale"`
}


//Minecraft

type MCSPost struct {
	Name string `form:"name"`
	Item string `form:"item"`
	Price string `form:"price"`
}
type Postmsg struct { 
	MCuser string `form:"MCuser"`
	Name string `form:"name"`
}

type MCShop struct { //DB
	Name string `json:"name"`
	Item string `json:"item"`
	Price string `json:"price"`
}

type Whitelist struct {
	Mcuuid string `json:"uuid"`
	Mcuser string `json:"name"`
	Name string `json:"rname"`
	Ip string `json:"ip"`
}