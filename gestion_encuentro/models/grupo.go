package models

type Grupo struct {
	Id_grupo 		   int 		   `json:"id_grupo"`
	Grupo 			   string	   `json:"grupo"`
	Activo 			   bool 	   `json:"activo"`
	Fecha_Creacion     string      `json:"fecha_creacion"`
	Fecha_Modificacion string      `json:"fecha_modificacion"`
}
