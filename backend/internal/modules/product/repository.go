package product

import (
	"context"

	"github.com/JeanVCN/affiliate_saas/backend/internal/modules/common"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	ListProducts(ctx context.Context, workspaceID string) ([]Product, error)
	CreateProduct(ctx context.Context, workspaceID string, input CreateProductInput) (Product, error)
	GetProduct(ctx context.Context, workspaceID string, productID string) (Product, error)
	CreateOffer(ctx context.Context, workspaceID string, productID string, input CreateOfferInput) (Offer, error)
}

type PostgresRepository struct {
	db *pgxpool.Pool
}

func NewPostgresRepository(db *pgxpool.Pool) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (repo *PostgresRepository) ListProducts(ctx context.Context, workspaceID string) ([]Product, error) {
	rows, err := repo.db.Query(ctx, `
		SELECT id, workspace_id, name, COALESCE(category, ''), COALESCE(description, ''),
		       status, created_at, updated_at
		FROM products
		WHERE workspace_id = $1 AND archived_at IS NULL
		ORDER BY created_at DESC`, workspaceID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []Product
	for rows.Next() {
		var item Product
		if err := rows.Scan(&item.ID, &item.WorkspaceID, &item.Name, &item.Category, &item.Description, &item.Status, &item.CreatedAt, &item.UpdatedAt); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, rows.Err()
}

func (repo *PostgresRepository) CreateProduct(ctx context.Context, workspaceID string, input CreateProductInput) (Product, error) {
	item := Product{ID: common.NewID("prd")}
	err := repo.db.QueryRow(ctx, `
		INSERT INTO products (id, workspace_id, name, category, description)
		VALUES ($1, $2, $3, NULLIF($4, ''), NULLIF($5, ''))
		RETURNING id, workspace_id, name, COALESCE(category, ''), COALESCE(description, ''),
		          status, created_at, updated_at`,
		item.ID, workspaceID, input.Name, input.Category, input.Description,
	).Scan(&item.ID, &item.WorkspaceID, &item.Name, &item.Category, &item.Description, &item.Status, &item.CreatedAt, &item.UpdatedAt)
	return item, common.NormalizePostgresErr(err)
}

func (repo *PostgresRepository) GetProduct(ctx context.Context, workspaceID string, productID string) (Product, error) {
	var item Product
	err := repo.db.QueryRow(ctx, `
		SELECT id, workspace_id, name, COALESCE(category, ''), COALESCE(description, ''),
		       status, created_at, updated_at
		FROM products
		WHERE workspace_id = $1 AND id = $2 AND archived_at IS NULL`,
		workspaceID, productID,
	).Scan(&item.ID, &item.WorkspaceID, &item.Name, &item.Category, &item.Description, &item.Status, &item.CreatedAt, &item.UpdatedAt)
	return item, common.NormalizePostgresErr(err)
}

func (repo *PostgresRepository) CreateOffer(ctx context.Context, workspaceID string, productID string, input CreateOfferInput) (Offer, error) {
	item := Offer{ID: common.NewID("off")}
	err := repo.db.QueryRow(ctx, `
		INSERT INTO offers (id, workspace_id, product_id, workspace_program_id, title, price_cents, currency)
		SELECT $1, $2, p.id, wp.id, NULLIF($5, ''), $6, NULLIF($7, '')
		FROM products p
		JOIN workspace_programs wp ON wp.id = $4 AND wp.workspace_id = $2 AND wp.archived_at IS NULL
		WHERE p.id = $3 AND p.workspace_id = $2 AND p.archived_at IS NULL
		RETURNING id, workspace_id, product_id, workspace_program_id, COALESCE(title, ''),
		          price_cents, COALESCE(currency, ''), status, created_at, updated_at`,
		item.ID, workspaceID, productID, input.WorkspaceProgramID, input.Title, input.PriceCents, input.Currency,
	).Scan(&item.ID, &item.WorkspaceID, &item.ProductID, &item.WorkspaceProgramID, &item.Title, &item.PriceCents, &item.Currency, &item.Status, &item.CreatedAt, &item.UpdatedAt)
	return item, common.NormalizePostgresErr(err)
}
