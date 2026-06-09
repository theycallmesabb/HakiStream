HakiStream

HakiStream is a scalable backend video streaming system built with Go that supports efficient video upload, storage, and HTTP range-based streaming. It is designed with production-level architecture using modern backend practices like JWT authentication, Redis token management, and Cloudflare R2 object storage.

🚀 Features
📤 Video upload with Cloudflare R2 object storage
🎥 HTTP Range-based streaming for smooth playback and seeking
📃 Video listing API
👤 User authentication system (Register & Login)
🔐 JWT-based authentication with Redis token session handling
🧠 Secure middleware-protected routes
🗄️ MongoDB integration for users and video metadata
⚡ Optimized for large video file handling and scalable delivery

🧰 Tech Stack
Backend: Go (Golang), Gin
Database: MongoDB
Cache / Session Store: Redis
Auth: JWT, bcrypt
Storage: Cloudflare R2 (S3-compatible object storage)
API Style: REST API
Streaming: HTTP Range Requests
🏗️ System Architecture
Client
  ↓
Gin REST API
  ↓
JWT Auth Middleware + Redis Validation
  ↓
Business Logic Layer
  ↓
MongoDB (Metadata) + Cloudflare R2 (Video Storage)
⚙️ How It Works
User registers and logs in
Server generates JWT token and stores session in Redis
Authenticated users upload videos to Cloudflare R2
Metadata is stored in MongoDB
Videos are streamed using HTTP Range requests for efficient playback
Redis validates active sessions for each protected request
📡 API Endpoints
Auth
POST /register   → Create user
POST /login      → Get JWT token
Videos
POST /upload     → Upload video (Protected)
GET /videos      → List all videos
GET /stream/:id  → Stream video with range support
🔐 Security
Passwords hashed using bcrypt
JWT authentication for stateless security
Redis used for session/token validation
Middleware-based route protection
☁️ Storage Design
Videos stored in Cloudflare R2 (S3-compatible)
Only metadata stored in MongoDB
Optimized for scalable and low-cost media delivery
📦 Project Structure
HakiStream/
├── config/
├── controllers/
├── middleware/
├── models/
├── routes/
├── services/
├── utils/
├── main.go
└── go.mod
🔥 Future Improvements
🎬 Video transcoding pipeline (FFmpeg)
📊 Analytics dashboard (views, watch time)
🔁 Refresh token system
🌍 CDN integration for faster streaming
👥 Role-based access control (Admin/User)
👨‍💻 Author

Built by Sabyasachee 🚀
Backend Engineer focused on scalable systems and distributed architecture.
