package store

import (
	"database/sql"
	"fmt"

	"github.com/comtom/portsService/logger"
	"github.com/comtom/portsService/ports"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

type Database struct {
	client *sql.DB
	logger *logger.Logger
}

func LatLngToString(coordinates []float64) string {
	if len(coordinates) != 2 {
		return "(0,0)"
	}
	return fmt.Sprintf("(%f, %f)", coordinates[0], coordinates[1])
}

func (d Database) Put(p ports.Port) error {
	stmt := `INSERT INTO ports(code_id, name, city, province, country, coordinates, timezone, unlocs, code) VALUES(
	$1, $2, $3,	$4,	$5,	$6,	$7,	$8, $9) ON CONFLICT (code_id) DO UPDATE
	SET name=EXCLUDED.name, city=EXCLUDED.city, province=EXCLUDED.province, country=EXCLUDED.country, coordinates=EXCLUDED.coordinates, timezone=EXCLUDED.timezone, unlocs=EXCLUDED.unlocs, code=EXCLUDED.code`

	_, err := d.client.Exec(stmt, p.Unlocs[0], p.Name, p.City, p.Province, p.Country, LatLngToString(p.Coordinates), p.Timezone, pq.Array(p.Unlocs), p.Code)
	if err != nil {
		return err
	}

	return nil
}

func (d Database) Get(unloc string) (ports.Port, error) {
	panic("not implemented")
}

func (d Database) Shutdown() {
	d.client.Close()
	d.logger.Info("disconnected from database")
}

func NewDBStore(host string, port string, user string, password string, database string, logger *logger.Logger) (Database, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, database)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		logger.Error("error connecting to the database")
		return Database{}, err
	}

	logger.Info("connected to database")
	return Database{
		client: db,
		logger: logger,
	}, nil
}
