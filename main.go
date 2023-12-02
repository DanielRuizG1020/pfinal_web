package main

import (
	"log"
	"net/http"

	"proyectofinal/controllers"
	"proyectofinal/handlers"
	"proyectofinal/models"
	repositorio "proyectofinal/repository" /* importando el paquete de repositorio */

	GorillaHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

/*
función para conectarse a la instancia de PostgreSQL, en general sirve para cualquier base de datos SQL.
Necesita la URL del host donde está instalada la base de datos y el tipo de base datos (driver)
*/
func ConectarDB(url, driver string) (*sqlx.DB, error) {
	pgUrl, _ := pq.ParseURL(url)
	db, err := sqlx.Connect(driver, pgUrl) // driver: postgres
	if err != nil {
		log.Printf("fallo la conexion a PostgreSQL, error: %s", err.Error())
		return nil, err
	}

	log.Printf("Nos conectamos bien a la base de datos db: %#v", db)
	return db, nil
}

func main() {
	/* creando un objeto de conexión a PostgreSQL */
	db, err := ConectarDB("postgres://wxxtppbw:C5FOPj-ozMvMaHw3FNT3yJnB6UHoNPek@flora.db.elephantsql.com/wxxtppbw", "postgres")
	if err != nil {
		log.Fatalln("error conectando a la base de datos", err.Error())
		return
	}

	/* creando una instancia del tipo Repository del paquete repository
	se debe especificar el tipo de struct que va a manejar la base de datos
	para este ejemplo es Empleado y se le pasa como parámetro el objeto de
	conexión a PostgreSQL */
	repo, err := repositorio.NewRepository[models.Empleado](db)
	if err != nil {
		log.Fatalln("fallo al crear una instancia de repositorio", err.Error())
		return
	}

	controller, err := controllers.NewController(repo)
	if err != nil {
		log.Fatalln("fallo al crear una instancia de controller", err.Error())
		return
	}

	handler, err := handlers.NewHandler(controller)
	if err != nil {
		log.Fatalln("fallo al crear una instancia de handler", err.Error())
		return
	}

	/* router (multiplexador) a los endpoints de la API (implementado con el paquete gorilla/mux) */
	router := mux.NewRouter()

	/* rutas a los endpoints de la API */
	router.Handle("/empleados", http.HandlerFunc(handler.LeerEmpleados)).Methods(http.MethodGet)
	router.Handle("/empleados", http.HandlerFunc(handler.CrearEmpleado)).Methods(http.MethodPost)
	router.Handle("/empleados/{id}", http.HandlerFunc(handler.LeerUnEmpleado)).Methods(http.MethodGet)
	router.Handle("/empleados/{id}", http.HandlerFunc(handler.ActualizarUnEmpleado)).Methods(http.MethodPatch)
	router.Handle("/empleados/{id}", http.HandlerFunc(handler.EliminarUnEmpleado)).Methods(http.MethodDelete)

	//Solucion problema CORS

	corsOptions := GorillaHandlers.CORS(
		GorillaHandlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		GorillaHandlers.AllowedMethods([]string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}),
		GorillaHandlers.AllowedOrigins([]string{"*"}), // Permitir solicitudes desde cualquier origen (*)
	)

	// Usa el manejador CORS para todas las rutas
	Gorillahandler := corsOptions(router)

	// Inicia el servidor
	log.Fatal(http.ListenAndServe(":8080", Gorillahandler))
	/* servidor escuchando en localhost por el puerto 8080 y entrutando las peticiones con el router
	headers := GorillaHandlers.AllowedHeaders([]string{"X-Requested-With", "Content Type", "Authorization"})
	methods := GorillaHandlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "PATCH", "HEAD", "OPTIONS"})
	origins := GorillaHandlers.AllowedOrigins([]string{"*"})
	http.ListenAndServe(":8080", GorillaHandlers.CORS(headers, methods, origins)(router))*/
}
