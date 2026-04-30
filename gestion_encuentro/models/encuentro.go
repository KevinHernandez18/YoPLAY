package models

import "time"

type Resultado string

const (
	Ganador  Resultado = "Ganador"
	Perdedor Resultado = "Perdedor"
	Empate   Resultado = "Empate"
)

type Fase_Torneo string

const (
	Grupos    Fase_Torneo = "Grupos"
	Octavos   Fase_Torneo = "Octavos"
	Cuartos   Fase_Torneo = "Cuartos"
	Semifinal Fase_Torneo = "Semifinal"
	Final     Fase_Torneo = "Final"
)

type Encuentro struct {
	Id_Encuentro         int         `json:"id_encuentro"`
	Fecha                time.Time   `json:"fecha"`
	Id_Torneo            int         `json:"id_torneo"`
	Id_tipo_distribucion int         `json:"id_tipo_distribucion"`
	Fase_torneo          Fase_Torneo `json:"fase_torneo"`
	Id_Equipo1           int         `json:"id_equipo1"`
	Resultado_Equipo1    Resultado   `json:"resultado_equipo1"`
	Id_Equipo2           int         `json:"id_equipo2"`
	Resultado_Equipo2    Resultado   `json:"resultado_equipo2"`
	Activo               bool        `json:"activo"`
	Fecha_Creacion       string      `json:"fecha_creacion"`
	Fecha_Modificacion   string      `json:"fecha_modificacion"`
}
