# Go hexagonal architecture starter
A reusable basic hexagonal architecture in go
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
## Lanch the database
```bash
make db
```
## run
```bash
make run
```
## build
```bash
make build
```
