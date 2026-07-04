package analytics

import "testing"

func TestParseConversionCSV(t *testing.T) {
	rows, err := ParseConversionCSV("product_id,order_reference,gross_amount_cents,commission_cents,currency,occurred_at\nprd_1,ord_1,129900,12990,brl,2026-07-04\n")
	if err != nil {
		t.Fatalf("ParseConversionCSV error = %v", err)
	}
	if len(rows) != 1 {
		t.Fatalf("rows len = %d, want 1", len(rows))
	}
	if rows[0].ProductID != "prd_1" || rows[0].OrderReference != "ord_1" || rows[0].Currency != "BRL" {
		t.Fatalf("row = %+v", rows[0])
	}
	if rows[0].GrossAmountCents == nil || *rows[0].GrossAmountCents != 129900 {
		t.Fatalf("gross amount = %+v", rows[0].GrossAmountCents)
	}
	if rows[0].OccurredAt == nil {
		t.Fatal("occurred_at was not parsed")
	}
}

func TestParseConversionCSVWithAffiliateLinkOnly(t *testing.T) {
	rows, err := ParseConversionCSV("affiliate_link_id,order_reference,gross_amount_cents\nlnk_1,ord_1,129900\n")
	if err != nil {
		t.Fatalf("ParseConversionCSV error = %v", err)
	}
	if len(rows) != 1 {
		t.Fatalf("rows len = %d, want 1", len(rows))
	}
	if rows[0].AffiliateLinkID != "lnk_1" || rows[0].ProductID != "" {
		t.Fatalf("row = %+v", rows[0])
	}
}

func TestParseConversionCSVValidation(t *testing.T) {
	_, err := ParseConversionCSV("product_id,gross_amount_cents\n,-1\n")
	if err == nil {
		t.Fatal("expected validation error")
	}
}
