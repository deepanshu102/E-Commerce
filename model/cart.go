package model

func (cart Cart) Add(UserId int) []Cart {
	var User Users
	var Product Product
	var TCart Cart
	var CartList []Cart
	db, _ := ConnectionStablish()
	db.Find(&User, UserId)
	db.Find(&Product, cart.ProductId)
	TCart.Quantity = cart.Quantity
	TCart.Amount = float64(cart.Quantity) * Product.Price
	TCart.Product = &Product
	TCart.User = &User
	db.Create(&TCart)
	db.Find(&CartList)

	return CartList
}

func (cart Cart) Delete(Id int) {
	db, _ := ConnectionStablish()
	db.Delete(&cart, Id)

}
func (cart Cart) Update() []Cart {
	var CartList []Cart
	var TCart Cart
	db, _ := ConnectionStablish()
	db.Preload("Product").Find(&TCart, cart.ID)
	TCart.Quantity = cart.Quantity
	TCart.Amount = float64(cart.Quantity) * TCart.Product.Price
	db.Save(&TCart)
	db.Find(&CartList)
	return CartList

}

func (cart Cart) View(UserId int) []Cart {
	db, _ := ConnectionStablish()
	var User Users
	db.Preload("CartList").Find(&User, UserId)
	return User.CartList
}
