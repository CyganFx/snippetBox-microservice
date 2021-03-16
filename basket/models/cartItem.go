package models

type Products struct {
	ID            int `json:"product_id"`
	Title         string
	Discription   string
	Price         float32
	Image         string
	Imagepath     string
	Overallrating int
	Wishlistcount int
	Createddate   string
	Modifieddate  string
	Value         int
	Productcount  int
}
