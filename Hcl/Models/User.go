package Models

import (
	"Hcl/Config"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//GetAllUsers Fetch all user data
func GetAllHotels(hotel *[]Hotel) (err error) {
	if err = Config.DB.Find(hotel).Error; err != nil {
		return err
	}
	return nil
}

//CreateUser ... Insert New data
func CreateHotel(hotel *Hotel) (err error) {
	if err = Config.DB.Create(&hotel).Error; err != nil {
		return err
	}
	return nil
}

//GetUserByID ... Fetch only one user by Id
func GetHotelByID(hotel *Hotel, Hotel_Id string) (err error) {
	if err = Config.DB.Where("Hotel_Id = ?", Hotel_Id).First(hotel).Error; err != nil {
		return err
	}
	return nil
}

//UpdateUser ... Update user
func UpdateHotel(hotel *Hotel, Hotel_Id string) (err error) {
	fmt.Println(hotel)
	Config.DB.Save(hotel)
	return nil
}

//DeleteUser ... Delete user
func DeleteHotel(hotel *Hotel, Hotel_Id string) (err error) {
	Config.DB.Where("Hotel_Id = ?", Hotel_Id).Delete(hotel)
	return nil
}
