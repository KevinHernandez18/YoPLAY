package routes

import(
	"gestion_encuentro/controllers"
	
	"github.com/gorilla/mux"
)

func RegisterGrupoEquipoRoutes(r *mux.Router){
	r.HandleFunc("/grupoencuentro", controllers.GetAllGrupoEquipo).Methods("GET")
	//r.HandleFunc("/grupoencuentro/{id}", controllers.GetGrupoEquipoByID).Methods("GET")
	//r.HandleFunc("/grupoencuentro", controllers.CreateGrupoEquipo).Methods("POST")
	//r.HandleFunc("/grupoencuentro/{id}", controllers.UpdateGrupoEquipo).Methods("PUT")
	//r.HandleFunc("/grupoencuentro/{id}", controllers.DeleteGrupoEquipo).Methods("DELETE")
}