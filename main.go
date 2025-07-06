package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	mux := http.NewServeMux()
	
	staticHandler := http.StripPrefix("/static/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Security: Block hidden files
        if strings.HasPrefix(r.URL.Path, ".") {
            http.NotFound(w, r)
            return
        }
        w.Header().Set("Cache-Control", "public, max-age=31536000")
        http.FileServer(http.Dir("./static")).ServeHTTP(w, r)
    }))
	
    mux.Handle("/static/", staticHandler)
	

	// Основные обработчики
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	mux.HandleFunc("/project/", projectHandler)

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Graceful shutdown
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)
		<-sigint

		log.Println("Shutting down server...")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		server.Shutdown(ctx)
	}()

	log.Printf("Server started on :%s", port)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}
}

func projectHandler(w http.ResponseWriter, r *http.Request) {
	projectName := r.URL.Path[len("/project/"):]

	// Здесь можно загружать данные проекта из БД
	// и динамически генерировать страницу

	html := fmt.Sprintf(`
	<!DOCTYPE html>
	<html>
	<head>
		<title>Проект %s</title>
	</head>
	<body>
		<h1>Страница проекта: %s</h1>
		<p>Детальное описание проекта...</p>
		<a href="/">На главную</a>
	</body>
	</html>
	`, projectName, projectName)

	fmt.Fprint(w, html)
}
