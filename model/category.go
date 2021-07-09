package model

func (category Category) Add() Category {

	db, _ := ConnectionStablish()
	db.Create(&category)
	db.Where("name=?", category.Name).Find(&category)
	return category
}

func (Category Category) Update() Category {
	db, _ := ConnectionStablish()
	db.Save(&Category)
	db.Find(&Category, Category.ID)
	return Category

}

func (category Category) Delete(Id int) Category {
	db, _ := ConnectionStablish()
	db.Delete(&category, Id)
	return category
}
func (category Category) ViewAll() []Category {
	var CategoryList []Category
	db, _ := ConnectionStablish()
	db.Preload("ProductList").Find(&CategoryList)
	return CategoryList
}

func (tempcategory Category) View(name string) Category {

	db, _ := ConnectionStablish()
	db.Where("name=?", name).First(&tempcategory)
	return tempcategory

}
