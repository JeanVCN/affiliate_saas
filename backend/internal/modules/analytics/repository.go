package analytics

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	ClickMetrics(ctx context.Context, workspaceID string, groupBy string) ([]ClickMetric, error)
}

type PostgresRepository struct {
	db *pgxpool.Pool
}

func NewPostgresRepository(db *pgxpool.Pool) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (repo *PostgresRepository) ClickMetrics(ctx context.Context, workspaceID string, groupBy string) ([]ClickMetric, error) {
	query := `
		SELECT 'product' AS group_name, p.id, p.name, COUNT(*)::bigint
		FROM click_events ce
		JOIN products p ON p.id = ce.product_id
		WHERE ce.workspace_id = $1
		GROUP BY p.id, p.name
		ORDER BY COUNT(*) DESC, p.name`
	if groupBy == "link" {
		query = `
			SELECT 'link' AS group_name, al.id, COALESCE(al.label, al.destination_url), COUNT(*)::bigint
			FROM click_events ce
			JOIN affiliate_links al ON al.id = ce.affiliate_link_id
			WHERE ce.workspace_id = $1
			GROUP BY al.id, al.label, al.destination_url
			ORDER BY COUNT(*) DESC, COALESCE(al.label, al.destination_url)`
	}

	rows, err := repo.db.Query(ctx, query, workspaceID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []ClickMetric
	for rows.Next() {
		var item ClickMetric
		if err := rows.Scan(&item.Group, &item.GroupID, &item.GroupLabel, &item.Clicks); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, rows.Err()
}
