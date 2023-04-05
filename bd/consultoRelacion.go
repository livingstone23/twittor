package bd

import (
	"context"
	"fmt"
	"git/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

/*ConsultoRelacion funcion que permite consultar la relacion */
func ConsultoRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittordb")
	col := db.Collection("relacion")

	condicion := bson.M{
		"usuarioId":         t.UsuarioId,
		"usuarioRelacionID": t.UsuarioRelacionID,
	}

	var resultado models.Relacion
	err := col.FindOne(ctx, condicion).Decode(&resultado)
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}

	return true, nil
}
