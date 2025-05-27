# Player Management API
api/
├── main.go           # Server setup, middleware, routing  
├── handlers.go       # HTTP handlers with proper error handling
├── service.go        # Business logic with thread safety
├── types.go          # Data structures, validation, custom errors
├── handlers_test.go  # Comprehensive test suite
└── README.md         # Complete documentation


A RESTful API for managing football players built with Go's standard library.

## 🚀 Features

### ✅ Improvements Made

- **Thread Safety**: All data operations are protected with read-write mutexes
- **JSON API**: Proper JSON request/response handling (no more form data)
- **Input Validation**: Comprehensive validation with meaningful error messages
- **HTTP Status Codes**: Proper status codes for different scenarios
- **Error Handling**: Structured error responses with detailed messages
- **Middleware**: Request logging and CORS support
- **Graceful Shutdown**: Server handles shutdown signals gracefully
- **Health Check**: `/health` endpoint for monitoring
- **Service Layer**: Separation of concerns with proper architecture
- **Resource Validation**: Checks for resource existence before operations
- **Duplicate Prevention**: Prevents duplicate players (same name + jersey number)

## 📋 API Endpoints

### Health Check
```
GET /health
```

### Player Operations
```
GET    /players           # Get all players
GET    /players/{id}      # Get player by ID
POST   /players           # Create new player
PUT    /players/{id}      # Update existing player
DELETE /players/{id}      # Delete player
```

## 🔧 Request/Response Format

### Standard Response Structure
```json
{
  "status": "success|error",
  "message": "Human readable message",
  "data": {},
  "error": "Error details (only on error)"
}
```

### Player Object
```json
{
  "id": "1",
  "name": "Messi",
  "jersey_number": 10,
  "rating": 99
}
```

## 📝 API Usage Examples

### 1. Get All Players
```bash
curl http://localhost:8080/players
```

### 2. Get Player by ID
```bash
curl http://localhost:8080/players/1
```

### 3. Create New Player
```bash
curl -X POST http://localhost:8080/players \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Mbappé",
    "jersey_number": 7,
    "rating": 91
  }'
```

### 4. Update Player
```bash
curl -X PUT http://localhost:8080/players/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Lionel Messi",
    "jersey_number": 10,
    "rating": 95
  }'
```

### 5. Delete Player
```bash
curl -X DELETE http://localhost:8080/players/1
```

## 🛠 Running the Application

### Prerequisites
- Go 1.21+ (uses new HTTP routing features)

### Run
```bash
cd api
go run .
```

### Environment Variables
```bash
export PORT=8080  # Optional, defaults to 8080
```

## 📊 Validation Rules

- **Name**: Required, non-empty string
- **Jersey Number**: 1-99 (inclusive)
- **Rating**: 1-99 (inclusive)
- **Uniqueness**: No two players can have the same name AND jersey number

## 🔒 Error Handling

The API returns appropriate HTTP status codes:

- `200 OK`: Successful GET, PUT, DELETE operations
- `201 Created`: Successful POST operations
- `400 Bad Request`: Invalid input, malformed JSON
- `404 Not Found`: Player not found
- `409 Conflict`: Player already exists (duplicate name + jersey number)
- `500 Internal Server Error`: Unexpected server errors

## 🏗 Architecture

```
main.go           # Server setup, middleware, routing
handlers.go       # HTTP handlers with proper error handling
service.go        # Business logic with thread safety
types.go          # Data structures, validation, custom errors
```

### Key Components

1. **PlayerService**: Thread-safe data operations with RWMutex
2. **PlayerHandler**: HTTP request/response handling
3. **Middleware**: Logging and CORS support
4. **Validation**: Input validation with custom error types
5. **Graceful Shutdown**: Proper server lifecycle management

## 📈 Performance Considerations

- **Read-Write Mutex**: Allows concurrent reads while protecting writes
- **Connection Pooling**: Proper HTTP server timeouts configured
- **Memory Efficient**: Pre-allocated slices and minimal allocations
- **Request Logging**: Performance monitoring with request duration

## 🔄 Backward Compatibility

Legacy function signatures are maintained for backward compatibility:
- `GetPlayers(w, r)`
- `CreatePlayer(w, r)`
- `UpdatePlayer(w, r)`
- `DeletePlayer(w, r)`

## 🧪 Testing

### Manual Testing
Use the provided curl examples above or tools like Postman.

### Sample Test Data
The API starts with 3 sample players:
1. Messi (Jersey: 10, Rating: 99)
2. Ronaldo (Jersey: 7, Rating: 98)
3. Neymar (Jersey: 10, Rating: 95)

## 🚨 Known Limitations

1. **In-Memory Storage**: Data is lost on server restart
2. **No Authentication**: API is open to all requests
3. **No Rate Limiting**: No request throttling implemented
4. **No Pagination**: All players returned in single response

## 🔮 Future Improvements

1. **Database Integration**: PostgreSQL/MySQL support
2. **Authentication**: JWT-based auth system
3. **Caching**: Redis integration for performance
4. **Pagination**: Limit/offset support for large datasets
5. **Filtering**: Query parameters for searching players
6. **Rate Limiting**: Request throttling middleware
7. **Unit Tests**: Comprehensive test coverage
8. **Docker Support**: Containerization
9. **API Documentation**: OpenAPI/Swagger integration
10. **Metrics**: Prometheus metrics collection 