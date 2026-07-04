package analytics

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/JeanVCN/affiliate_saas/backend/internal/modules/common"
)

type ClickMetric struct {
	Group      string `json:"group"`
	GroupID    string `json:"group_id"`
	GroupLabel string `json:"group_label"`
	Clicks     int64  `json:"clicks"`
}

type Overview struct {
	Clicks              int64 `json:"clicks"`
	ImportedConversions int64 `json:"imported_conversions"`
	GrossAmountCents    int64 `json:"gross_amount_cents"`
	CommissionCents     int64 `json:"commission_cents"`
}

type TopProduct struct {
	ProductID           string `json:"product_id"`
	ProductName         string `json:"product_name"`
	Clicks              int64  `json:"clicks"`
	ImportedConversions int64  `json:"imported_conversions"`
	GrossAmountCents    int64  `json:"gross_amount_cents"`
	CommissionCents     int64  `json:"commission_cents"`
}

type ConversionImport struct {
	ID          string                 `json:"id"`
	WorkspaceID string                 `json:"workspace_id"`
	Source      string                 `json:"source"`
	Status      string                 `json:"status"`
	CreatedAt   time.Time              `json:"created_at"`
	UpdatedAt   time.Time              `json:"updated_at"`
	Rows        []ConversionImportRow  `json:"rows,omitempty"`
	Summary     *ReconciliationSummary `json:"reconciliation_summary,omitempty"`
}

type ConversionImportRow struct {
	ID                   string          `json:"id"`
	WorkspaceID          string          `json:"workspace_id"`
	ConversionImportID   string          `json:"conversion_import_id"`
	ProductID            string          `json:"product_id,omitempty"`
	AffiliateLinkID      string          `json:"affiliate_link_id,omitempty"`
	OccurredAt           *time.Time      `json:"occurred_at,omitempty"`
	OrderReference       string          `json:"order_reference,omitempty"`
	GrossAmountCents     *int            `json:"gross_amount_cents,omitempty"`
	CommissionCents      *int            `json:"commission_cents,omitempty"`
	Currency             string          `json:"currency,omitempty"`
	RawPayload           json.RawMessage `json:"raw_payload"`
	ReconciliationStatus string          `json:"reconciliation_status"`
	ReconciliationNote   string          `json:"reconciliation_note,omitempty"`
	ReconciledAt         *time.Time      `json:"reconciled_at,omitempty"`
	CreatedAt            time.Time       `json:"created_at"`
}

type ReconciliationSummary struct {
	Total     int64 `json:"total"`
	Pending   int64 `json:"pending"`
	Matched   int64 `json:"matched"`
	Unmatched int64 `json:"unmatched"`
	Ignored   int64 `json:"ignored"`
}

type CreateConversionImportInput struct {
	Source string `json:"source"`
}

type CreateConversionImportRowInput struct {
	ProductID        string         `json:"product_id"`
	AffiliateLinkID  string         `json:"affiliate_link_id"`
	OccurredAt       *time.Time     `json:"occurred_at"`
	OrderReference   string         `json:"order_reference"`
	GrossAmountCents *int           `json:"gross_amount_cents"`
	CommissionCents  *int           `json:"commission_cents"`
	Currency         string         `json:"currency"`
	RawPayload       map[string]any `json:"raw_payload"`
}

type CreateConversionImportCSVInput struct {
	CSV string `json:"csv"`
}

type UpdateConversionImportRowInput struct {
	ProductID            *string `json:"product_id"`
	AffiliateLinkID      *string `json:"affiliate_link_id"`
	ReconciliationStatus *string `json:"reconciliation_status"`
	ReconciliationNote   *string `json:"reconciliation_note"`
}

func (input *CreateConversionImportInput) Normalize() {
	input.Source = strings.ToLower(strings.TrimSpace(input.Source))
	if input.Source == "" {
		input.Source = "manual"
	}
}

func (input CreateConversionImportInput) Validate() error {
	if input.Source != "manual" && input.Source != "csv" {
		return common.NewValidationError("source must be manual or csv")
	}
	return nil
}

func (input *CreateConversionImportRowInput) Normalize() {
	input.ProductID = strings.TrimSpace(input.ProductID)
	input.AffiliateLinkID = strings.TrimSpace(input.AffiliateLinkID)
	input.OrderReference = strings.TrimSpace(input.OrderReference)
	input.Currency = strings.ToUpper(strings.TrimSpace(input.Currency))
}

func (input CreateConversionImportRowInput) Validate() error {
	if input.GrossAmountCents != nil && *input.GrossAmountCents < 0 {
		return common.NewValidationError("gross_amount_cents must be greater than or equal to zero")
	}
	if input.CommissionCents != nil && *input.CommissionCents < 0 {
		return common.NewValidationError("commission_cents must be greater than or equal to zero")
	}
	if input.ProductID == "" && input.AffiliateLinkID == "" && input.OrderReference == "" {
		return common.NewValidationError("product_id, affiliate_link_id, or order_reference is required")
	}
	return nil
}

func (input *UpdateConversionImportRowInput) Normalize() {
	trimStringPtr(&input.ProductID)
	trimStringPtr(&input.AffiliateLinkID)
	trimStringPtr(&input.ReconciliationNote)
	if input.ReconciliationStatus != nil {
		value := strings.ToLower(strings.TrimSpace(*input.ReconciliationStatus))
		input.ReconciliationStatus = &value
	}
}

func (input UpdateConversionImportRowInput) Validate() error {
	if input.ReconciliationStatus != nil && !isReconciliationStatus(*input.ReconciliationStatus) {
		return common.NewValidationError("reconciliation_status must be pending, matched, unmatched, or ignored")
	}
	return nil
}

func isReconciliationStatus(value string) bool {
	return value == "pending" || value == "matched" || value == "unmatched" || value == "ignored"
}

func trimStringPtr(value **string) {
	if *value == nil {
		return
	}
	trimmed := strings.TrimSpace(**value)
	*value = &trimmed
}
