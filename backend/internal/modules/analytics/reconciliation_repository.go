package analytics

import (
	"context"
	"database/sql"

	"github.com/JeanVCN/affiliate_saas/backend/internal/modules/common"
)

func (repo *PostgresRepository) ReconciliationSummary(ctx context.Context, workspaceID string, importID string) (ReconciliationSummary, error) {
	var item ReconciliationSummary
	err := repo.db.QueryRow(ctx, `
		SELECT COUNT(*)::bigint,
		       COUNT(*) FILTER (WHERE reconciliation_status = 'pending')::bigint,
		       COUNT(*) FILTER (WHERE reconciliation_status = 'matched')::bigint,
		       COUNT(*) FILTER (WHERE reconciliation_status = 'unmatched')::bigint,
		       COUNT(*) FILTER (WHERE reconciliation_status = 'ignored')::bigint
		FROM conversion_import_rows
		WHERE workspace_id = $1 AND conversion_import_id = $2`,
		workspaceID, importID,
	).Scan(&item.Total, &item.Pending, &item.Matched, &item.Unmatched, &item.Ignored)
	return item, err
}

func (repo *PostgresRepository) UpdateConversionImportRow(ctx context.Context, workspaceID string, importID string, rowID string, input UpdateConversionImportRowInput) (ConversionImportRow, error) {
	var item ConversionImportRow
	var occurredAt sql.NullTime
	var reconciledAt sql.NullTime
	productID := nullableString(input.ProductID)
	affiliateLinkID := nullableString(input.AffiliateLinkID)
	reconciliationStatus := nullableString(input.ReconciliationStatus)
	reconciliationNote := nullableString(input.ReconciliationNote)
	err := repo.db.QueryRow(ctx, `
		WITH resolved AS (
			SELECT
				COALESCE(NULLIF($4::text, ''), (
					SELECT al.product_id
					FROM affiliate_links al
					WHERE al.id = NULLIF($5::text, '')
					  AND al.workspace_id = $1
					  AND al.archived_at IS NULL
				)) AS product_id,
				NULLIF($5::text, '') AS affiliate_link_id
		)
		UPDATE conversion_import_rows cir
		SET product_id = COALESCE(r.product_id, cir.product_id),
		    affiliate_link_id = COALESCE(r.affiliate_link_id, cir.affiliate_link_id),
		    reconciliation_status = COALESCE($6::text, CASE
		      WHEN COALESCE(r.product_id, cir.product_id) IS NOT NULL
		        OR COALESCE(r.affiliate_link_id, cir.affiliate_link_id) IS NOT NULL
		      THEN 'matched'
		      ELSE cir.reconciliation_status
		    END),
		    reconciliation_note = COALESCE($7::text, cir.reconciliation_note),
		    reconciled_at = CASE
		      WHEN $6::text IS NOT NULL OR $4::text IS NOT NULL OR $5::text IS NOT NULL OR $7::text IS NOT NULL THEN now()
		      ELSE cir.reconciled_at
		    END
		FROM resolved r
		WHERE cir.workspace_id = $1
		  AND cir.conversion_import_id = $2
		  AND cir.id = $3
		  AND (
		    r.product_id IS NULL
		    OR EXISTS (
		      SELECT 1 FROM products p
		      WHERE p.id = r.product_id AND p.workspace_id = $1 AND p.archived_at IS NULL
		    )
		  )
		  AND (
		    COALESCE(r.affiliate_link_id, cir.affiliate_link_id) IS NULL
		    OR EXISTS (
		      SELECT 1 FROM affiliate_links al
		      WHERE al.id = COALESCE(r.affiliate_link_id, cir.affiliate_link_id)
		        AND al.workspace_id = $1
		        AND al.archived_at IS NULL
		        AND al.product_id = COALESCE(r.product_id, cir.product_id)
		    )
		  )
		RETURNING cir.id, cir.workspace_id, cir.conversion_import_id, COALESCE(cir.product_id, ''),
		          COALESCE(cir.affiliate_link_id, ''), cir.occurred_at, COALESCE(cir.order_reference, ''),
		          cir.gross_amount_cents, cir.commission_cents, COALESCE(cir.currency, ''), cir.raw_payload,
		          cir.reconciliation_status, COALESCE(cir.reconciliation_note, ''), cir.reconciled_at, cir.created_at`,
		workspaceID, importID, rowID, productID, affiliateLinkID, reconciliationStatus, reconciliationNote,
	).Scan(&item.ID, &item.WorkspaceID, &item.ConversionImportID, &item.ProductID, &item.AffiliateLinkID,
		&occurredAt, &item.OrderReference, &item.GrossAmountCents, &item.CommissionCents,
		&item.Currency, &item.RawPayload, &item.ReconciliationStatus, &item.ReconciliationNote,
		&reconciledAt, &item.CreatedAt)
	if occurredAt.Valid {
		item.OccurredAt = &occurredAt.Time
	}
	if reconciledAt.Valid {
		item.ReconciledAt = &reconciledAt.Time
	}
	return item, common.NormalizePostgresErr(err)
}

func nullableString(value *string) any {
	if value == nil {
		return nil
	}
	return *value
}
