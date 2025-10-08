# Shared Package — Golang (Gin Compatible)

`shared-pkg` is a **common utility library** for Golang microservices that standardizes **API responses**, **error handling**, and **internationalization (i18n)**.  
It’s designed to be **framework-agnostic** but integrates perfectly with **Gin**.


## dependencies

```
gin     -> `go get github.com/gin-gonic/gin`
goil8n  -> `go get github.com/nicksnyder/go-i18n/v2/goi18n`
        -> `go get golang.org/x/text`
Zap Logger -> `go.uber.org/zap`
```

---

## 📂 Folder Structure

```
shared-pkg/
├── go.mod
├── response/
│   ├── meta.go
│   └── response.go
├── errors/
│   ├── error_factory.go
│   └── http_status_map.go
│   └── internal_error_map.go
│   └── static_code.go
└── translator/
    ├── translator.go
    └── locales/
        ├── en.json
        └── id.json
└── utils/
│   ├── request_id.go

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

## Format Logger

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
Message: Resource tidak ditemukan: user
Data: {
  "error": "error test",
  "method": "GET",
  "path": "/error",
  "request_id": "28ec7947-9468-47a1-a6a5-e61e01036d2b",
  "status": 404
}
Error: error test
───────────────────────────────────────────────────────────────
```
