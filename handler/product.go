package handler

import (
	"Online-Shop-API/model"
	"database/sql"
	"errors"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ListProduct(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get data from database
		products, err := model.SelectProduct(db)
		if err != nil {
			log.Printf("Having trouble getting data from database...%v\n", err)
			c.JSON(500, gin.H{"Error": "Have trouble"})
			return
		}
		// Give the response
		c.JSON(200, products)
		return
	}
}
func GetProducts(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: read id from url
		id := c.Param("id")

		// TODO: get data from database with id
		product, err := model.SelectProductById(db, id)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				log.Printf("Having trouble getting data from database...%v\n", err)
				c.JSON(404, gin.H{"Error": "Products not found"})
				return
			}

			log.Printf("Having trouble getting data from database...%v\n", err)
			c.JSON(500, gin.H{"Error": "Have trouble"})
			return
		}

		// TODO: give the responses
		c.JSON(200, product)
	}

}

func CreateProducts(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var product model.Product
		if err := c.Bind(&product); err != nil {
			log.Printf("Having trouble getting data from body...%v\n", err)
			c.JSON(400, gin.H{"Error": "Invalid request body"})
			return
		}

		product.ID = uuid.New().String()
		if err := model.InsertProduct(db, product); err != nil {
			log.Printf("Having trouble creating data from database...%v\n", err)
			c.JSON(500, gin.H{"Error": "Having Trouble in Server..."})
			return
		}

		c.JSON(201, product)
	}
}

func UpdateProducts(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var product model.Product
		if err := c.Bind(&product); err != nil {
			log.Printf("Having trouble getting data from body...%v\n", err)
			c.JSON(400, gin.H{"Error": "Invalid request body"})
			return
		}

		productExisting, err := model.SelectProductById(db, id)
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("Product not found...%v\n", id)
			c.JSON(404, gin.H{"Error": "Product not found"})
			return
		} else if err != nil {
			log.Printf("Having trouble getting data from database...%v\n", err)
			c.JSON(500, gin.H{"Error": "Having Trouble in Server..."})
			return
		} else if productExisting == (model.Product{}) {
			log.Printf("Product not found...%v\n", id)
			c.JSON(404, gin.H{"Error": "Product not found"})
			return
		}
		if product.Name != "" {
			productExisting.Name = product.Name
		}

		if product.Price != 0 {
			productExisting.Price = product.Price
		}

		if err := model.UpdateProducts(db, productExisting); err != nil {
			log.Printf("Having trouble to update data from database...%v\n", err)
			c.JSON(500, gin.H{"Error": "Having Trouble in Server..."})
			return
		}

		c.JSON(201, product)
	}
}

func DeleteProducts(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		if err := model.DeleteProducts(db, id); err != nil {
			log.Printf("Having trouble to update data from database...%v\n", err)
			c.JSON(500, gin.H{"Error": "Having Trouble in Server..."})
			return
		}

		c.JSON(204, nil)
	}
}
