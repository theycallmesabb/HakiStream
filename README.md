#  HakiStream

HakiStream is a high-performance, scalable backend video streaming system built with **Go**. It supports efficient video uploading, cloud object storage, and chunked **HTTP range-based streaming**. Designed with production-level architecture, it integrates robust security practices including JWT authentication, Redis-backed session/token management, and Cloudflare R2 storage.

---

## 🚀 Features

* **Optimized Streaming:** HTTP Range-based requests for smooth, buffering-free playback and instant seeking.
* **Cloud Native Storage:** Direct-to-cloud video uploading utilizing Cloudflare R2 (S3-compatible object storage) to keep the app stateless.
* **Robust Security:** Secure, middleware-protected routes using stateless JWTs paired with a stateful Redis token validation layer (allowing immediate session revocation).
* **Database Isolation:** Decoupled storage architecture—video files live in R2, while structural user and video metadata live in MongoDB.
* **High Performance:** Built on Go's concurrent architecture and the Gin framework, optimized for large file handling and low-latency delivery.

---

## 🧰 Tech Stack

| Layer | Technology |
| :--- | :--- |
| **Language & Framework** | Go (Golang), Gin Web Framework |
| **Databases** | MongoDB (Metadata), Redis (Session & Token Cache) |
| **Storage** | Cloudflare R2 (S3-Compatible Object Storage) |
| **Authentication** | JSON Web Tokens (JWT), `bcrypt` password hashing |
| **API Style** | RESTful Architecture |
| **Streaming Protocol** | HTTP Range Requests |

---

## ⚙️ Core Workflow

1. **Authentication:** User registers or logs in. The server issues a JWT access token and caches the active session identifier in Redis.
2. **Video Upload:** Authenticated users upload video files via multipart forms. The file stream is piped directly to Cloudflare R2, and a structural metadata record containing the R2 target URL is created in MongoDB.
3. **Playback & Streaming:** When a client requests a video, the system validates the active session in Redis and processes incoming **HTTP Range Requests**, serving the file back in dynamic byte-chunks instead of loading the entire asset into server memory.

---

## 📡 API Endpoints

### Auth
* `POST /register` → Create user account (Passwords hashed via `bcrypt`)
* `POST /login`    → Authenticate user and retrieve JWT token

### Videos
* `POST /upload`     → Upload video file to Cloudflare R2 **[Protected]**
* `GET /videos`      → List all available videos and metadata
* `GET /stream/:id`  → Stream video chunk data with partial content support

---

## 📦 Project Structure

The project follows a clean, decoupled layer-based architecture:

```text
HakiStream/
├── config/       # Database connections (Mongo, Redis) & Env variables
├── controllers/  # Request parsers and HTTP handlers
├── middleware/   # JWT verification & Redis session validators
├── models/       # MongoDB schemas and Data Transfer Objects (DTOs)
├── routes/       # API router groups mapping endpoints to controllers
├── services/     # Core business logic (R2 upload clients, streaming chunk calculators)
├── utils/        # Crypto helpers, token generators, error definitions
├── main.go       # Application entry point
└── go.mod        # Dependency management


Built with ⚡ by Sabyasachee

GitHub: @theycallmesabb

Role: Backend Engineer focused on scalable systems, performance tuning, and distributed architecture.
