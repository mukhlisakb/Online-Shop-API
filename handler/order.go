package handler

import (
	"Online-Shop-API/model"
	"database/sql"
	"log"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func CheckOutOrder(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: get data order from request
		var checkOutOrder model.Checkout
		if err := c.BindJSON(&checkOutOrder); err != nil {
			log.Printf("Having trouble to get data from request body...%v\n", err)
			c.JSON(400, gin.H{"Error": "Products not found..."})
			return
		}
		// TODO: get product from database
		ids := []string{}
		orderQty := make(map[string]int32)
		for _, o := range checkOutOrder.Products {
			ids = append(ids, o.ID)
			orderQty[o.ID] = int32(o.Quantity)
		}

		products, err := model.SelectProductIn(db, ids)
		if err != nil {
			log.Printf("Having trouble to get data Product...%v\n", err)
			c.JSON(500, gin.H{"Error": "Having trouble in the Server.."})
			return
		}
		// TODO: create password
		passcode := generatePasscode(5)
		// TODO: hasing password
		hashcode, err := bcrypt.GenerateFromPassword([]byte(passcode), 10)
		if err != nil {
			log.Printf("Having trouble to create hash password...%v\n", err)
			c.JSON(500, gin.H{"Error": "Having trouble in the Server.."})
			return
		}

		hashcodestring := string(hashcode)
		// TODO: detail and orders

		order := model.Order{
			ID:         uuid.New().String(),
			Email:      checkOutOrder.Email,
			Address:    checkOutOrder.Address,
			Passcode:   &hashcodestring,
			GrandTotal: 0,
		}

		details := []model.OrderDetail{}

		for _, p := range products {
			total := p.Price * int64(orderQty[p.ID])

			detail := model.OrderDetail{
				ID:        uuid.New().String(),
				OrderID:   order.ID,
				ProductID: p.ID,
				Quantity:  orderQty[p.ID],
				Price:     p.Price,
				Total:     total,
			}
			details = append(details, detail)

			order.GrandTotal += total

		}
		if err := model.CreateOrder(db, order, details); err != nil {
			log.Printf("Having trouble to create Order...%v\n", err)
			c.JSON(500, gin.H{"Error": "Having trouble in the Server.."})
			return
		}

		orderWithDetail := model.OrderWithDetails{
			Order:   order,
			Details: details,
		}

		orderWithDetail.Order.Passcode = &passcode

		c.JSON(200, orderWithDetail)
	}
}

func generatePasscode(lenght int) string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	randomGenerator := rand.New(rand.NewSource(time.Now().UnixNano()))
	code := make([]byte, lenght)
	for i := range code {
		code[i] = charset[randomGenerator.Intn(len(charset))]
	}
	return string(code)
}
func Confirm(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: Getting id from params
		id := c.Param("id")

		// TODO: Read body request
		var confirmReq model.Confirm
		if err := c.BindJSON(&confirmReq); err != nil {
			log.Printf("Kesulitan mengambil data dari body request: %v\n", err)
			c.JSON(400, gin.H{"Error": "Data permintaan tidak valid"})
			return
		}

		// TODO: Get order from database
		order, err := model.SelectOrderById(db, id)
		if err != nil {
			log.Printf("Kesulitan mengambil data dari database: %v\n", err)
			c.JSON(500, gin.H{"Error": "Kesulitan pada server"})
			return
		}

		if order.Passcode == nil {
			log.Println("Passcode tidak valid")
			c.JSON(400, gin.H{"Error": "Passcode tidak valid"})
			return
		}

		// TODO: Matching Passcode
		if err := bcrypt.CompareHashAndPassword([]byte(*order.Passcode), []byte(confirmReq.Passcode)); err != nil {
			log.Printf("Data tidak memiliki izin untuk mengakses: %v\n", err)
			c.JSON(401, gin.H{"Error": "Data tidak memiliki izin untuk mengakses"})
			return
		}

		// TODO: Makesure oderd is not paid.
		if order.PaidAt != nil {
			log.Println("Order sudah dibayar")
			c.JSON(400, gin.H{"Error": "Order sudah dibayar"})
			return
		}

		// TODO: Matching amount.
		if order.GrandTotal != confirmReq.Amount {
			log.Println("Jumlah tidak valid")
			c.JSON(400, gin.H{"Error": "Jumlah tidak valid"})
			return
		}

		// TODO: Information order update.
		current := time.Now()
		if err = model.UpdateOrderByID(db, id, confirmReq, current); err != nil {
			log.Printf("Kesulitan memperbarui data di database: %v\n", err)
			c.JSON(500, gin.H{"Error": "Kesulitan pada server"})
			return
		}

		// TODO: Clean Passcode and new paid information.
		order.Passcode = nil
		order.PaidAt = &current
		order.PaidBank = &confirmReq.Bank
		order.PaidAccountNumber = &confirmReq.AccountNumber

		// TODO: Sending responses
		c.JSON(200, order)
	}
}

func GetOrder(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: Getting id from params
		id := c.Param("id")

		// TODO: Getting passcode from query params
		passcode := c.Query("passcode")

		// TODO: Get data from database
		order, err := model.SelectOrderById(db, id)
		if err != nil {
			log.Printf("Kesulitan mengambil data dari database: %v\n", err)
			c.JSON(500, gin.H{"Error": "Kesulitan pada server"})
			return
		}

		if order.Passcode == nil {
			log.Println("Passcode tidak valid")
			c.JSON(400, gin.H{"Error": "Passcode tidak valid"})
			return
		}

		// TODO: Matching passcode
		if err := bcrypt.CompareHashAndPassword([]byte(*order.Passcode), []byte(passcode)); err != nil {
			log.Printf("Data tidak memiliki izin untuk mengakses: %v\n", err)
			c.JSON(401, gin.H{"Error": "Data tidak memiliki izin untuk mengakses"})
			return
		}

		// TODO: Remove passcode and update payment information
		order.Passcode = nil

		// TODO: Sending Responses
		c.JSON(200, order)
	}
}
