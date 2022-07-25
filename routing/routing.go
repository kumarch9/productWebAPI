package routing

import (
	"log"
	"net/http"
	hd "productwebapi/handlers"
	"productwebapi/signin"
	"productwebapi/signup"

	"github.com/gorilla/mux"
)

func HandlerRouting() {
	route := mux.NewRouter()

	//can't used "/" after route name e.g. `www.host.com/api/route + /`
	route.StrictSlash(false)
	newRoute := route.PathPrefix("/api").Subrouter()

	newRoute.HandleFunc("/signup", signup.RegUserHandler).Methods("POST")
	newRoute.HandleFunc("/signin", signin.SigninHandler).Methods("POST")
	newRoute.HandleFunc("/product", hd.CreateProduct).Methods("POST")
	newRoute.HandleFunc("/product", hd.GetProduct).Methods("GET")
	newRoute.HandleFunc("/product/{id}", hd.GetProductByID).Methods("GET")
	newRoute.HandleFunc("/product/{id}", hd.UpdateProduct).Methods("PUT")
	newRoute.HandleFunc("/product/{id}", hd.DeleteProduct).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8055", route))

}
