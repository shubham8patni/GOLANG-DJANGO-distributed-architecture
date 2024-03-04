package models

import "gorm.io/gorm"

type Products struct {
	gorm.Model
	Name        string  `gorm:"column:name;not null"`
	Description string  `gorm:"column:description;not null"`
	Price       float64 `gorm:"column:price;not null"`
	BrandID     uint    `gorm:"column:brand_id"` // Foreign key
	Brand       Brands  `gorm:"foreignKey:BrandID;references:ID"`
	//Brand   Brands  // Define the relationship
}

//DUMMY DATA QUERY
//INSERT INTO products (name, description, price, brand_id)
//
//VALUES
//
//('Product 1 from Brand 1', 'Description of Product 1', 12.99, 1),
//
//('Product 2 from Brand 2', 'Description of Product 2', 24.50, 2),
//
//('Product 3 from Brand 3', 'Description of Product 3', 39.95, 3),
//
//('Product 3 from Brand 5', 'Description of Product 4', 40.95, 4)
//
//;

//CREATE TABLE IF NOT EXISTS products (
//id INT PRIMARY KEY AUTO_INCREMENT,  -- Auto-incrementing primary key
//name VARCHAR(255) NOT NULL,  -- Product name (up to 255 characters)
//description TEXT NOT NULL,  -- Product description
//price DECIMAL(10, 2) NOT NULL,  -- Price with 2 decimal places
//brand_id INT NOT NULL,  -- Foreign key referencing the Brands table
//FOREIGN KEY (brand_id) REFERENCES brands(id)  -- Enforce foreign key constraint
//);
