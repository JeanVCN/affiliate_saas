package compliance

import (
	"context"
	"strings"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (service *Service) RunCampaignCheck(ctx context.Context, workspaceID string, campaignID string, input RunCampaignCheckInput) (Check, error) {
	input.Normalize()
	content := input.content()
	if content == "" {
		var err error
		content, err = service.repo.CampaignContent(ctx, workspaceID, campaignID)
		if err != nil {
			return Check{}, err
		}
	}
	findings := evaluateContent(content)
	return service.repo.CreateCampaignCheck(ctx, workspaceID, campaignID, findings)
}

func evaluateContent(content string) []FindingInput {
	content = strings.ToLower(content)
	findings := make([]FindingInput, 0)
	if !containsAny(content, "afiliado", "affiliate", "publi", "publicidade", "sponsored", "patrocinado") {
		findings = append(findings, FindingInput{
			Severity: "blocker",
			Code:     "missing_affiliate_disclosure",
			Message:  "Affiliate disclosure is missing or unclear.",
		})
	}
	if containsAny(content, "garantido", "guaranteed", "cura", "cure", "renda garantida", "guaranteed income") {
		findings = append(findings, FindingInput{
			Severity: "blocker",
			Code:     "unsupported_absolute_claim",
			Message:  "Content includes an absolute or regulated claim that needs proof before publishing.",
		})
	}
	if containsAny(content, "menor preço", "lowest price", "mais barato", "cheapest", "disponível agora", "always available") {
		findings = append(findings, FindingInput{
			Severity: "warning",
			Code:     "price_or_availability_claim",
			Message:  "Price or availability claims must be checked against current product truth before publishing.",
		})
	}
	if containsAny(content, "bot", "automação de clique", "auto click", "fake engagement", "engajamento artificial") {
		findings = append(findings, FindingInput{
			Severity: "blocker",
			Code:     "prohibited_automation",
			Message:  "Content references prohibited automation or artificial engagement.",
		})
	}
	if len(findings) == 0 {
		findings = append(findings, FindingInput{
			Severity: "info",
			Code:     "basic_check_passed",
			Message:  "Basic MVP compliance checklist did not find obvious blockers.",
		})
	}
	return findings
}

func containsAny(content string, needles ...string) bool {
	for _, needle := range needles {
		if strings.Contains(content, needle) {
			return true
		}
	}
	return false
}
