package database

import (
	"encoding/json"
	"net"
	"os"
	"runtime"

	"github.com/jmoiron/sqlx"
)

type LogRepo struct {
	db *sqlx.DB
}

func NewLogRepo(db *sqlx.DB) *LogRepo {
	return &LogRepo{db: db}
}

// PushLog : La fonction qui fait tout le travail sale pour toi
func (r *LogRepo) PushLog(userID int, action, module string, targetID int, severity string, payload interface{}) error {
	// 1. Collecte automatique des infos de cohérence
	hostname, _ := os.Hostname()
	osInfo := runtime.GOOS + " (" + runtime.GOARCH + ")"

	// Récupération IP locale simple
	ip := "127.0.0.1"
	addrs, _ := net.InterfaceAddrs()
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip = ipnet.IP.String()
				break
			}
		}
	}

	// 2. Conversion du payload en JSON string
	payloadJSON, _ := json.Marshal(payload)

	// 3. Insertion SQLX (plus propre avec les Named Queries)
	query := `INSERT INTO logs (
		user_id, action, module, target_id, severity, 
		machine_name, os_info, ip_address, payload
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`

	_, err := r.db.Exec(query,
		userID, action, module, targetID, severity,
		hostname, osInfo, ip, string(payloadJSON),
	)

	return err
}
