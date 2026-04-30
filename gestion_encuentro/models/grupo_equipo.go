package models

type Grupo_Equipo struct {
	Id_Grupo_Equipo    int         `json:"id_grupo_equipo"`
	Id_Grupo 		   int		   `json:"id_grupo"`
	Id_Equipo		   int		   `json:"id_equipo"`
	Partidos_Jugados   int		   `json:"partidos_jugados"`
	Empates			   int		   `json:"empates"`
	Victorias		   int		   `json:"victorias"`
	Derrotas 		   int		   `json:"derrotas"`
	Activo			   bool		   `json:"activo"`
	Fecha_Creacion	   string	   `json:"fecha_creacion"`
	Fecha_Modificacion string	   `json:"fecha_modificacion"`
}