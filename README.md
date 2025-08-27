
# Logger API

**Logger API** is a high-performance, minimalist **Go** REST API for logging user actions and providing analytics.  
It uses **PostgreSQL** for persistent storage and is fully containerized with **Docker** and **Docker Compose** for easy deployment.

---

## Features

- Log user actions with timestamps and latency metrics.
- Retrieve logs and basic analytics such as count and average latency.
- Fully containerized using Docker for development and production.
- Configurable using environment variables.
- Clean, maintainable Go code using best practices.

---

## Tech Stack

- **Backend:** Go (Echo Framework)
- **Database:** PostgreSQL
- **Containerization:** Docker & Docker Compose
- **Database Driver:** pgx (PostgreSQL driver for Go)

---

## Getting Started

### Prerequisites

- Docker & Docker Compose
- Go 1.25+ (for local development/testing)
- Git

---

### Installation & Setup

1. Clone the repository:

```bash
git clone https://github.com/your-username/logger-api.git
cd logger-api
````

2. Create a `.env` file in the root directory:

```env
POSTGRES_USER=admin
POSTGRES_PASSWORD=123
POSTGRES_DB=loggerdb
```

3. Build and start the services using Docker Compose:

```bash
docker-compose up --build
```

4. The API will be available at:

* Logs endpoint: `http://localhost:8080/api/logs`
* Stats endpoint: `http://localhost:8080/api/stats`

---

## API Endpoints

| Method | Endpoint     | Description                          |
| ------ | ------------ | ------------------------------------ |
| GET    | `/api/logs`  | Retrieve all logs                    |
| GET    | `/api/stats` | Retrieve log count & average latency |

> You can expand the API to add POST, DELETE, or filter queries as needed.

---

## Project Structure

```
logger-api/
├── Dockerfile
├── docker-compose.yml
├── .env.example
├── go.mod
├── go.sum
├── main.go
├── db.go           # Database connection and pool initialization
└── README.md
```


Docker Compose mounts this SQL file, so the table is automatically created when the PostgreSQL container starts.

---

## Environment Variables

| Variable           | Description                  |
| ------------------ | ---------------------------- |
| POSTGRES\_USER     | Database username            |
| POSTGRES\_PASSWORD | Database password            |
| POSTGRES\_DB       | Database name                |
| DATABASE\_URL      | Connection string for Go API |

The API reads the connection string from `DATABASE_URL` or environment variables.

---

## Running Locally

You can also run the API without Docker:

```bash
go mod download
go run main.go
```

Make sure you have a PostgreSQL instance running locally and update `DATABASE_URL` in `.env`.

---

## Contributing

1. Fork the repository.
2. Create a new branch: `git checkout -b feature/my-feature`
3. Make your changes.
4. Commit your changes: `git commit -m "Add new feature"`
5. Push to your branch: `git push origin feature/my-feature`
6. Open a Pull Request.

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

## Acknowledgements

* [Echo Framework](https://echo.labstack.com/) - Go web framework
* [pgx](https://github.com/jackc/pgx) - PostgreSQL driver for Go
* Docker & Docker Compose for containerization


---

If you want, I can **also create a `.env.example` and `.gitignore` template** for this repo so you don’t push sensitive data like passwords.  

Do you want me to do that too?
