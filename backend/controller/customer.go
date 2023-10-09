package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/Kami0rn/SoyJuu/entities"
)


//POST /customer
func CreateCustomer(c *gin.Context) {
	var customer entities.Customer
	
	//bind
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}	

	// Create a new customer entity with the given customer information.
	u := entities.Customer{
		// FirstName: customer.FirstName,
		// LastName: customer.LastName,
		// Username: customer.Username,
		// Password: customer.Password,
		// Address: customer.Address,
		// Email: customer.Email,
		// Phone: customer.Phone,
		// Gender: customer.Gender,
	}

	if err := entities.DB().Create(&u).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": u})

}

// GET /user/:id
func GetCustomer(c *gin.Context) {
	var customer entities.Customer
	id := c.Param("id")
	if err := entities.DB().Preload("Gender").Raw("SELECT * FROM customers WHERE id = ?", id).Find(&customer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": customer})
}

// GET /users
func ListCustomer(c *gin.Context) {
	var customer []entities.Customer
	if err := entities.DB().Preload("Gender").Raw("SELECT * FROM customers").Find(&customer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": customer})
}

func DeleteCustomer(c *gin.Context) {
	id := c.Param("id")
	if tx := entities.DB().Exec("DELETE FROM customers WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}


// PATCH /users
func UpdateCustomer(c *gin.Context) {
	var customer entities.Customer
	var result entities.Customer

	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// ค้นหา user ด้วย id
	if tx := entities.DB().Where("id = ?", customer.ID).First(&result); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	if err := entities.DB().Save(&customer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": customer})
}