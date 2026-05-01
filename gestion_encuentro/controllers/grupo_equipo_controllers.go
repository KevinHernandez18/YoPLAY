package controllers

import (
	"encoding/json"
	"gestion_encuentro/config"
	"gestion_encuentro/models"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAllGrupoEquipo(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query("SELECT id_grupo_equipo, id_grupo, id_equipo, partidos_jugados, empates, victorias, derrotas, activo, fecha_creacion, fecha_modificacion FROM grupo_equipo")

	if err != nil {
		respondJSON(w, 500, map[string]string{"Error:": err.Error()})
		return
	}
	defer rows.Close()

	var list []models.Grupo_Equipo

	for rows.Next() {
		var GEQ models.Grupo_Equipo
		rows.Scan(&GEQ.Id_Grupo_Equipo, &GEQ.Id_Grupo, &GEQ.Id_Equipo, &GEQ.Partidos_Jugados, &GEQ.Empates, &GEQ.Victorias, &GEQ.Derrotas, &GEQ.Activo, &GEQ.Fecha_Creacion, &GEQ.Fecha_Modificacion)
		list = append(list, GEQ)
	}

	respondJSON(w, 200, list)
}

func GetGrupoEquipoByID(w http.ResponseWriter, r *http.Request) {
	id_grupo_equipo := mux.Vars(r)["id"]

	var GEQ models.Grupo_Equipo

	err := config.DB.QueryRow(
		"SELECT id_grupo_equipo, id_grupo, id_equipo, partidos_jugados, empates, victorias, derrotas, activo, fecha_creacion, fecha_modificacion FROM grupo_equipo WHERE id_grupo_equipo = $1", id_grupo_equipo,
	).Scan(&GEQ.Id_Grupo_Equipo, &GEQ.Id_Grupo, &GEQ.Id_Equipo, &GEQ.Partidos_Jugados, &GEQ.Empates, &GEQ.Victorias, &GEQ.Derrotas, &GEQ.Activo, &GEQ.Fecha_Creacion, &GEQ.Fecha_Modificacion)

	if err != nil {
		respondJSON(w, 404, map[string]string{"Error:": "Id no encontrado"})
		return
	}

	respondJSON(w, 200, GEQ)
}

func CreateGrupoEquipo(w http.ResponseWriter, r *http.Request) {
	var GEQ models.Grupo_Equipo
	json.NewDecoder(r.Body).Decode(&GEQ)

	err := config.DB.QueryRow(
		"INSERT INTO grupo_equipo (id_grupo, id_equipo) VALUES ($1, $2) RETURNING id_grupo_equipo",
		GEQ.Id_Grupo, GEQ.Id_Equipo,
	).Scan(&GEQ.Id_Grupo_Equipo)

	if err != nil {
		respondJSON(w, 500, map[string]string{"Error:": err.Error()})
		return
	}

	respondJSON(w, 201, GEQ)
}

func UpdateGrupoEquipo(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var GEQ models.Grupo_Equipo
	json.NewDecoder(r.Body).Decode(&GEQ)

	_, err := config.DB.Exec(
		"UPDATE grupo_equipo SET id_grupo=$1, id_equipo=$2, partidos_jugados=$3, empates=$4, victorias=$5, derrotas=$6, activo=true, fecha_modificacion=now() WHERE id_grupo_equipo=$7",
		GEQ.Id_Grupo, GEQ.Id_Equipo, GEQ.Partidos_Jugados, GEQ.Empates, GEQ.Victorias, GEQ.Derrotas, id,
	)

	if err != nil {
		respondJSON(w, 500, map[string]string{"Error:": err.Error()})
		return
	}
	respondJSON(w, 200, map[string]string{"Message": "Dato Actualizado"})
}

func DeleteGrupoEquipo(w http.ResponseWriter, r *http.Request){
	id := mux.Vars(r)["id"]

	_, err := config.DB.Exec("DELETE FROM grupo_equipo WHERE id_grupo_equipo=$1", id)

	if err != nil {
		respondJSON(w, 500, map[string]string{"Error": err.Error()})
		return
	}
	respondJSON(w, 200, map[string]string{"Message": "Dato eliminado"})
}

// //	Finalizado.
