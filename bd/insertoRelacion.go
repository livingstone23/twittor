package bd

import (
	"context"
	"git/twittor/models"
	"time"
)

/*InsertoRelacion funcion que permite insertar la relacion*/
func InsertoRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittordb")
	col := db.Collection("relacion")

	_, err := col.InsertOne(ctx, t)

	if err != nil {
		return false, err
	}

	return true, nil
}
