export type ApiErrorBody = {
  error?: {
    code?: string;
    message?: string;
    fields?: Record<string, string>;
  };
};

export type User = {
  id: string;
  email: string;
  display_name?: string;
  status: string;
};

export type Membership = {
  id: string;
  workspace_id: string;
  user_id: string;
  role: string;
  status: string;
};

export type AuthResponse = {
  user: User;
  session?: {
    expires_at: string;
  };
  workspaces: Membership[];
};

export type MarketplaceProgram = {
  id: string;
  workspace_id: string;
  program_id: string;
  marketplace_name: string;
  marketplace_slug: string;
  program_name: string;
  program_slug: string;
  external_account_label?: string;
  status: string;
};

export type Product = {
  id: string;
  workspace_id: string;
  name: string;
  category?: string;
  description?: string;
  status: string;
};

export type Offer = {
  id: string;
  workspace_id: string;
  product_id: string;
  workspace_program_id: string;
  title?: string;
  price_cents?: number;
  currency?: string;
  status: string;
};

export type AffiliateLink = {
  id: string;
  workspace_id: string;
  product_id: string;
  offer_id?: string;
  destination_url: string;
  label?: string;
  status: string;
  short_links?: ShortLink[];
};

export type ShortLink = {
  id: string;
  workspace_id: string;
  affiliate_link_id: string;
  slug: string;
  status: string;
};

export type AnalyticsOverview = {
  clicks: number;
  imported_conversions: number;
  gross_amount_cents: number;
  commission_cents: number;
};
