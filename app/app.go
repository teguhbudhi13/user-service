package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/teguhbudhi13/user-service/app/handler"
	"github.com/teguhbudhi13/user-service/app/model"
	"github.com/teguhbudhi13/user-service/config"
)

// App has router and db instances
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

// Initialize initializes the app with predefined configuration
func (a *App) Initialize(config *config.Config) {
	// dbURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True",
	dbURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		config.DB.Host,
		config.DB.Port,
		config.DB.Username,
		config.DB.Name,
		config.DB.Password,
		config.DB.SslMode,
	)

	db, err := gorm.Open(config.DB.Dialect, dbURI)
	if err != nil {
		log.Fatal(err.Error())
	}

	a.DB = model.DBMigrate(db)
	a.Router = mux.NewRouter()
	a.setRouters()
}

// setRouters sets the all required routers
func (a *App) setRouters() {
	// Routing for handling the users
	a.Get("/users", a.handleRequest(handler.GetAllUsers))
	a.Post("/users", a.handleRequest(handler.CreateUser))
	a.Get("/users/{id:[0-9]+}", a.handleRequest(handler.GetUser))
	a.Put("/users/{id:[0-9]+}", a.handleRequest(handler.UpdateUser))
	a.Delete("/users/{id:[0-9]+}", a.handleRequest(handler.DeleteUser))
}

// Get wraps the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Post wraps the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Put wraps the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// Delete wraps the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}

// RequestHandlerFunction request handler
type RequestHandlerFunction func(db *gorm.DB, w http.ResponseWriter, r *http.Request)

func (a *App) handleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(a.DB, w, r)
	}
}
