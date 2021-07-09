package model

func (User Users) Add() Users {

	db, _ := ConnectionStablish()
	db.Create(&User)
	db.Find(&User, "name=?", User.Name)
	return User
}
func (User Users) Update(Id int) Users {
	db, _ := ConnectionStablish()
	var user Users
	db.Find(&user, Id)
	User.ID = user.ID
	db.Save(&User)
	db.Find(&User, User.ID)
	return User
}
func (User Users) Delete(Id int) {
	db, _ := ConnectionStablish()
	db.Delete(&User, Id)

}

func (User Users) ViewAll() []Users {
	var UsersList []Users
	db, _ := ConnectionStablish()
	db.Find(&UsersList)
	return UsersList
}

func (User Users) View(Id int) Users {
	db, _ := ConnectionStablish()
	db.Find(&User, Id)
	return User
}

func (User Users) Login(email, pass string) Users {
	db, _ := ConnectionStablish()
	db.Find(&User, "email=? and password=?", email, pass)
	return User

}
