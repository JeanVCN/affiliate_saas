import type {
  AffiliateLink,
  AnalyticsOverview,
  ApiErrorBody,
  AuthResponse,
  MarketplaceProgram,
  Offer,
  Product,
  ShortLink
} from "@/types/api";

const API_BASE_URL = process.env.NEXT_PUBLIC_API_BASE_URL ?? "/backend";

type RequestOptions = {
  method?: "GET" | "POST" | "PATCH";
  body?: unknown;
};

export class ApiClientError extends Error {
  code: string;
  fields?: Record<string, string>;
  status: number;

  constructor(message: string, status: number, code = "request_failed", fields?: Record<string, string>) {
    super(message);
    this.name = "ApiClientError";
    this.code = code;
    this.status = status;
    this.fields = fields;
  }
}

async function request<T>(path: string, options: RequestOptions = {}): Promise<T> {
  const response = await fetch(`${API_BASE_URL}${path}`, {
    method: options.method ?? "GET",
    credentials: "include",
    headers: {
      Accept: "application/json",
      ...(options.body ? { "Content-Type": "application/json" } : {})
    },
    body: options.body ? JSON.stringify(options.body) : undefined
  });

  if (!response.ok) {
    let payload: ApiErrorBody | undefined;
    try {
      payload = (await response.json()) as ApiErrorBody;
    } catch {
      payload = undefined;
    }

    throw new ApiClientError(
      payload?.error?.message ?? `Request failed with status ${response.status}`,
      response.status,
      payload?.error?.code,
      payload?.error?.fields
    );
  }

  if (response.status === 204) {
    return undefined as T;
  }

  return (await response.json()) as T;
}

export const api = {
  signup(input: {
    email: string;
    password: string;
    display_name: string;
    workspace_name: string;
    workspace_slug: string;
  }) {
    return request<AuthResponse>("/api/v1/auth/signup", { method: "POST", body: input });
  },
  login(input: { email: string; password: string }) {
    return request<AuthResponse>("/api/v1/auth/login", { method: "POST", body: input });
  },
  me() {
    return request<AuthResponse>("/api/v1/auth/me");
  },
  logout() {
    return request<void>("/api/v1/auth/logout", { method: "POST" });
  },
  async listPrograms(workspaceID: string) {
    const programs = await request<MarketplaceProgram[] | null>(`/api/v1/workspaces/${workspaceID}/programs`);
    return programs ?? [];
  },
  enableProgram(
    workspaceID: string,
    input: {
      marketplace_name: string;
      marketplace_slug: string;
      program_name: string;
      program_slug: string;
      external_account_label: string;
    }
  ) {
    return request<MarketplaceProgram>(`/api/v1/workspaces/${workspaceID}/programs`, {
      method: "POST",
      body: input
    });
  },
  createProduct(workspaceID: string, input: { name: string; category: string; description: string }) {
    return request<Product>(`/api/v1/workspaces/${workspaceID}/products`, { method: "POST", body: input });
  },
  createOffer(
    workspaceID: string,
    productID: string,
    input: { workspace_program_id: string; title: string; price_cents: number; currency: string }
  ) {
    return request<Offer>(`/api/v1/workspaces/${workspaceID}/products/${productID}/offers`, {
      method: "POST",
      body: input
    });
  },
  createLink(
    workspaceID: string,
    input: { product_id: string; offer_id: string; destination_url: string; label: string }
  ) {
    return request<AffiliateLink>(`/api/v1/workspaces/${workspaceID}/links`, { method: "POST", body: input });
  },
  createShortLink(workspaceID: string, linkID: string, input: { slug: string }) {
    return request<ShortLink>(`/api/v1/workspaces/${workspaceID}/links/${linkID}/short-links`, {
      method: "POST",
      body: input
    });
  },
  analyticsOverview(workspaceID: string) {
    return request<AnalyticsOverview>(`/api/v1/workspaces/${workspaceID}/analytics/overview`);
  }
};

export function getErrorMessage(error: unknown) {
  if (error instanceof ApiClientError) {
    const fieldMessages = error.fields ? Object.values(error.fields).join(" ") : "";
    return fieldMessages ? `${error.message} ${fieldMessages}` : error.message;
  }
  if (error instanceof Error) {
    return error.message;
  }
  return "Unexpected error.";
}
