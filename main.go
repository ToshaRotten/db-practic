package main

import (
	"database/sql"
	_ "database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type FullInfo struct {
	box   Box
	prod  Product
	order Order
}

type Box struct {
	ID   int
	Name string
}

type Product struct {
	ID    int
	Name  string
	Price int
}

type Order struct {
	ID         int
	ClientName string
	Phone      string
}

func (b *Box) GetBoxByID(id int, db *sql.DB) *Box {
	sqlString := fmt.Sprintf("SELECT b.ID_box, b.name FROM `box` b WHERE b.ID_box = %d", id)
	row := db.QueryRow(sqlString)
	err := row.Scan(&b.ID, &b.Name)
	if err != nil {
		log.Fatal(err)
	}
	return b
}

func (o *Order) GetOrderById(id int, db *sql.DB) *Order {
	sqlString := fmt.Sprintf("SELECT o.ID_order, o.client_name, o.phone FROM `orders` o WHERE o.ID_order = %d", id)
	row := db.QueryRow(sqlString)
	err := row.Scan(&o.ID, &o.ClientName, &o.Phone)
	if err != nil {
		log.Fatal(err)
	}
	return o
}

func (p *Product) GetProductById(id int, db *sql.DB) *Product {
	sqlString := fmt.Sprintf("SELECT p.ID_product, p.name, p.price FROM `products` p WHERE p.ID_product = %d", id)
	row := db.QueryRow(sqlString)
	err := row.Scan(&p.ID, &p.Name, &p.Price)
	if err != nil {
		log.Fatal(err)
	}
	return p
}

func (i *FullInfo) GetProductByIdWithBox(id int, db *sql.DB) *FullInfo {
	sqlString := fmt.Sprintf("SELECT p.ID_product, p.name, p.price, b.ID_box, b.name FROM `products` p JOIN `box-products` bp ON p.ID_product = bp.ID_products JOIN box b ON bp.ID_box = b.ID_box WHERE p.ID_product = %d", id)
	row := db.QueryRow(sqlString)
	err := row.Scan(&i.prod.ID, &i.prod.Name, &i.prod.Price, &i.box.ID, &i.box.Name)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func (i *FullInfo) GetOrderByIdWithProduct(id int, db *sql.DB) *FullInfo {
	sqlString := fmt.Sprintf("SELECT o.ID_order, o.client_name, o.phone, p.ID_product, p.name, p.price FROM orders o JOIN `orders-products` op ON o.ID_order = op.ID_order JOIN products p ON op.ID_product = p.ID_product WHERE o.ID_order = %d", id)
	row := db.QueryRow(sqlString)
	err := row.Scan(&i.order.ID, &i.order.ClientName, &i.order.Phone, &i.prod.ID, &i.prod.Name, &i.prod.Price)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func (i *FullInfo) GetAllUsedDataByProdID(id int, db *sql.DB) *FullInfo {
	sqlString := fmt.Sprintf(
		"SELECT T1.name, T1.price, T1.box_name, T2.client_name, T2.phone FROM (SELECT p.ID_product, p.name, p.price, b.name AS box_name FROM products p JOIN `box-products` bp ON p.ID_product = bp.ID_products JOIN box b ON bp.ID_box = b.ID_box) T1,(SELECT p.ID_product, o.client_name, o.phone FROM orders o JOIN `orders-products` op ON o.ID_order = op.ID_order JOIN products p ON op.ID_product = p.ID_product WHERE o.ID_order = '%d') T2 WHERE  T1.ID_product = T2.ID_product", id)
	row := db.QueryRow(sqlString)
	err := row.Scan(&i.prod.Name, &i.prod.Price, &i.box.Name, &i.order.ClientName, &i.order.Phone)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func main() {
	var inf FullInfo

	db, err := sql.Open("sqlite3", "file:test")
	if err != nil {
		log.Fatal(err)
	}

	OrderIds := []int{1, 2, 3}

	for i := 1; i < len(OrderIds)+1; i++ {
		fmt.Println(i)
		inf.GetAllUsedDataByProdID(i, db)
		fmt.Println(inf)
	}

	defer db.Close()
}
