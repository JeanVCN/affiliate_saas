package integration_test

type authResponse struct {
	Workspaces []membershipResponse `json:"workspaces"`
}

type membershipResponse struct {
	WorkspaceID string `json:"workspace_id"`
}

type workspaceProgramResponse struct {
	ID        string `json:"id"`
	ProgramID string `json:"program_id"`
}

type productResponse struct {
	ID string `json:"id"`
}

type campaignResponse struct {
	ID              string                   `json:"id"`
	ProductID       string                   `json:"product_id"`
	Status          string                   `json:"status"`
	ChannelPackages []channelPackageResponse `json:"channel_packages"`
}

type channelPackageResponse struct {
	ID string `json:"id"`
}

type publishingTaskResponse struct {
	ID          string `json:"id"`
	Status      string `json:"status"`
	CompletedAt string `json:"completed_at"`
}

type complianceCheckResponse struct {
	ID       string                      `json:"id"`
	Findings []complianceFindingResponse `json:"findings"`
}

type complianceFindingResponse struct {
	Code string `json:"code"`
}

type programPolicyNoteResponse struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Severity string `json:"severity"`
}

type offerResponse struct {
	ID string `json:"id"`
}

type linkResponse struct {
	ID string `json:"id"`
}

type conversionImportResponse struct {
	ID     string                        `json:"id"`
	Source string                        `json:"source"`
	Rows   []conversionImportRowResponse `json:"rows"`
}

type conversionImportRowResponse struct {
	ID                   string `json:"id"`
	ProductID            string `json:"product_id"`
	AffiliateLinkID      string `json:"affiliate_link_id"`
	ReconciliationStatus string `json:"reconciliation_status"`
	ReconciliationNote   string `json:"reconciliation_note"`
}

type reconciliationSummaryResponse struct {
	Matched int64 `json:"matched"`
}

type conversionCSVRowsResponse struct {
	Rows []conversionImportRowResponse `json:"rows"`
}

type shortLinkResponse struct {
	Slug string `json:"slug"`
}

type clickMetricsResponse struct {
	GroupBy string            `json:"group_by"`
	Items   []clickMetricItem `json:"items"`
}

type clickMetricItem struct {
	GroupID string `json:"group_id"`
	Clicks  int64  `json:"clicks"`
}

type analyticsOverviewResponse struct {
	Clicks              int64 `json:"clicks"`
	ImportedConversions int64 `json:"imported_conversions"`
	CommissionCents     int64 `json:"commission_cents"`
}

type topProductResponse struct {
	ProductID           string `json:"product_id"`
	Clicks              int64  `json:"clicks"`
	ImportedConversions int64  `json:"imported_conversions"`
}
