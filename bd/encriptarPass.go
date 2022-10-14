package bd

import (
	"golang.org/x/crypto/bcrypt"
)

func EncriptarPass(pass string) (string, error) {
	//Se recomienda utilizar para un pass comun usar 6 y si utilizan la pass para un super usuario utilizar 8
	costo := 8 //Es el algorimo basado elevado al costo al cuadrado, a mayor costo es una mejor encriptacion
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	return string(bytes), err
}
