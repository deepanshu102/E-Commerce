package model

func (product Product) Add() Product {
	db, _ := ConnectionStablish()
	var category Category
	db.First(&category, product.Category.ID)
	product.Category = &category
	db.Create(&product)
	db.Where("name=?", product.Name).Find(&product)
	return product

}

func (product Product) Update() Product {
	db, _ := ConnectionStablish()
	db.Save(&product)
	db.Find(&product, product.ID)
	return product
}

func (product Product) Delete(Id int) {
	db, _ := ConnectionStablish()
	db.Delete(&product, Id)

}
func (User Product) ViewAll() []Product {
	db, _ := ConnectionStablish()
	var ProductList []Product
	db.Preload("Category").Find(&ProductList)
	return ProductList
}

func (product Product) View(Id int) Product {
	db, _ := ConnectionStablish()
	db.First(&product, Id)
	return product

}
