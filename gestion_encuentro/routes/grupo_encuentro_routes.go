package routes

import(
	"gestion_encuentro/controllers"
	
	"github.com/gorilla/mux"
)

func RegisterGrupoEncuentroRoutes(r *mux.Router){
	r.HandleFunc("/grupoencuentro", controllers.GetAllGrupoencuentro).Methods("GET")
	//r.HandleFunc("/grupoencuentro/{id}", controllers.GetGrupoEncuentroByID).Methods("GET")
	//r.HandleFunc("/grupoencuentro", controllers.CreateGrupoEncuentro).Methods("POST")
	//r.HandleFunc("/grupoencuentro/{id}", controllers.UpdateGrupoEncuentro).Methods("PUT")
	//r.HandleFunc("/grupoencuentro/{id}", controllers.DeleteGrupoEncuentro).Methods("DELETE")
}