package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
	"snippetBox-microservice/basket/api/grpc/protobuffs"
	mysqldb "snippetBox-microservice/basket/connection"
	"snippetBox-microservice/basket/models"
	"strconv"
	"time"
)

func Addwishlist(c *gin.Context) {

	session, _ := store.Get(c.Request, "mysession")

	userid := session.Values["userid"]

	if userid == nil {
		c.JSON(301, "/")
	} else {
		db := mysqldb.SetupDB()
		fmt.Println("sess uid", userid)

		var product models.Products
		err := c.BindJSON(&product)
		if err != nil {
			panic(err.Error())
		}
		pid := product.ID

		fmt.Println("pid:", pid)

		created_date := time.Now().Format("Jan 2, 2006 03:04:05 PM")

		var count int

		er := db.QueryRow("SELECT count(product_id) cnt FROM tbl_wishlist WHERE user_id = ? AND product_id = ?", userid, pid).Scan(&count)
		if er != nil {
			fmt.Println(er)
		}
		fmt.Println("count", count)
		if count != 1 {

			insWish, err := db.Prepare("INSERT INTO  tbl_wishlist(product_id,user_id,created_date) VALUES(?,?,?)")

			if err != nil {
				fmt.Println(err)
			}
			insWish.Exec(pid, userid, created_date)

			json.NewEncoder(c.Writer).Encode("Added in your wishlist")

		} else {

			json.NewEncoder(c.Writer).Encode("Already in your wishlist")
		}
		defer db.Close()
	}

}

func WishListFromCatalog(c *gin.Context) {

	//grpc client connection
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	grpc_client := protobuffs.NewCatalogServiceClient(conn)

	session, _ := store.Get(c.Request, "mysession")

	userid := session.Values["userid"]
	fmt.Println("userid", userid)
	if userid == nil {
		c.JSON(301, "/viewwishlist")
	} else {
		db := mysqldb.SetupDB()

		WishRows, err := db.Query("SELECT product_id FROM tbl_wishlist where user_id=?", userid)
		fmt.Println("Wish rows", *WishRows)
		if err != nil {
			fmt.Println(err)
		}
		defer WishRows.Close()
		//defer db.Close()

		var res1 []models.ProductModel

		for WishRows.Next() {
			//fmt.Println("Wish rows", WishRows)
			//fmt.Println("Wish rows next: ", WishRows.Next())
			wishlist := models.ProductModel{}
			//var product_id int
			_ = WishRows.Scan(&wishlist.ID)

			fmt.Println("Product id: ", wishlist.ID)

			request := &protobuffs.ProductSendRequest{Id: int32(wishlist.ID)}
			fmt.Println("Request id: ", request.Id)
			response, err := grpc_client.SendProduct(context.Background(), request)
			if err != nil {
				log.Fatalf("error while calling SendProduct %v", err)
			}
			fmt.Println("Response title: ", response.Title)
			wishlist.ID = int(response.Id)
			wishlist.Title = response.Title
			wishlist.Category = response.Category
			wishlist.Description = response.Description
			wishlist.Price = response.Price

			res1 = append(res1, wishlist)
		}

		name := session.Values["firstname"]

		c.JSON(200, gin.H{"wishlist": res1, "name": name})

	}
}

func Wishlist(c *gin.Context) {

	session, _ := store.Get(c.Request, "mysession")

	userid := session.Values["userid"]
	fmt.Println("userid", userid)
	if userid == nil {
		c.JSON(301, "/viewwishlist")
	} else {
		db := mysqldb.SetupDB()

		WishRows, err := db.Query("SELECT product_id,id FROM tbl_wishlist where user_id=?", userid)

		if err != nil {
			fmt.Println(err)
		}
		wishlist := models.Products{}
		res1 := []models.Products{}
		for WishRows.Next() {
			var wid, product_id int
			_ = WishRows.Scan(&product_id, &wid)

			var title, image_path string
			var price float32
			var quantity int
			_ = db.QueryRow("SELECT title,price,image_path FROM tbl_products WHERE id=?", product_id).Scan(&title, &price, &image_path)

			db.QueryRow("SELECT quantity FROM tbl_orders WHERE product_id=? AND user_id=?", product_id, userid).Scan(&quantity)
			wishlist.Title = title

			wishlist.Price = price
			wishlist.Imagepath = image_path
			wishlist.ID = wid
			wishlist.Overallrating = product_id
			wishlist.Value = quantity
			res1 = append(res1, wishlist)
		}
		var ordercount int
		_ = db.QueryRow("SELECT count(id)  FROM tbl_orders WHERE user_id = ? ", userid).Scan(&ordercount)

		fmt.Println("ordercount", ordercount)

		name := session.Values["firstname"]

		c.JSON(200, gin.H{"wishlist": res1, "orderscount": ordercount, "name": name})

		defer db.Close()
	}

}

func Deletewishlist(c *gin.Context) {
	var wishList models.Wishlist

	err := c.BindJSON(&wishList)
	if err != nil {
		panic(err.Error())
	}
	wid := wishList.ID

	db := mysqldb.SetupDB()
	wid, err = strconv.Atoi(c.Query("wid"))
	if err != nil {
		log.Fatal("error in delete wish list")
	}
	fmt.Println("wid", wid)
	delWish, err := db.Prepare("DELETE FROM  tbl_wishlist WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	delWish.Exec(wid)

	defer db.Close()
	c.JSON(301, "/wishlist")
}
func Deletewishlistall(c *gin.Context) {
	session, _ := store.Get(c.Request, "mysession")

	userid := session.Values["userid"]

	fmt.Println("sess user:", userid)

	if userid == nil {
		c.JSON(301, "/")
	}
	db := mysqldb.SetupDB()
	delWish, err := db.Prepare("DELETE FROM  tbl_wishlist WHERE user_id=?")
	if err != nil {
		panic(err.Error())
	}
	delWish.Exec(userid)

	defer db.Close()
	c.JSON(301, "/wishlist")
}
