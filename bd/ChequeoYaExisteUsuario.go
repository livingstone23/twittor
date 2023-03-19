package bd

import (
	"context"
	"git/twittor/models"
	"time"

"go.mongodb.org/mongo-driver/bson"
)

/*ChequeoYaExisteUsuario: Funcion que recibe un email y chequea si el usuario ya existe*/
func ChequeoYaExisteUsuario(email string) (models.Usuario, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()


	db := MongoCN.Database("twittordb")
	col := db.Collection("usuarios")
condicion := bson.M{"email": email}

var resultado models.Usuario
	err := col.FindOne(ctx, condicion).Decode(&resultado)
	ID := resultado.ID.Hex()

	if err != nil {
		return resultado, false, ID
	}
	return resultado, true, ID
}
