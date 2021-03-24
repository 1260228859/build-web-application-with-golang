package models

import (
	db "gin/databases"
)

type Order struct {
	orderNumber int64
	productCode string
}

func (o *Order) AddOrder() (id int64, err error) {
	rs, err := db.SqlDB.Exec("INSERT INTO orderdetails(productCode) VALUES (?)", o.productCode)
	if err != nil {

		return
	}
	id, err = rs.LastInsertId()
	return
}

func (o *Order) GetOrder() (orders []Order, err error) {
	rs, err := db.SqlDB.Query("SELECT orderNumber, productCode FROM orderdetails")
	defer rs.Close()
	if err != nil {
		return
	}
	orders = make([]Order, 0)
	for rs.Next() {
		var order Order
		rs.Scan(&order.orderNumber, &order.productCode)
		orders = append(orders, order)
	}

	if err = rs.Err(); err != nil {
		return
	}
	return
}
