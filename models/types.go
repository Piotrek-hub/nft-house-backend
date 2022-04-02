package models

type House struct {
	Id string `form:"id"`
	Title string `form:"title"`
	Type string `form:"type"`
	Location string `form:"location"`
	Price string `form:"price"`
	ImageUrl string `form:"imageUrl"`
	Owner string `form:"owner"`
}