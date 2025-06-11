users-service/
├── cmd/
│   └── server/
│       └── main.go               ← инициализация DB и запуск gRPC
├── internal/
│   ├── database/
│   │   └── db.go                 ← InitDB() (GORM)
│   ├── user/
│   │   ├── model.go              ← ORM-модель User
│   │   ├── repository.go         ← UserRepository
│   │   └── service.go            ← UserService (бизнес-логика)
│   └── transport/
│       └── grpc/
│           ├── handler.go        ← реализация всех gRPC-методов
│           └── server.go         ← настройка и запуск grpc.Server
├── go.mod
└── go.sum

<!-- для запуска -->
```zsh
    docker compose up --build
```