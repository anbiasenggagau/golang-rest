package model

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func ConnectDatabase() {
	var postgres_user = "postgres"
	var postgres_pass = "supersecret"
	var postgres_port = "5432"
	var postgres_db = "postgres"

	godotenv.Load(".env")

	envUser := os.Getenv("POSTGRES_USER")
	envPass := os.Getenv("POSTGRES_PASS")
	envPort := os.Getenv("POSTGRES_PORT")
	envDb := os.Getenv("POSTGRES_USER")

	if envUser != "" {
		postgres_user = envUser
	}
	if envPass != "" {
		postgres_pass = envPass
	}
	if envDb != "" {
		postgres_db = envDb
	}
	if envPort != "" {
		postgres_port = envPort
	}

	dsn := fmt.Sprintf(`postgres://%s:%s@localhost:%s/%s`, postgres_user, postgres_pass, postgres_port, postgres_db)
	database, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&Customer{}, &CustomerAddress{}, &Product{}, &PaymentMethod{}, &Transaction{})

	// Seeding Database
	var countProduct int64
	var countPayment int64
	var countCustomer int64
	var countCustomerAddress int64

	database.Model(&Product{}).Count(&countProduct)
	database.Model(&PaymentMethod{}).Count(&countPayment)
	database.Model(&Customer{}).Count(&countCustomer)
	database.Model(&CustomerAddress{}).Count(&countCustomerAddress)

	if countProduct == 0 {
		database.Exec(`
        INSERT INTO products(name,price) VALUES('Sabun',12000);
        INSERT INTO products(name,price) VALUES('Shampoo',30000);
        INSERT INTO products(name,price) VALUES('Buku',10000);
        `)
	}

	if countPayment == 0 {
		database.Exec(`
        INSERT INTO payment_methods(name, is_active) VALUES('Cash',true);
        INSERT INTO payment_methods(name, is_active) VALUES('Online',false);
        `)
	}

	if countCustomer == 0 && countCustomerAddress == 0 {
		database.Exec(`
	    INSERT INTO customers(customer_name) VALUES('Anbia');
	    INSERT INTO customers(customer_name) VALUES('Senggagau');

        INSERT INTO customer_addresses(address, customer_id) VALUES('Semarang',1);
        INSERT INTO customer_addresses(address, customer_id) VALUES('Jakarta',2);
	    `)
	}

	Db = database
}
