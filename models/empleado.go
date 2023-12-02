package models

/*
es mejor conservar un estándar entre las etiquetas de json y db para no tener problemas al parsear
de json a db en el método ActualizarUnEmpleado
*/
type Empleado struct {
	Id           int64  `db:"id" json:"id"`
	Nombre       string `db:"nombre" json:"nombre"`
	Edad         uint   `db:"edad" json:"edad"`
	Nacionalidad string `db:"nacionalidad" json:"nacionalidad"`
	Ciudad       string `db:"ciudad" json:"ciudad"`
	Casado       bool   `db:"casado" json:"casado"`
	Discapacidad bool   `db:"discapacidad" json:"discapacidad"`
}
