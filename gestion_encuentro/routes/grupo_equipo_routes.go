package routes

import(
	"gestion_encuentro/controllers"
	
	"github.com/gorilla/mux"
)

func RegisterGrupoEquipoRoutes(r *mux.Router){
	r.HandleFunc("/grupoequipo", controllers.GetAllGrupoEquipo).Methods("GET")
	r.HandleFunc("/grupoequipo/{id}", controllers.GetGrupoEquipoByID).Methods("GET")
	r.HandleFunc("/grupoequipo", controllers.CreateGrupoEquipo).Methods("POST")
	r.HandleFunc("/grupoequipo/{id}", controllers.UpdateGrupoEquipo).Methods("PUT")
	//r.HandleFunc("/grupoequipo/{id}", controllers.DeleteGrupoEquipo).Methods("DELETE")
}