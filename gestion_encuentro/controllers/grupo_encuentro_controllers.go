package controllers

import (
	//"encoding/json"
	"encoding/json"
	"gestion_encuentro/config"
	"gestion_encuentro/models"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAllGrupoEncuentro(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query("SELECT id_grupo_encuentro, id_grupo, id_encuentro, activo, fecha_creacion, fecha_modificacion FROM grupo_encuentro")

	if err != nil {
		respondJSON(w, 500, map[string]string{"Error:": err.Error()})
		return
	}
	defer rows.Close()

	var list []models.Grupo_Encuentro

	for rows.Next() {
		var GE models.Grupo_Encuentro
		rows.Scan(&GE.Id_Grupo_Encuentro, &GE.Id_Grupo, &GE.Id_Encuentro, &GE.Activo, &GE.Fecha_Creacion, &GE.Fecha_Modificacion)
		list = append(list, GE)
	}

	respondJSON(w, 200, list)
}

func GetGrupoEncuentroByID(w http.ResponseWriter, r *http.Request) {
	id_grupo_encuentro := mux.Vars(r)["id"]

	var GE models.Grupo_Encuentro

	err := config.DB.QueryRow(
		"SELECT id_grupo_encuentro, id_grupo, id_encuentro, activo, fecha_creacion, fecha_modificacion FROM grupo_encuentro WHERE id_grupo_encuentro = $1", id_grupo_encuentro,
	).Scan(&GE.Id_Grupo_Encuentro, &GE.Id_Grupo, &GE.Id_Encuentro, &GE.Activo, &GE.Fecha_Creacion, &GE.Fecha_Modificacion)

	if err != nil {
		respondJSON(w, 404, map[string]string{"Error:": "Id no encontrado"})
		return
	}

	respondJSON(w, 200, GE)
}

func CreateGrupoEncuentro(w http.ResponseWriter, r *http.Request){
	var GE models.Grupo_Encuentro
	json.NewDecoder(r.Body).Decode(&GE)

	err := config.DB.QueryRow(
		"INSERT INTO grupo_encuentro (id_grupo, id_encuentro) VALUES ($1, $2) RETURNING id_grupo_encuentro",
		GE.Id_Grupo, GE.Id_Encuentro,
	).Scan(&GE.Id_Grupo_Encuentro)

	if err != nil {
		respondJSON(w, 500, map[string]string{"Error:": err.Error()})
		return
	}

	respondJSON(w, 201, GE)
}

func UpdateGrupoEncuentro(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var GE models.Grupo_Encuentro
	json.NewDecoder(r.Body).Decode(&GE)

	_, err := config.DB.Exec(
		"UPDATE grupo_encuentro SET id_grupo=$1, id_encuentro=$2, activo=$3, fecha_modificacion=now() WHERE id_grupo_encuentro=$4",
		GE.Id_Grupo, GE.Id_Encuentro, GE.Activo, id,
	)

	if err != nil {
		respondJSON(w, 500, map[string]string{"Error:": err.Error()})
		return
	}
	respondJSON(w, 200, map[string]string{"Message": "Dato Actualizado"})
}

func DeleteGrupoEncuentro(w http.ResponseWriter, r *http.Request){
	id := mux.Vars(r)["id"]

	_, err := config.DB.Exec("DELETE FROM grupo_encuentro WHERE id_grupo_encuentro=$1", id)

	if err != nil {
		respondJSON(w, 500, map[string]string{"Error": err.Error()})
		return
	}
	respondJSON(w, 200, map[string]string{"Message": "Dato eliminado"})
}

//	Finalizado.