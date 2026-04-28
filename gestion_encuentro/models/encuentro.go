package models

import "time"

type Encuentro struct {
	Id_Encuentro         int `json: "id_encuentro"`
	Fecha                time.Time `json: "fecha"`
	Id_Torneo            int `json: "id_torneo"`
	Id_tipo_distribucion int `json: "id_tipo_distribucion"`
	Fase_torneo 		 string `json: "fase_torneo"`
	Id_Equipo1			 int `json: "id_Equipo1"`
	Resultado_Equipo1 	 string `json: "resultado equipo1"`
	Id_Equipo2 			 int `json: "id_equipo2"`
	Resultado_Equipo2 	 string `json: "resultado_equipo2"`
	Activo 				 bool `json: "activo"`
	Fecha_Creacion 		 string `json: "fecha_creacion"`
	Fecha_Modificacion 	 string `json: "fecha_modificacion"`
}