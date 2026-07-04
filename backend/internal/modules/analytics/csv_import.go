package analytics

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/JeanVCN/affiliate_saas/backend/internal/modules/common"
)

func ParseConversionCSV(input string) ([]CreateConversionImportRowInput, error) {
	return ParseConversionCSVReader(strings.NewReader(input))
}

func ParseConversionCSVReader(input io.Reader) ([]CreateConversionImportRowInput, error) {
	reader := csv.NewReader(input)
	reader.TrimLeadingSpace = true
	headers, err := reader.Read()
	if err != nil {
		return nil, common.NewValidationError("csv header is required")
	}
	index := mapHeaders(headers)
	rows := make([]CreateConversionImportRowInput, 0)
	line := 1
	for {
		line++
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, common.NewValidationError(fmt.Sprintf("invalid csv row %d", line))
		}
		item, err := csvRecordToInput(index, record)
		if err != nil {
			return nil, common.NewValidationError(fmt.Sprintf("row %d: %s", line, err.Error()))
		}
		item.Normalize()
		if err := item.Validate(); err != nil {
			return nil, common.NewValidationError(fmt.Sprintf("row %d: %s", line, err.Error()))
		}
		rows = append(rows, item)
	}
	if len(rows) == 0 {
		return nil, common.NewValidationError("csv must contain at least one data row")
	}
	return rows, nil
}

func mapHeaders(headers []string) map[string]int {
	index := make(map[string]int, len(headers))
	for i, header := range headers {
		index[strings.ToLower(strings.TrimSpace(header))] = i
	}
	return index
}

func csvRecordToInput(index map[string]int, record []string) (CreateConversionImportRowInput, error) {
	var item CreateConversionImportRowInput
	item.ProductID = csvValue(index, record, "product_id")
	item.AffiliateLinkID = csvValue(index, record, "affiliate_link_id")
	item.OrderReference = csvValue(index, record, "order_reference")
	item.Currency = csvValue(index, record, "currency")
	item.RawPayload = map[string]any{"source": "csv"}
	if value := csvValue(index, record, "occurred_at"); value != "" {
		occurredAt, err := parseCSVTime(value)
		if err != nil {
			return item, common.NewValidationError("occurred_at must be RFC3339 or YYYY-MM-DD")
		}
		item.OccurredAt = &occurredAt
	}
	if value := csvValue(index, record, "gross_amount_cents"); value != "" {
		amount, err := strconv.Atoi(value)
		if err != nil {
			return item, common.NewValidationError("gross_amount_cents must be an integer")
		}
		item.GrossAmountCents = &amount
	}
	if value := csvValue(index, record, "commission_cents"); value != "" {
		amount, err := strconv.Atoi(value)
		if err != nil {
			return item, common.NewValidationError("commission_cents must be an integer")
		}
		item.CommissionCents = &amount
	}
	return item, nil
}

func csvValue(index map[string]int, record []string, key string) string {
	i, ok := index[key]
	if !ok || i >= len(record) {
		return ""
	}
	return strings.TrimSpace(record[i])
}

func parseCSVTime(value string) (time.Time, error) {
	if parsed, err := time.Parse(time.RFC3339, value); err == nil {
		return parsed, nil
	}
	return time.Parse("2006-01-02", value)
}
