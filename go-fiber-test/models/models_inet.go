package models

import "gorm.io/gorm"

type Person struct {
	Name string `json:"name"`
	Pass string `json:"pass"`
}

type User struct {
	Name     string `json:"name" validate:"required,min=3,max=32"`
	IsActive *bool  `json:"isactive" validate:"required"`
	Email    string `json:"email,omitempty" validate:"required,email,min=3,max=32"`
}

type RegisterRequest struct {
	Email        string `json:"email" validate:"required,email"`
	Username     string `json:"username" validate:"required"`
	Password     string `json:"password" validate:"required,min=6,max=20"`
	LineID       string `json:"line_id" validate:"required"`
	Tel          string `json:"tel" validate:"required"`
	BusinessType string `json:"business_type" validate:"required"`
	Website      string `json:"website"`
}

type Dogs struct {
	gorm.Model
	Name      string         `json:"name"`
	DogID     int            `json:"dog_id"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type DogsRes struct {
	Name  string `json:"name"`
	DogID int    `json:"dog_id"`
	Type  string `json:"type"`
}

type ResultData struct {
	Data       []DogsRes `json:"data"`
	Name       string    `json:"name"`
	Count      int       `json:"count"`
	SumRed     int       `json:"sum_red"`
	SumGreen   int       `json:"sum_green"`
	SumPink    int       `json:"sum_pink"`
	SumNoColor int       `json:"sum_nocolor"`
}

type Company struct {
	gorm.Model
	CompanyID    string `json:"company_id"`
	Name         string `json:"name"`
	Address      string `json:"address"`
	Phone        string `json:"phone"`
	Employees    int    `json:"employees"`
	BusinessType string `json:"business_type"`
}

type UsersProfile struct {
	gorm.Model
	EmployeeID string `json:"employee_id"`
	Name       string `json:"name"`
	Lastname   string `json:"lastname"`
	Birthday   string `json:"birthday"`
	Age        int    `json:"age"`
	Email      string `json:"email"`
	Tel        string `json:"tel"`
}
