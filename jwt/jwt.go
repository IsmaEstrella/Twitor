package jwt

import (
	jwt "github.com/dgrijalva/jwt-go"

	"github.com/IsmaEstrella/Twitor/models"
)

func GeneroJWT(t models.Usuario) (string, error) {
	miClave := []byte("DesarrolloMaster")

	payload := jwt.MapClaims{
		"email":     t.Email,
		"nombre":    t.Nombre,
		"apellidos": t.Apellidos,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokeStr, err := token.SignedString(miClave)
	if err != nil {
		return tokeStr, err
	}

	return tokeStr, nil
}
