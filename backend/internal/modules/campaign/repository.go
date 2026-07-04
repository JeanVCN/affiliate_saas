package campaign

import (
	"context"
	"database/sql"

	"github.com/JeanVCN/affiliate_saas/backend/internal/modules/common"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	ListCampaigns(ctx context.Context, workspaceID string) ([]Campaign, error)
	CreateCampaign(ctx context.Context, workspaceID string, input CreateCampaignInput) (Campaign, error)
	GetCampaign(ctx context.Context, workspaceID string, campaignID string) (Campaign, error)
	UpdateCampaign(ctx context.Context, workspaceID string, campaignID string, input UpdateCampaignInput) (Campaign, error)
	CreateChannelPackage(ctx context.Context, workspaceID string, campaignID string, input CreateChannelPackageInput) (ChannelPackage, error)
	ListPublishingTasks(ctx context.Context, workspaceID string, campaignID string) ([]PublishingTask, error)
	CreatePublishingTask(ctx context.Context, workspaceID string, campaignID string, input CreatePublishingTaskInput) (PublishingTask, error)
	UpdatePublishingTask(ctx context.Context, workspaceID string, campaignID string, taskID string, input UpdatePublishingTaskInput) (PublishingTask, error)
}

type PostgresRepository struct {
	db *pgxpool.Pool
}

func NewPostgresRepository(db *pgxpool.Pool) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (repo *PostgresRepository) ListCampaigns(ctx context.Context, workspaceID string) ([]Campaign, error) {
	rows, err := repo.db.Query(ctx, `
		SELECT id, workspace_id, COALESCE(product_id, ''), name, status, created_at, updated_at
		FROM campaigns
		WHERE workspace_id = $1 AND archived_at IS NULL
		ORDER BY created_at DESC`, workspaceID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []Campaign
	for rows.Next() {
		var item Campaign
		if err := rows.Scan(&item.ID, &item.WorkspaceID, &item.ProductID, &item.Name, &item.Status, &item.CreatedAt, &item.UpdatedAt); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, rows.Err()
}

func (repo *PostgresRepository) CreateCampaign(ctx context.Context, workspaceID string, input CreateCampaignInput) (Campaign, error) {
	item := Campaign{ID: common.NewID("cmp")}
	err := repo.db.QueryRow(ctx, `
		INSERT INTO campaigns (id, workspace_id, product_id, name)
		SELECT $1, $2, NULLIF($3, ''), $4
		WHERE $3 = '' OR EXISTS (
			SELECT 1 FROM products p
			WHERE p.id = $3 AND p.workspace_id = $2 AND p.archived_at IS NULL
		)
		RETURNING id, workspace_id, COALESCE(product_id, ''), name, status, created_at, updated_at`,
		item.ID, workspaceID, input.ProductID, input.Name,
	).Scan(&item.ID, &item.WorkspaceID, &item.ProductID, &item.Name, &item.Status, &item.CreatedAt, &item.UpdatedAt)
	return item, common.NormalizePostgresErr(err)
}

func (repo *PostgresRepository) GetCampaign(ctx context.Context, workspaceID string, campaignID string) (Campaign, error) {
	item, err := repo.getCampaign(ctx, workspaceID, campaignID)
	if err != nil {
		return Campaign{}, err
	}
	packages, err := repo.listChannelPackages(ctx, workspaceID, campaignID)
	if err != nil {
		return Campaign{}, err
	}
	item.Packages = packages
	return item, nil
}

func (repo *PostgresRepository) UpdateCampaign(ctx context.Context, workspaceID string, campaignID string, input UpdateCampaignInput) (Campaign, error) {
	item := Campaign{}
	err := repo.db.QueryRow(ctx, `
		UPDATE campaigns
		SET name = COALESCE($3, name),
		    status = COALESCE($4, status),
		    updated_at = now(),
		    archived_at = CASE WHEN $4 = 'archived' THEN now() ELSE archived_at END
		WHERE workspace_id = $1 AND id = $2 AND archived_at IS NULL
		RETURNING id, workspace_id, COALESCE(product_id, ''), name, status, created_at, updated_at`,
		workspaceID, campaignID, input.Name, input.Status,
	).Scan(&item.ID, &item.WorkspaceID, &item.ProductID, &item.Name, &item.Status, &item.CreatedAt, &item.UpdatedAt)
	return item, common.NormalizePostgresErr(err)
}

func (repo *PostgresRepository) CreateChannelPackage(ctx context.Context, workspaceID string, campaignID string, input CreateChannelPackageInput) (ChannelPackage, error) {
	item := ChannelPackage{ID: common.NewID("pkg")}
	err := repo.db.QueryRow(ctx, `
		INSERT INTO channel_packages (id, workspace_id, campaign_id, channel, title, body)
		SELECT $1, $2, c.id, $4, NULLIF($5, ''), NULLIF($6, '')
		FROM campaigns c
		WHERE c.id = $3 AND c.workspace_id = $2 AND c.archived_at IS NULL
		RETURNING id, workspace_id, campaign_id, channel, COALESCE(title, ''),
		          COALESCE(body, ''), status, created_at, updated_at`,
		item.ID, workspaceID, campaignID, input.Channel, input.Title, input.Body,
	).Scan(&item.ID, &item.WorkspaceID, &item.CampaignID, &item.Channel, &item.Title, &item.Body, &item.Status, &item.CreatedAt, &item.UpdatedAt)
	return item, common.NormalizePostgresErr(err)
}

func (repo *PostgresRepository) ListPublishingTasks(ctx context.Context, workspaceID string, campaignID string) ([]PublishingTask, error) {
	rows, err := repo.db.Query(ctx, `
		SELECT pt.id, pt.workspace_id, pt.campaign_id, COALESCE(pt.channel_package_id, ''),
		       pt.channel, pt.title, COALESCE(pt.notes, ''), pt.scheduled_for, pt.status,
		       pt.created_at, pt.updated_at, pt.completed_at
		FROM publishing_tasks pt
		JOIN campaigns c ON c.id = pt.campaign_id AND c.workspace_id = pt.workspace_id
		WHERE pt.workspace_id = $1 AND pt.campaign_id = $2 AND c.archived_at IS NULL
		ORDER BY pt.created_at DESC`, workspaceID, campaignID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []PublishingTask
	for rows.Next() {
		var item PublishingTask
		var scheduledFor sql.NullTime
		var completedAt sql.NullTime
		if err := rows.Scan(&item.ID, &item.WorkspaceID, &item.CampaignID, &item.ChannelPackageID, &item.Channel,
			&item.Title, &item.Notes, &scheduledFor, &item.Status, &item.CreatedAt, &item.UpdatedAt, &completedAt); err != nil {
			return nil, err
		}
		applyTaskTimes(&item, scheduledFor, completedAt)
		items = append(items, item)
	}
	return items, rows.Err()
}

func (repo *PostgresRepository) CreatePublishingTask(ctx context.Context, workspaceID string, campaignID string, input CreatePublishingTaskInput) (PublishingTask, error) {
	item := PublishingTask{ID: common.NewID("tsk")}
	var scheduledFor sql.NullTime
	var completedAt sql.NullTime
	err := repo.db.QueryRow(ctx, `
		INSERT INTO publishing_tasks (id, workspace_id, campaign_id, channel_package_id, channel, title, notes, scheduled_for, status)
		SELECT $1, $2, c.id, NULLIF($4, ''), $5, $6, NULLIF($7, ''), $8,
		       CASE WHEN $8::timestamptz IS NULL THEN 'todo' ELSE 'scheduled' END
		FROM campaigns c
		WHERE c.id = $3 AND c.workspace_id = $2 AND c.archived_at IS NULL
		  AND (
		    NULLIF($4, '') IS NULL
		    OR EXISTS (
		      SELECT 1 FROM channel_packages cp
		      WHERE cp.id = $4 AND cp.workspace_id = $2 AND cp.campaign_id = c.id
		    )
		  )
		RETURNING id, workspace_id, campaign_id, COALESCE(channel_package_id, ''),
		          channel, title, COALESCE(notes, ''), scheduled_for, status,
		          created_at, updated_at, completed_at`,
		item.ID, workspaceID, campaignID, input.ChannelPackageID, input.Channel, input.Title, input.Notes, input.ScheduledFor,
	).Scan(&item.ID, &item.WorkspaceID, &item.CampaignID, &item.ChannelPackageID, &item.Channel, &item.Title,
		&item.Notes, &scheduledFor, &item.Status, &item.CreatedAt, &item.UpdatedAt, &completedAt)
	applyTaskTimes(&item, scheduledFor, completedAt)
	return item, common.NormalizePostgresErr(err)
}

func (repo *PostgresRepository) UpdatePublishingTask(ctx context.Context, workspaceID string, campaignID string, taskID string, input UpdatePublishingTaskInput) (PublishingTask, error) {
	item := PublishingTask{}
	var scheduledFor sql.NullTime
	var completedAt sql.NullTime
	err := repo.db.QueryRow(ctx, `
		UPDATE publishing_tasks pt
		SET status = COALESCE($4, pt.status),
		    notes = COALESCE($5, pt.notes),
		    scheduled_for = COALESCE($6, pt.scheduled_for),
		    completed_at = CASE WHEN $4 = 'done' THEN now() WHEN $4 IS NULL THEN pt.completed_at ELSE NULL END,
		    updated_at = now()
		FROM campaigns c
		WHERE pt.workspace_id = $1 AND pt.campaign_id = $2 AND pt.id = $3
		  AND c.id = pt.campaign_id AND c.workspace_id = pt.workspace_id AND c.archived_at IS NULL
		RETURNING pt.id, pt.workspace_id, pt.campaign_id, COALESCE(pt.channel_package_id, ''),
		          pt.channel, pt.title, COALESCE(pt.notes, ''), pt.scheduled_for, pt.status,
		          pt.created_at, pt.updated_at, pt.completed_at`,
		workspaceID, campaignID, taskID, input.Status, input.Notes, input.ScheduledFor,
	).Scan(&item.ID, &item.WorkspaceID, &item.CampaignID, &item.ChannelPackageID, &item.Channel, &item.Title,
		&item.Notes, &scheduledFor, &item.Status, &item.CreatedAt, &item.UpdatedAt, &completedAt)
	applyTaskTimes(&item, scheduledFor, completedAt)
	return item, common.NormalizePostgresErr(err)
}

func applyTaskTimes(item *PublishingTask, scheduledFor sql.NullTime, completedAt sql.NullTime) {
	if scheduledFor.Valid {
		item.ScheduledFor = &scheduledFor.Time
	}
	if completedAt.Valid {
		item.CompletedAt = &completedAt.Time
	}
}

func (repo *PostgresRepository) getCampaign(ctx context.Context, workspaceID string, campaignID string) (Campaign, error) {
	var item Campaign
	err := repo.db.QueryRow(ctx, `
		SELECT id, workspace_id, COALESCE(product_id, ''), name, status, created_at, updated_at
		FROM campaigns
		WHERE workspace_id = $1 AND id = $2 AND archived_at IS NULL`,
		workspaceID, campaignID,
	).Scan(&item.ID, &item.WorkspaceID, &item.ProductID, &item.Name, &item.Status, &item.CreatedAt, &item.UpdatedAt)
	return item, common.NormalizePostgresErr(err)
}

func (repo *PostgresRepository) listChannelPackages(ctx context.Context, workspaceID string, campaignID string) ([]ChannelPackage, error) {
	rows, err := repo.db.Query(ctx, `
		SELECT id, workspace_id, campaign_id, channel, COALESCE(title, ''),
		       COALESCE(body, ''), status, created_at, updated_at
		FROM channel_packages
		WHERE workspace_id = $1 AND campaign_id = $2
		ORDER BY created_at DESC`, workspaceID, campaignID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []ChannelPackage
	for rows.Next() {
		var item ChannelPackage
		if err := rows.Scan(&item.ID, &item.WorkspaceID, &item.CampaignID, &item.Channel, &item.Title, &item.Body, &item.Status, &item.CreatedAt, &item.UpdatedAt); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, rows.Err()
}
