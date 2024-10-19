package routes

import (
	"net/http"
	"nodenet/internal/controllers"
)

// InitializeRoutes sets up the API routes
func InitializeRoutes(controller *controllers.NodeController) {
	http.HandleFunc("/get", controller.GetValue)
	http.HandleFunc("/set", controller.SetValue)
	http.HandleFunc("/delete", controller.DeleteValue)
	http.HandleFunc("/getall", controller.GetAllValues)
}
