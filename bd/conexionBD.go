package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Creamos dos variable una se va a exportar y la otra de uso interno
// mongoC Es donde es la conexion de la base de datos.
var MongoCN = ConectarBD()
var clientOpions = options.Client().ApplyURI("mongodb+srv://ismaduarte:Iduarte16@cluster0.dvax7ee.mongodb.net/?retryWrites=true&w=majority")

func ConectarBD() *mongo.Client {
	//Conexion a la base de datos y tiene un contexto en el cual va a trabajar
	client, err := mongo.Connect(context.TODO(), clientOpions)
	//Declaramos un if para la validar la conexion.
	if err != nil {
		log.Fatal(err.Error()) //.Error() su funcion es convertir un objeto a String
		return client
	}
	//Realizamos una verificacion de comunicacion con la base de datos usando PING
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexion Exitosa con la Base de datos")
	return client
}

/*checarConexion del servirod*/
func ChecarConexion() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
