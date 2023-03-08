package entity

import "errors"

//create a struc to define the properties and its types

type Order struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

// func to create newOrder.
// 1-Receive properites by param,
// 2-assign the properties to the Order memory position
// 3-create err and call validate func
// if err is not nil, return nil, err
// else, return order, nil

//We dont have try catch to validate errors, so we need to validate the err and if the err is nil we have error and if
//the err is empty we have the order, always return the value wanted + error info like <order, nil>

func NewOrder(id string, price float64, tax float64) (*Order, error) {

	order := &Order{
		ID:    id,
		Price: price,
		Tax:   tax,
	}

	err := order.Validate()

	if err != nil {
		return nil, err
	}

	return order, nil
}

//func to validate if the order properties are good

func (o *Order) Validate() error {
	if o.ID == "" {
		return errors.New("id is empty")
	}

	if o.Price <= 0 {
		return errors.New("invalid price")
	}

	if o.Tax <= 0 {
		return errors.New("invalid tax")
	}

	return nil
}

// func to Calculate the final price as price + tax values, we return an error and if not have error return nil

func (o *Order) CalculateFinalPrice() error {
	o.FinalPrice = o.Price + o.Tax

	err := o.Validate()

	if err != nil {
		return err
	}

	return nil
}
