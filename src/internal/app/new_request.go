package app

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
	"main/src/internal/pkg"
	"net/http"
	"strings"
)

// SaveRequestUseCase use case for saving request to database for statistics
type SaveRequestUseCase struct {
	db *pgx.Conn
}

// NewSaveRequestUseCase constructor for SaveRequestUseCase
func NewSaveRequestUseCase(db *pgx.Conn) *SaveRequestUseCase {
	return &SaveRequestUseCase{db: db}
}

func (u *SaveRequestUseCase) SaveRequest(req *http.Request) error {
	clientIP := req.Header.Get("X-Forwarded-For")
	ipInfo, err := pkg.GetIPInfo(clientIP)
	if err != nil {
		return errors.Wrap(err, "get ip info")
	}

	sqlQuery := `INSERT INTO request_info (ip, os, country, region, city, latitude, longitude, user_agent, headers, endpoint)
					VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);`
	_, err = u.db.Exec(context.Background(), sqlQuery,
		clientIP,
		getOsTypeFromUserAgent(req.Header.Get("User-Agent")),
		ipInfo.Country,
		ipInfo.Region,
		ipInfo.City,
		ipInfo.Latitude,
		ipInfo.Longitude,
		req.Header.Get("User-Agent"),
		getAllHeaders(req),
		req.URL.Path)
	if err != nil {
		return errors.Wrap(err, "insert into request_info")
	}

	return nil
}

// getOsTypeFromUserAgent returns approximate os type determined based of UserAgent
func getOsTypeFromUserAgent(userAgent string) string {
	if userAgent == "" {
		return ""
	}
	tmpString := strings.Split(userAgent, "(")
	systemInfo := ""
	if len(tmpString) == 3 {
		tmpString2 := strings.Split(tmpString[1], ")")
		if len(tmpString2) == 2 {
			systemInfo = tmpString2[0]
		}
	}
	return systemInfo
}

func getAllHeaders(req *http.Request) string {
	headers := ""
	for key, values := range req.Header {
		headers += fmt.Sprintf("%s: %s\n", key, values)
	}
	return headers
}
