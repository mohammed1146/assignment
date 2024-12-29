package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/mohammed1146/skelton/internal/domain"
)

type SpacecraftRepository interface {
	ListSpacecrafts(ctx context.Context, name string, class string, status string) ([]domain.Spacecraft, error)
}

type spacecraftRepository struct {
	db *sql.DB
}

func NewSpacecraftRepository(db *sql.DB) SpacecraftRepository {
	return &spacecraftRepository{db: db}
}

// ListSpacecrafts is to list all spacecrafts.
func (r *spacecraftRepository) ListSpacecrafts(ctx context.Context, name string, class string, status string) ([]domain.Spacecraft, error) {
	var (
		query  string
		rows   *sql.Rows
		err    error
		params []interface{}
	)

	query = `SELECT sc.id, sc.name, sc.class, sc.crew, sc.value, sc.status FROM spacecrafts sc`
	conditions := []string{}

	if name != "" {
		conditions = append(conditions, "name LIKE ?")
		params = append(params, "%"+name+"%")
	}

	if class != "" {
		conditions = append(conditions, "class = ?")
		params = append(params, "%"+class+"%")
	}

	if status != "" {
		conditions = append(conditions, "status = ?")
		params = append(params, "%"+status+"%")
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " OR ")
	}

	rows, err = r.db.QueryContext(ctx, query, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var spacecrafts []domain.Spacecraft
	for rows.Next() {
		fmt.Print("recrods")
		var spacecraft domain.Spacecraft
		if err := rows.Scan(&spacecraft.ID, &spacecraft.Name, &spacecraft.Class, &spacecraft.Crew, &spacecraft.Value, &spacecraft.Status); err != nil {
			return nil, err
		}
		spacecrafts = append(spacecrafts, spacecraft)
	}

	return spacecrafts, nil
}
