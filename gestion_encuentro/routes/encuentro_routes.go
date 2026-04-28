package routes

import(
	"API_GO_CRUD/controllers"
	
	"github.com/gorilla/mux"
)

func RegisterCredentialRoutes(r *mux.Router){
	r.HandleFunc("/encuentro", controllers.GetALLEncuentros).Methods("GET")
	r.HandleFunc("/encuentro/{id}", controllers.GetEncuentroByID).Methods("GET")
	r.HandleFunc("/encuentro", controllers.CreateEncuentro).Methods("POST")
	r.HandleFunc("/encuentro/{id}", controllers.UpdateEncuentro).Methods("PUT")
	r.HandleFunc("/encuentro/{id}", controllers.DeleteEncuentro).Methods("DELETE")
}