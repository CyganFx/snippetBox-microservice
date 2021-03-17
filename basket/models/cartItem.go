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

type ProductModel struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Category    string  `json:"category"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
}
