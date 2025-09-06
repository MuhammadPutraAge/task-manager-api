package main

import (
	"log"
	"net/http"

	"github.com/muhammadputraage/task-manager-api/internal/config"
	"github.com/muhammadputraage/task-manager-api/internal/db"
	"github.com/muhammadputraage/task-manager-api/internal/task"
)

func main() {
	config.LoadEnv()
	database := db.Init()

	if err := database.AutoMigrate(&task.Task{}); err != nil {
		log.Fatalf("❌ Failed to migrate database: %v", err)
	}

	taskRepository := task.NewRepository(database)
	taskHandler := task.NewHandler(taskRepository)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /tasks", taskHandler.GetAll)
	mux.HandleFunc("POST /tasks", taskHandler.Create)
	mux.HandleFunc("GET /tasks/{id}", taskHandler.GetById)
	mux.HandleFunc("PUT /tasks/{id}", taskHandler.Update)
	mux.HandleFunc("DELETE /tasks/{id}", taskHandler.Delete)

	log.Println("✅ Server started on http://localhost:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("❌ Server failed: %v", err)
	}
}
