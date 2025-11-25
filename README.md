# Go hexagonal architecture starter
```bash
.
├── cmd
│   └── main.go
├── go.mod
├── go.sum
├── internal
│   ├── adaptaters
│   │   ├── primary
│   │   └── secondary
│   │       └── gorm.go
│   ├── core
│   │   ├── domain
│   │   │   └── user.go
│   │   └── ports
│   │       └── db.go
│   └── infrastructure
│       ├── docker-compose.yaml
│       ├── migrations
│       │   └── user.go
│       └── migrations.go
├── main
├── Makefile
└── README.md
```
