package pollDataMysql

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/lukemakhanu/learning/internal/domains/pollData"
)

var _ pollData.PollDataRepository = (*MysqlRepository)(nil)

type MysqlRepository struct {
	db *sql.DB
}

// Create a new mysql repository
func New(connectionString string) (*MysqlRepository, error) {
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(5)
	// Maximum Open Connections
	db.SetMaxOpenConns(10)
	// Idle Connection Timeout
	db.SetConnMaxIdleTime(5 * time.Second)
	// Connection Lifetime
	db.SetConnMaxLifetime(15 * time.Second)

	return &MysqlRepository{
		db: db,
	}, nil
}

// Save : saves live score record into db
func (mr *MysqlRepository) Save(ctx context.Context, t pollData.Sports) (int, error) {
	var d int
	rs, err := mr.db.Exec("INSERT sports SET sport_id=?,sport_name=?,s_binomen=?,match_count=?, \n"+
		"sport_type_id=?,created=now(),modified=now() ON DUPLICATE KEY UPDATE modified=now()",
		t.SportID, t.SportName, t.SBinomen, t.MatchCount, t.SportTypeID)

	if err != nil {
		return d, fmt.Errorf("Unable to save sports : %v", err)
	}

	lastInsertedID, err := rs.LastInsertId()
	if err != nil {
		return d, fmt.Errorf("Unable to retrieve last id [primary key] : %v", err)
	}

	return int(lastInsertedID), nil
}

// LiveScores : query data used when querying for live scores.
func (mr *MysqlRepository) Get(ctx context.Context) ([]pollData.Sports, error) {

	var gc []pollData.Sports

	statement := fmt.Sprintf("SELECT sport_id,sport_name,s_binomen,match_count, \n" +
		"sport_type_id from sports ")
	raws, err := mr.db.Query(statement)
	if err != nil {
		return gc, err
	}
	for raws.Next() {
		var g pollData.Sports
		err := raws.Scan(&g.SportID, &g.SportName, &g.SBinomen, &g.MatchCount,
			&g.SportTypeID)
		if err != nil {
			return gc, err
		}
		gc = append(gc, g)
	}
	if raws.Err(); err != nil {
		return gc, err
	}
	raws.Close()

	return gc, nil
}
