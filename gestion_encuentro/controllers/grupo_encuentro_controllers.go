package controllers

import (
	//"encoding/json"
	"gestion_encuentro/config"
	"gestion_encuentro/models"
	"net/http"

	"github.com/gorilla/mux"
	//"github.com/gorilla/mux" //Librería para crear rutas
)

func GetAllGrupoencuentro(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query("SELECT id_grupo_encuentro, id_grupo, id_encuentro, activo FROM grupo_encuentro")

	if err != nil {
		respondJSON(w, 500, map[string]string{"Error:": err.Error()})
		return
	}
	defer rows.Close()

	var list []models.Grupo_Encuentro

	for rows.Next() {
		var GE models.Grupo_Encuentro
		rows.Scan(&GE.Id_Grupo_Encuentro, &GE.Id_Grupo, &GE.Id_Encuentro, &GE.Activo)
		list = append(list, GE)
	}

	respondJSON(w, 200, list)
}

func GetGrupoEncuentroByID(w http.ResponseWriter, r *http.Request) {
	id_grupo_encuentro := mux.Vars(r)["id"]

	var ge models.Grupo_Encuentro

	err := config.DB.QueryRow(
		"SELECT id_grupo_encuentro, id_grupo, id_encuentro, activo FROM grupo_encuentro WHERE id_grupo_encuentro = $1", id_grupo_encuentro,
	).Scan(&ge.Id_Grupo_Encuentro, &ge.Id_Grupo, &ge.Id_Encuentro, &ge.Activo)

	if err != nil {
		respondJSON(w, 404, map[string]string{"Error:": "Id no encontrado"})
		return
	}

	respondJSON(w, 200, ge)
}
