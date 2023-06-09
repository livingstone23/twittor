package bd

import (
	"context"
	"git/twittor/models"
	"time"
)

/*BorroRelacion funcion que permite el borrado de la relacion*/
func BorroRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittordb")
	col := db.Collection("relacion")

	_, err := col.DeleteOne(ctx, t)

	if err != nil {
		return false, err
	}

	return true, nil

}
