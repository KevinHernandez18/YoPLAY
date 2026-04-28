package controllers

import (
	"encoding/json"
	"gestion_encuentro/config"
	"gestion_encuentro/models"
	"net/http"

	"github.com/gorilla/mux"//Librería para crear rutas
)

func respondJSON(w http.ResponseWriter, status int, payload interface{}){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

func GetAllEncuentros(w http.ResponseWriter, r *http.Request){
	rows, err := config.DB.Query("SELECT id_encuentro, fecha, id_torneo, id_tipo_distribucion, fase_torneo, id_equipo1, resultado_equipo1, id_equipo2, resultado_equipo2, activo, fecha_creacion, fecha_modificacion FROM encuentro")

	if err != nil {
		respondJSON(w, 500, map[string]string{"Error": err.Error()})
		return
	}
	defer rows.Close()

	var list []models.Encuentro

	for rows.Next() {
		var e models.Encuentro
		rows.Scan(&e.Id_Encuentro, &e.Fecha, &e.Id_Torneo, &e.Id_tipo_distribucion, &e.Fase_torneo, &e.Id_Equipo1, &e.Resultado_Equipo1, &e.Id_Equipo2, e.Resultado_Equipo2, e.Activo, e.Fecha_Creacion, e.Fecha_Modificacion)
		list = append(list, e)
	}

	respondJSON(w, 200, list)
}
