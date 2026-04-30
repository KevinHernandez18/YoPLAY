package controllers

import (
	//"encoding/json"
	"gestion_encuentro/config"
	"gestion_encuentro/models"
	"net/http"

	"github.com/gorilla/mux" //Librería para crear rutas
)

func GetALLGrupos(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query("SELECT id_grupo, grupo, activo, fecha_creacion, fecha_modificacion FROM grupo")

	if err != nil {
		respondJSON(w, 500, map[string]string{"Error": err.Error()})
		return
	}
	defer rows.Close()

	var list []models.Grupo

	for rows.Next() {
		var g models.Grupo
		rows.Scan(&g.Id_grupo, &g.Grupo, &g.Activo, &g.Fecha_Creacion, &g.Fecha_Modificacion)
		list = append(list, g)
	}

	respondJSON(w, 200, list)
}


