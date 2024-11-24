package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Name       string  `json:"name" gorm:"size:100;not null"`             // Nombre con un límite de 100 caracteres y no nulo
	Email      string  `json:"email" gorm:"unique;not null"`              // Email único y obligatorio
	Password   string  `json:"password" gorm:"not null"`                  // Contraseña no nula
	Money      float64 `json:"money"`                                     // Manejamos el dinero como float
	CardNumber string  `json:"cardNumber" gorm:"unique;size:16;not null"` // Guardar número de tarjeta, límite de 16 caracteres
}
