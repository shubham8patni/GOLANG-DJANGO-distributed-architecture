package models

import "gorm.io/gorm"

type Brands struct {
	gorm.Model
	Name string `gorm:"column:name;not null;uniqueIndex"`
	//Products []Products // One-to-Many relationship
}

//DUMMY DATA QUERY
//brands query:
//INSERT INTO brands (name) VALUES ('Brand Name 1'),       ('Brand Name 2'),       ('Brand Name 3'),       ('Brand Name 4'),       ('Brand Name 5'),       ('Brand Name 6'),       ('Brand Name 7'),       ('Brand Name 8');
//

//CREATE TABLE IF NOT EXISTS brands (
//id INT PRIMARY KEY AUTO_INCREMENT,  -- Auto-incrementing primary key
//name VARCHAR(255) UNIQUE NOT NULL  -- Brand name (unique, up to 255 characters)
//);
