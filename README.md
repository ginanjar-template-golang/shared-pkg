# Shared Package — Golang (Gin Compatible)

`shared-pkg` is a **common utility library** for Golang microservices that standardizes **API responses**, **error handling**, and **internationalization (i18n)**.  
integrates perfectly with **Gin**.


## dependencies

```
gin         -> `go get github.com/gin-gonic/gin`
goil8n      -> `go get github.com/nicksnyder/go-i18n/v2/goi18n` & `go get golang.org/x/text`
Zap Logger  -> `go.uber.org/zap`
```

---

## 📂 Folder Structure

```
shared-pkg/
├── go.mod
├── constants/
│   ├── http_code.go
│   └── internal_code.go
├── db/
│   ├── translator.go
│    └── gorm/
│        └── repository
│            └── base_repository.go
├── errors/
│   ├── errors.go
├── logger/
│   └── logger.go
├── middleware/
│   ├── auth_jwt.go
│   ├── cors.go
│   ├── recovery.go
│   └── request.go
├── response/
│   └── response.go
└── translator/
│   ├── translator.go
│   └── messages/
│       ├── errors
│       │   ├── en.json
│       │   └── id.json
│       └── success
│           ├── en.json
│           └── id.json
└── utils/
│   ├── request_id.go
│   └── sensitive_field.go

```

## Format Standar Response

```

✅ Response Sukses
{
  "meta_data": {
    "status": true,
    "request_id": "uuid-v4",
    "code": 200,
    "message": "Success"
  },
  "data": {
    "results": {
      "user": "Genjer Dotkom",
      "role": "Admin"
    }
  }
}

❌ Response Error
{
  "meta_data": {
    "status": false,
    "request_id": "uuid-v4",
    "code": 404,
    "message": "Resource tidak ditemukan"
  },
  "errors": "resource_not_found"
}
```

## Logger

### Initialize Logger

```
logger.Init(logger.Config{
		LogglyUrl:   "https://logs-01.loggly.com/inputs/%s/tag/%s",
		LogglyToken: "", //your-loggly-token
		LogglyTag:   "service-shared-pkg",
		Environment: "dev", // dev (TRACE,DEBUG,INFO,WARN,ERROR) | staging (TRACE,INFO,WARN,ERROR) | prod (WARN,ERROR)
		AllLogLevel: false,
})
```

### Format Logger

```
───────────────────────────────────────────────────────────────
[INFO] 2025-10-08T16:26:15+07:00
RequestID: d3c5c1f6-c1ba-47a5-bc62-3ebf98405220
Message: success get users
Data: {
  "limit": 10,
  "method": "GET",
  "page": 1,
  "path": "/success-pagination",
  "request_id": "d3c5c1f6-c1ba-47a5-bc62-3ebf98405220",
  "status": 200,
  "total": 2
}
───────────────────────────────────────────────────────────────

[ERROR] 2025-10-08T16:25:56+07:00
RequestID: 28ec7947-9468-47a1-a6a5-e61e01036d2b
Message: Something Wrong
Error: error test
───────────────────────────────────────────────────────────────
```

```
 protoc \      
  --proto_path=./proto \
  --go_out=. --go_opt=module=github.com/ginanjar-template-golang/shared-pkg \
  --go-grpc_out=. --go-grpc_opt=module=github.com/ginanjar-template-golang/shared-pkg \
  proto/*.proto
```
