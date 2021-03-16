package models

type Wishlist struct {
	ID         int `json:"wishlist_id"`
	Productid  int `json:"product_id"`
	Userid     int `json:"user_id"`
	Createdate string
}
