package bd

import (
	"context"
	"time"

	"github.com/IsmaEstrella/Twitor/models"
	"go.mongodb.org/mongo-driver/bson"
)

func UsuarioExiste(email string) (models.Usuario, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("Twitor")
	col := db.Collection("usuarios")

	codicion := bson.M{"email": email}

	var resultado models.Usuario
	err := col.FindOne(ctx, codicion).Decode(&resultado)
	ID := resultado.ID.Hex()
	if err != nil {
		return resultado, false, ID
	}
	return resultado, true, ID
}
