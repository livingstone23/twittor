package models

/* Relacion modelo para grabar la relacion de un usuario con otro */
type Relacion struct {
	UsuarioId         string `bson:"usuarioId" json:"usuarioId"`
	UsuarioRelacionID string `bson:"usuarioRelacionID" json:"usuarioRelacionID"`
}
