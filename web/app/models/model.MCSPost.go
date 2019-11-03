package models

type MCSPost struct {
	Name string `form:"name"`
	Item string `form:"item"`
	Price string `form:"price"`
}