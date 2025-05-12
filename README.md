# go-svc-template

[![Go](https://img.shields.io/badge/go-1.24%2B-blue)](https://golang.org) [![Bazel](https://img.shields.io/badge/bazel-8.x%2B-green)](https://bazel.build) [![License: MIT](https://img.shields.io/badge/license-MIT-yellow.svg)](LICENSE)

Minimal, idiomatic Go microservice boilerplate—standalone or monorepo-ready—featuring:

- **HTTP** & **gRPC** servers (Uber FX, Cobra, Viper, Zap)
- **Config** via YAML + flags + env (Viper)
- **Logging** with Zap (dev config + cleanup)
- **Layered DDD** structure: `repo` → `service` → `transport` → `server`
- **Lifecycle** management (FX modules)
- **Bazel** build (Bzlmod: rules_go + Gazelle + go_deps)

---

## 🚀 Features

- HTTP endpoints via `net/http`
- gRPC stubs via `protoc` & Bazel
- Configurable via `configs/config.yml`, `--config` flag or `CONFIG` env
- Zap logger with `Sync()` on shutdown
- Dependency injection & graceful start/stop with FX
- Easy to extend or embed in a monorepo

## 🛠 Prerequisites

- Go ≥ 1.24
- Bazel ≥ 8.x
- `protoc`, `protoc-gen-go`, `protoc-gen-go-grpc` in PATH

---

## 🎬 Quick Start

```bash
git clone https://github.com/froppa/go-svc-template.git
cd go-svc-template

# Update Go deps and BUILD files
./scripts/update-deps.sh

# Build & run HTTP service
./scripts/build.sh serve-http
./scripts/run.sh   serve-http -- --config=configs/config.yml
./scripts/run.sh --direct serve-http -- --config=configs/config.yml

# Or debug locally with Go
go run ./cmd/serve-http --config=configs/config.yml
```
---

## 🧩 Components & Roles

- **cmd/serve-http**
  Cobra CLI entrypoint (`main.go`): parses flags, initializes FX app, starts HTTP.

- **configs/config.yml**
  YAML config (e.g. ports). Loaded by Viper in `pkg/config`.

- **internal/repo**
  Repository layer (data access). Swap in-memory stub for real DB.

- **internal/service**
  Business logic. Implements use-cases by calling repo.

- **internal/transport**
  Transport layer.
  - `http.go`: builds an `http.Handler`.
  - `grpc.go`: builds gRPC services.

- **pkg/config**
  Viper + Cobra flag wiring. Provides `*Config` to FX.

- **pkg/logger**
  Zap logger setup + FX cleanup hook.

- **pkg/server/httpfx**
  FX module: consumes `http.Handler`, port, logger; registers OnStart/OnStop hooks.

- **pkg/server/grpcfx**
  FX module: consumes gRPC services, port, logger; registers OnStart/OnStop hooks.

- **proto/**
  `.proto` definitions + generated Go stubs.

- **scripts/**
  Helpers: `build.sh`, `run.sh`, `update-deps.sh`.

---
