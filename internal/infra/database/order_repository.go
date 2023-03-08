package database

import (
	"database/sql"

	"github.com/herlloncardoso/go-orders/internal/entity"
)

//struct using connection with the database from a go native package

type OrderRepository struct {
	Db *sql.DB
}

//create constructor function

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{Db: db}
}

func (r *OrderRepository) Save(order *entity.Order) error {

	//do the insert into orders, the values marked with ? will be replaced with the variables in sequence of the comma
	_, err := r.Db.Exec("Insert into orders (id, price, tax, final_price) Values(?,?,?,?)",
		order.ID, order.Price, order.Tax, order.FinalPrice)

	if err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) GetTotal() (int, error) {
	//declare variable total
	var total int

	//do the select and count of the order table and use SCAN for save the result in the variable, in that case we
	//change the value of the variable pointing for the memory position
	err := r.Db.QueryRow("SELECT count(*) from orders ").Scan(&total)

	if err != nil {
		return 0, err
	}

	return total, nil
}
