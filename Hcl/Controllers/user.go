package Controllers

import (
	"Hcl/Models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Get Hotels ... Get all hotels
func GetHotels(c *gin.Context) {
	var hotel []Models.Hotel
	err := Models.GetAllHotels(&hotel)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, hotel)
	}
}

//CreateHotel  ...
func CreateHotel(c *gin.Context) {
	var hotel Models.Hotel
	c.BindJSON(&hotel)
	err := Models.CreateHotel(&hotel)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, hotel)
	}
}

//GetHotelByID ... Get the hotel by id
func GetHotelByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var hotel Models.Hotel
	err := Models.GetHotelByID(&hotel, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, hotel)
	}
}

// UpdateHotel ... Update the hotel information
func UpdateHotel(c *gin.Context) {
	var hotel Models.Hotel
	id := c.Params.ByName("id")
	err := Models.GetHotelByID(&hotel, id)
	if err != nil {
		c.JSON(http.StatusNotFound, hotel)
	}
	c.BindJSON(&hotel)
	err = Models.UpdateHotel(&hotel, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, hotel)
	}
}

//DeleteUser ... Delete the user
func Deletehotel(c *gin.Context) {
	var hotel Models.Hotel
	id := c.Params.ByName("id")
	err := Models.DeleteHotel(&hotel, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"Hotel_Id " + id: "is deleted"})
	}
}
