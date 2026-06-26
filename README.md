# SkillCycle Backend

Built with Go (Gin) + PostgreSQL

## Setup

1. Install Go and PostgreSQL
2. Create database:


CREATE DATABASE skillcycle_db;

CREATE USER skillcycle_user WITH PASSWORD 'skillcycle123';

GRANT ALL ON SCHEMA public TO skillcycle_user;
3. Create a `.env` file:
DB_HOST=localhost

DB_PORT=5432

DB_USER=skillcycle_user

DB_PASSWORD=skillcycle123

DB_NAME=skillcycle_db

JWT_SECRET=skillcycle_super_secret_key_2024

PORT=8080
4. Run:
go run main.go

## API Routes

### Auth
- POST /api/auth/register
- POST /api/auth/login

### Listings
- GET  /api/listings (public)
- GET  /api/listings/:id (public)
- POST /api/listings (protected)
- PUT  /api/listings/:id (protected)
- DELETE /api/listings/:id (protected)