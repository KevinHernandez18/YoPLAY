package routes

import(
	"gestion_encuentro/controllers"
	
	"github.com/gorilla/mux"
)

func RegisterEncuentroRoutes(r *mux.Router){
	r.HandleFunc("/encuentro", controllers.GetAllEncuentros).Methods("GET")
	//r.HandleFunc("/encuentro/{id}", controllers.GetEncuentroByID).Methods("GET")
	//r.HandleFunc("/encuentro", controllers.CreateEncuentro).Methods("POST")
	//r.HandleFunc("/encuentro/{id}", controllers.UpdateEncuentro).Methods("PUT")
	//r.HandleFunc("/encuentro/{id}", controllers.DeleteEncuentro).Methods("DELETE")
}