package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

type Log struct {
	ID        int       `json:"id"`
	UserID    string    `json:"user_id"`
	Action    string    `json:"action"`
	Latencyms int       `json:"latency_ms"`
	CreatedAt time.Time `json:"created_at"`
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system env")
	}

	dbURL := os.Getenv("DATABASE_URL")

	InitDB(dbURL)

	e := echo.New()

	e.GET("/api/logs", func(c echo.Context) error {
		r, err := DB.Query(c.Request().Context(), "SELECT id, user_id, action, latency_ms, created_at FROM logs;")

		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		defer r.Close()

		var logs []Log
		for r.Next() {
			var log Log

			if err := r.Scan(&log.ID, &log.UserID, &log.Action, &log.Latencyms, &log.CreatedAt); err != nil {
				return c.String(http.StatusInternalServerError, err.Error())
			}
			logs = append(logs, log)
		}

		return c.JSON(http.StatusOK, logs)

	})

	e.POST("/api/log", func(c echo.Context) error {
		var input struct {
			UserID    string `json:"user_id"`
			Action    string `json:"action"`
			Latencyms int    `json:"latency_ms"`
		}

		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
		}

		_, err := DB.Exec(
			c.Request().Context(),
			"INSERT INTO logs (user_id, action, latency_ms) VALUES ($1, $2, $3)",
			input.UserID, input.Action, input.Latencyms,
		)

		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusCreated, map[string]string{"message": "log inserted"})
	})

	e.GET("/api/stats", func(c echo.Context) error {

		var total int
		var avg float64

		err := DB.QueryRow(c.Request().Context(), "SELECT COUNT(*), AVG(latency_ms) FROM logs;").Scan(&total, &avg)

		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		res := map[string]interface{}{
			"total_logs":  total,
			"avg_latency": avg,
		}

		return c.JSON(http.StatusOK, res)

	})

	e.Logger.Fatal(e.Start(":8080"))
}
