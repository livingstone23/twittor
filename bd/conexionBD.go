package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoCN es el objeto de conexion a la base de datos*/
var MongoCN = ConectarBD()
var clientOptions = options.Client().ApplyURI("")

/*ConectarBD funcion que me permite conectarme a la base de datos*/
func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	//Hacemos un ping para saber si esta la conexion a la base de datos
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexion exitosa con la BD")
	return client
}

/*ChequeoConnection es el ping a la base de datos*/
func ChequeoConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
