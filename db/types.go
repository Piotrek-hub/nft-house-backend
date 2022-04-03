package db

type House struct {
	Id       string `form:"id" json:"id"`
	Title    string `form:"title" json:"title"`
	Type     string `form:"type" json:"type"`
	Location string `form:"location" json:"location"`
	Price    string `form:"price" json:"price"`
	ImageUrl string `form:"imageUrl" json:"imageUrl"`
	Owner    string `form:"owner" json:"owner"`
}
