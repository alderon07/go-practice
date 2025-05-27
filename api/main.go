package main

import (
  "context"
  "encoding/json"
  "log"
  "net/http"
  "os"
  "os/signal"
  "syscall"
  "time"
)

// LoggingMiddleware logs HTTP requests
func LoggingMiddleware(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    start := time.Now()
    
    // Create a response recorder to capture status code
    recorder := &responseRecorder{ResponseWriter: w, statusCode: http.StatusOK}
    
    next.ServeHTTP(recorder, r)
    
    duration := time.Since(start)
    log.Printf("%s %s %d %v", r.Method, r.URL.Path, recorder.statusCode, duration)
  })
}

// CORS middleware to handle Cross-Origin Resource Sharing
func CORSMiddleware(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
    
    // Handle preflight requests
    if r.Method == "OPTIONS" {
      w.WriteHeader(http.StatusOK)
      return
    }
    
    next.ServeHTTP(w, r)
  })
}

// responseRecorder is a custom ResponseWriter to capture status codes
type responseRecorder struct {
  http.ResponseWriter
  statusCode int
}

func (r *responseRecorder) WriteHeader(statusCode int) {
  r.statusCode = statusCode
  r.ResponseWriter.WriteHeader(statusCode)
}

func main() {
  // Initialize service and handler
  playerService := NewPlayerService()
  playerHandler := NewPlayerHandler(playerService)
  
  // Create router
  router := http.NewServeMux()
  
  // Register routes with the improved handlers
  router.HandleFunc("GET /players", playerHandler.GetPlayers)
  router.HandleFunc("GET /players/{id}", playerHandler.GetPlayer)
  router.HandleFunc("POST /players", playerHandler.CreatePlayer)
  router.HandleFunc("PUT /players/{id}", playerHandler.UpdatePlayer)
  router.HandleFunc("DELETE /players/{id}", playerHandler.DeletePlayer)
  
  // Add health check endpoint
  router.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
    response := Response{
      Status:  "success",
      Message: "Service is healthy",
      Data:    map[string]string{"version": "1.0.0", "service": "player-api"},
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
  })
  
  // Apply middleware
  handler := LoggingMiddleware(CORSMiddleware(router))
  
  // Configure server
  port := os.Getenv("PORT")
  if port == "" {
    port = "8080"
  }
  
  server := &http.Server{
    Addr:         ":" + port,
    Handler:      handler,
    ReadTimeout:  15 * time.Second,
    WriteTimeout: 15 * time.Second,
    IdleTimeout:  60 * time.Second,
  }
  
  // Start server in a goroutine
  go func() {
    log.Printf("🚀 Server starting on port %s", port)
    log.Printf("📋 Available endpoints:")
    log.Printf("   GET    /health")
    log.Printf("   GET    /players")
    log.Printf("   GET    /players/{id}")
    log.Printf("   POST   /players")
    log.Printf("   PUT    /players/{id}")
    log.Printf("   DELETE /players/{id}")
    
    if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
      log.Fatalf("Failed to start server: %v", err)
    }
  }()
  
  // Wait for interrupt signal to gracefully shutdown the server
  quit := make(chan os.Signal, 1)
  signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
  <-quit
  
  log.Println("🛑 Shutting down server...")
  
  // Give outstanding requests 30 seconds to complete
  ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
  defer cancel()
  
  if err := server.Shutdown(ctx); err != nil {
    log.Fatalf("Server forced to shutdown: %v", err)
  }
  
  log.Println("✅ Server stopped gracefully")
}

