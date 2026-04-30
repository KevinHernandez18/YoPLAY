package controllers

import (
	"encoding/json"
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

func GetGrupoByID(w http.ResponseWriter, r *http.Request) {
	id_grupo := mux.Vars(r)["id"]

	var g models.Grupo

	err := config.DB.QueryRow(
		"SELECT id_grupo, grupo, activo, fecha_creacion, fecha_modificacion FROM grupo WHERE id_grupo=$1", id_grupo,
	).Scan(&g.Id_grupo, &g.Grupo, &g.Activo, &g.Fecha_Creacion, &g.Fecha_Modificacion)

	if err != nil {
		respondJSON(w, 404, map[string]string{"Error:": "Id no encontrado"})
		return
	}

	respondJSON(w, 200, g)
}

func CreateGrupo(w http.ResponseWriter, r *http.Request) {
	var g models.Grupo
	json.NewDecoder(r.Body).Decode(&g)

	err := config.DB.QueryRow(
		"INSERT INTO grupo (grupo) VALUES ($1) RETURNING id_grupo",
		g.Grupo,
	).Scan(&g.Id_grupo)

	if err != nil {
		respondJSON(w, 500, map[string]string{"Error:": err.Error()})
		return
	}

	respondJSON(w, 201, g)
}

func UpdateGrupo(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var g models.Grupo
	json.NewDecoder(r.Body).Decode(&g)

	_, err := config.DB.Exec(
		"UPDATE grupo SET grupo = $1, fecha_modificacion = now() WHERE id_grupo = $2",
		g.Grupo, id,
	)

	if err != nil {
		respondJSON(w, 500, map[string]string{"Error:": err.Error()})
		return
	}

	respondJSON(w, 200, map[string]string{"Message": "Actualizado correctamente"})
}
