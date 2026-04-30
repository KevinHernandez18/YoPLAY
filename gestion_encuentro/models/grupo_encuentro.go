package models

type Grupo_Encuentro struct {
	Id_Grupo_Encuentro int 	`json:"id_grupo_encuentro"`
	Id_Grupo 		   int 	`json:"id_grupo"`
	Id_Encuentro 	   int 	`json:"id_encuentro"`
	Activo 			   bool `json:"activo"`
	Fecha_Creacion     string      `json:"fecha_creacion"`
	Fecha_Modificacion string      `json:"fecha_modificacion"`
}
