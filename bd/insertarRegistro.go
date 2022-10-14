package bd

import (
	"context"
	"time"

	"github.com/IsmaEstrella/Twitor/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertarRegistro(u models.Usuario) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("Twitor")
	col := db.Collection("usuarios")

	//Aqui el pass se va a reescribir y se va a encriptar
	u.Password, _ = EncriptarPass(u.Password)

	result, err := col.InsertOne(ctx, u) //InsertOne devuelve un ID
	if err != nil {
		return "", false, err
	}
	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
