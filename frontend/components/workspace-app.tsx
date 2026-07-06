"use client";

import {
  BarChart3,
  CheckCircle2,
  Link2,
  ListRestart,
  Loader2,
  LogOut,
  PackagePlus,
  Plug,
  RefreshCw,
  ShieldCheck
} from "lucide-react";
import { FormEvent, useEffect, useMemo, useState } from "react";
import { api, getErrorMessage } from "@/lib/api";
import type { AffiliateLink, AnalyticsOverview, AuthResponse, MarketplaceProgram, Offer, Product, ShortLink } from "@/types/api";

type AuthMode = "login" | "signup";
type Status = { type: "idle" | "loading" | "success" | "error"; message: string };
type BusyAction = "session" | "auth" | "logout" | "programs" | "program" | "product" | "offer" | "link" | "shortLink" | "analytics" | null;

const initialStatus: Status = { type: "idle", message: "" };

export function WorkspaceApp() {
  const [authMode, setAuthMode] = useState<AuthMode>("signup");
  const [auth, setAuth] = useState<AuthResponse | null>(null);
  const [status, setStatus] = useState<Status>({ type: "loading", message: "Restoring session..." });
  const [busyAction, setBusyAction] = useState<BusyAction>("session");
  const [email, setEmail] = useState("creator@example.com");
  const [password, setPassword] = useState("affiliate1234");
  const [displayName, setDisplayName] = useState("Creator Operator");
  const [workspaceName, setWorkspaceName] = useState("Creator Lab");
  const [workspaceSlug, setWorkspaceSlug] = useState("creator-lab");
  const [programs, setPrograms] = useState<MarketplaceProgram[]>([]);
  const [program, setProgram] = useState<MarketplaceProgram | null>(null);
  const [product, setProduct] = useState<Product | null>(null);
  const [offer, setOffer] = useState<Offer | null>(null);
  const [link, setLink] = useState<AffiliateLink | null>(null);
  const [shortLink, setShortLink] = useState<ShortLink | null>(null);
  const [overview, setOverview] = useState<AnalyticsOverview | null>(null);

  const [programForm, setProgramForm] = useState({
    marketplace_name: "TikTok Shop",
    marketplace_slug: "tiktok-shop",
    program_name: "TikTok Shop Affiliate",
    program_slug: "tiktok-shop-affiliate",
    external_account_label: "Manual account"
  });
  const [productForm, setProductForm] = useState({
    name: "Creator Camera Kit",
    category: "Creator gear",
    description: "Manual product record for affiliate campaign planning."
  });
  const [offerForm, setOfferForm] = useState({
    title: "TikTok Shop manual offer",
    price: "199.90",
    currency: "BRL"
  });
  const [linkForm, setLinkForm] = useState({
    destination_url: "https://example.com/products/creator-camera-kit",
    label: "TikTok bio link"
  });
  const [shortLinkSlug, setShortLinkSlug] = useState("creator-camera");

  const workspaceID = auth?.workspaces[0]?.workspace_id;
  const isBusy = busyAction !== null;

  useEffect(() => {
    api
      .me()
      .then((response) => {
        setAuth(response);
        setStatus({ type: "success", message: "Session restored." });
      })
      .catch(() => {
        setStatus(initialStatus);
      })
      .finally(() => {
        setBusyAction(null);
      });
  }, []);

  useEffect(() => {
    if (!workspaceID) {
      return;
    }

    let isActive = true;
    setBusyAction("programs");

    api
      .listPrograms(workspaceID)
      .then((loadedPrograms) => {
        if (!isActive) {
          return;
        }
        setPrograms(loadedPrograms);
        setProgram((current) => current ?? loadedPrograms[0] ?? null);
      })
      .catch((error) => {
        if (isActive) {
          setStatus({ type: "error", message: getErrorMessage(error) });
        }
      })
      .finally(() => {
        if (isActive) {
          setBusyAction(null);
        }
      });

    api
      .analyticsOverview(workspaceID)
      .then((nextOverview) => {
        if (isActive) {
          setOverview(nextOverview);
        }
      })
      .catch((error) => {
        if (isActive) {
          setStatus({ type: "error", message: getErrorMessage(error) });
        }
      });

    return () => {
      isActive = false;
    };
  }, [workspaceID]);

  const steps = useMemo(
    () => [
      { label: "Workspace", value: workspaceID ?? "Waiting for session", done: Boolean(workspaceID), icon: ShieldCheck },
      { label: "Marketplace program", value: program?.program_name ?? `${programs.length} loaded`, done: Boolean(program), icon: Plug },
      { label: "Product", value: product?.name ?? "Not created", done: Boolean(product), icon: PackagePlus },
      { label: "Offer", value: offer?.title ?? "Not created", done: Boolean(offer), icon: CheckCircle2 },
      { label: "Affiliate link", value: link?.label ?? "Not created", done: Boolean(link), icon: Link2 },
      { label: "Short link", value: shortLink ? `/r/${shortLink.slug}` : "Not created", done: Boolean(shortLink), icon: Link2 },
      { label: "Analytics", value: overview ? `${overview.clicks} clicks` : "Not loaded", done: Boolean(overview), icon: BarChart3 }
    ],
    [link, offer, overview, product, program, programs.length, shortLink, workspaceID]
  );

  async function submitAuth(event: FormEvent<HTMLFormElement>) {
    event.preventDefault();
    setBusyAction("auth");
    setStatus({ type: "loading", message: authMode === "signup" ? "Creating account..." : "Signing in..." });
    try {
      const response =
        authMode === "signup"
          ? await api.signup({
              email,
              password,
              display_name: displayName,
              workspace_name: workspaceName,
              workspace_slug: workspaceSlug
            })
          : await api.login({ email, password });
      resetWorkspaceState();
      setAuth(response);
      setStatus({ type: "success", message: authMode === "signup" ? "Account ready." : "Signed in." });
    } catch (error) {
      setStatus({ type: "error", message: getErrorMessage(error) });
    } finally {
      setBusyAction(null);
    }
  }

  async function logout() {
    setBusyAction("logout");
    setStatus({ type: "loading", message: "Signing out..." });
    try {
      await api.logout();
      setAuth(null);
      resetWorkspaceState();
      setStatus({ type: "success", message: "Signed out." });
    } catch (error) {
      setStatus({ type: "error", message: getErrorMessage(error) });
    } finally {
      setBusyAction(null);
    }
  }

  async function loadPrograms(targetWorkspaceID = workspaceID, options: { quiet?: boolean } = {}) {
    if (!targetWorkspaceID) {
      setStatus({ type: "error", message: "Authenticate before loading programs." });
      return;
    }

    setBusyAction("programs");
    if (!options.quiet) {
      setStatus({ type: "loading", message: "Loading marketplace programs..." });
    }

    try {
      const loadedPrograms = await api.listPrograms(targetWorkspaceID);
      setPrograms(loadedPrograms);
      setProgram((current) => current ?? loadedPrograms[0] ?? null);
      if (!options.quiet) {
        setStatus({ type: "success", message: loadedPrograms.length ? "Programs loaded." : "No programs yet." });
      }
    } catch (error) {
      setStatus({ type: "error", message: getErrorMessage(error) });
    } finally {
      setBusyAction(null);
    }
  }

  async function submitProgram(event: FormEvent<HTMLFormElement>) {
    event.preventDefault();
    if (!workspaceID) {
      setStatus({ type: "error", message: "Authenticate before creating a program." });
      return;
    }

    setBusyAction("program");
    setStatus({ type: "loading", message: "Creating marketplace program..." });
    try {
      const createdProgram = await api.enableProgram(workspaceID, programForm);
      setProgram(createdProgram);
      setPrograms((current) => [createdProgram, ...current.filter((item) => item.id !== createdProgram.id)]);
      setStatus({ type: "success", message: "Program ready." });
    } catch (error) {
      setStatus({ type: "error", message: getErrorMessage(error) });
    } finally {
      setBusyAction(null);
    }
  }

  async function submitProduct(event: FormEvent<HTMLFormElement>) {
    event.preventDefault();
    if (!workspaceID) {
      setStatus({ type: "error", message: "Authenticate before creating a product." });
      return;
    }

    setBusyAction("product");
    setStatus({ type: "loading", message: "Creating product..." });
    try {
      const createdProduct = await api.createProduct(workspaceID, productForm);
      setProduct(createdProduct);
      setOffer(null);
      setLink(null);
      setShortLink(null);
      setStatus({ type: "success", message: "Product created." });
    } catch (error) {
      setStatus({ type: "error", message: getErrorMessage(error) });
    } finally {
      setBusyAction(null);
    }
  }

  async function submitOffer(event: FormEvent<HTMLFormElement>) {
    event.preventDefault();
    if (!workspaceID || !program || !product) {
      setStatus({ type: "error", message: "Create a program and product before creating an offer." });
      return;
    }

    setBusyAction("offer");
    setStatus({ type: "loading", message: "Creating offer..." });
    try {
      const createdOffer = await api.createOffer(workspaceID, product.id, {
        workspace_program_id: program.id,
        title: offerForm.title,
        price_cents: toCents(offerForm.price),
        currency: offerForm.currency
      });
      setOffer(createdOffer);
      setLink(null);
      setShortLink(null);
      setStatus({ type: "success", message: "Offer created." });
    } catch (error) {
      setStatus({ type: "error", message: getErrorMessage(error) });
    } finally {
      setBusyAction(null);
    }
  }

  async function submitLink(event: FormEvent<HTMLFormElement>) {
    event.preventDefault();
    if (!workspaceID || !product || !offer) {
      setStatus({ type: "error", message: "Create a product and offer before creating a link." });
      return;
    }

    setBusyAction("link");
    setStatus({ type: "loading", message: "Creating affiliate link..." });
    try {
      const createdLink = await api.createLink(workspaceID, {
        product_id: product.id,
        offer_id: offer.id,
        destination_url: linkForm.destination_url,
        label: linkForm.label
      });
      setLink(createdLink);
      setShortLink(null);
      setStatus({ type: "success", message: "Affiliate link created." });
    } catch (error) {
      setStatus({ type: "error", message: getErrorMessage(error) });
    } finally {
      setBusyAction(null);
    }
  }

  async function submitShortLink(event: FormEvent<HTMLFormElement>) {
    event.preventDefault();
    if (!workspaceID || !link) {
      setStatus({ type: "error", message: "Create an affiliate link before creating a short link." });
      return;
    }

    setBusyAction("shortLink");
    setStatus({ type: "loading", message: "Creating short link..." });
    try {
      const createdShortLink = await api.createShortLink(workspaceID, link.id, {
        slug: `${shortLinkSlug}-${Date.now()}`
      });
      setShortLink(createdShortLink);
      setStatus({ type: "success", message: "Short link created." });
    } catch (error) {
      setStatus({ type: "error", message: getErrorMessage(error) });
    } finally {
      setBusyAction(null);
    }
  }

  async function refreshAnalytics(targetWorkspaceID = workspaceID, options: { quiet?: boolean } = {}) {
    if (!targetWorkspaceID) {
      setStatus({ type: "error", message: "Authenticate before loading analytics." });
      return;
    }

    setBusyAction("analytics");
    if (!options.quiet) {
      setStatus({ type: "loading", message: "Refreshing analytics..." });
    }

    try {
      const nextOverview = await api.analyticsOverview(targetWorkspaceID);
      setOverview(nextOverview);
      if (!options.quiet) {
        setStatus({ type: "success", message: "Analytics refreshed." });
      }
    } catch (error) {
      setStatus({ type: "error", message: getErrorMessage(error) });
    } finally {
      setBusyAction(null);
    }
  }

  function resetWorkspaceState() {
    setPrograms([]);
    setProgram(null);
    setProduct(null);
    setOffer(null);
    setLink(null);
    setShortLink(null);
    setOverview(null);
  }

  return (
    <main className="app-shell">
      <section className="sidebar">
        <div>
          <p className="eyebrow">Affiliate SaaS</p>
          <h1>Commerce workspace</h1>
        </div>
        <nav aria-label="Workspace flow">
          {steps.map((step) => {
            const Icon = step.icon;
            return (
              <div className={`nav-step ${step.done ? "is-done" : ""}`} key={step.label}>
                <Icon aria-hidden="true" size={18} />
                <span>
                  <strong>{step.label}</strong>
                  <small>{step.value}</small>
                </span>
              </div>
            );
          })}
        </nav>
      </section>

      <section className="workspace">
        <header className="topbar">
          <div>
            <p className="eyebrow">Manual MVP flow</p>
            <h2>Product to analytics</h2>
          </div>
          {auth ? (
            <button className="icon-button secondary" disabled={isBusy} onClick={logout} type="button" title="Logout">
              {busyAction === "logout" ? <Loader2 className="spin" aria-hidden="true" size={18} /> : <LogOut aria-hidden="true" size={18} />}
              Logout
            </button>
          ) : null}
        </header>

        <StatusBanner status={status} />

        {!auth ? (
          <form className="auth-panel" onSubmit={submitAuth}>
            <div className="segmented" aria-label="Authentication mode">
              <button className={authMode === "signup" ? "active" : ""} onClick={() => setAuthMode("signup")} type="button">
                Signup
              </button>
              <button className={authMode === "login" ? "active" : ""} onClick={() => setAuthMode("login")} type="button">
                Login
              </button>
            </div>

            <div className="form-grid">
              <label>
                Email
                <input autoComplete="email" onChange={(event) => setEmail(event.target.value)} value={email} />
              </label>
              <label>
                Password
                <input autoComplete={authMode === "signup" ? "new-password" : "current-password"} onChange={(event) => setPassword(event.target.value)} type="password" value={password} />
              </label>
              {authMode === "signup" ? (
                <>
                  <label>
                    Display name
                    <input onChange={(event) => setDisplayName(event.target.value)} value={displayName} />
                  </label>
                  <label>
                    Workspace
                    <input onChange={(event) => setWorkspaceName(event.target.value)} value={workspaceName} />
                  </label>
                  <label>
                    Workspace slug
                    <input onChange={(event) => setWorkspaceSlug(event.target.value)} value={workspaceSlug} />
                  </label>
                </>
              ) : null}
            </div>

            <button className="primary-action" disabled={isBusy} type="submit">
              {busyAction === "auth" || busyAction === "session" ? <Loader2 className="spin" aria-hidden="true" size={18} /> : <ShieldCheck aria-hidden="true" size={18} />}
              {authMode === "signup" ? "Create workspace" : "Enter workspace"}
            </button>
          </form>
        ) : (
          <div className="manual-workspace">
            <section className="workspace-summary">
              <div>
                <p className="eyebrow">Authenticated</p>
                <h3>{auth.user.display_name || auth.user.email}</h3>
                <p className="muted">Workspace ID: {workspaceID}</p>
              </div>
              <button className="icon-button secondary" disabled={isBusy || !workspaceID} onClick={() => void loadPrograms()} type="button">
                {busyAction === "programs" ? <Loader2 className="spin" aria-hidden="true" size={18} /> : <ListRestart aria-hidden="true" size={18} />}
                Load programs
              </button>
            </section>

            <div className="flow-panels">
              <form className="flow-card" onSubmit={submitProgram}>
                <PanelHeader done={Boolean(program)} eyebrow="Step 1" title="Marketplace program" />
                <div className="form-grid compact">
                  <label>
                    Marketplace
                    <input onChange={(event) => setProgramForm({ ...programForm, marketplace_name: event.target.value })} value={programForm.marketplace_name} />
                  </label>
                  <label>
                    Marketplace slug
                    <input onChange={(event) => setProgramForm({ ...programForm, marketplace_slug: event.target.value })} value={programForm.marketplace_slug} />
                  </label>
                  <label>
                    Program
                    <input onChange={(event) => setProgramForm({ ...programForm, program_name: event.target.value })} value={programForm.program_name} />
                  </label>
                  <label>
                    Program slug
                    <input onChange={(event) => setProgramForm({ ...programForm, program_slug: event.target.value })} value={programForm.program_slug} />
                  </label>
                  <label className="span-2">
                    External account label
                    <input onChange={(event) => setProgramForm({ ...programForm, external_account_label: event.target.value })} value={programForm.external_account_label} />
                  </label>
                </div>
                <ResourceSummary items={programs.map((item) => `${item.program_name} (${item.status})`)} empty="No programs loaded." />
                <button className="primary-action" disabled={isBusy || !workspaceID} type="submit">
                  {busyAction === "program" ? <Loader2 className="spin" aria-hidden="true" size={18} /> : <Plug aria-hidden="true" size={18} />}
                  Save program
                </button>
              </form>

              <form className="flow-card" onSubmit={submitProduct}>
                <PanelHeader done={Boolean(product)} eyebrow="Step 2" title="Product" />
                <div className="form-grid compact">
                  <label>
                    Name
                    <input onChange={(event) => setProductForm({ ...productForm, name: event.target.value })} value={productForm.name} />
                  </label>
                  <label>
                    Category
                    <input onChange={(event) => setProductForm({ ...productForm, category: event.target.value })} value={productForm.category} />
                  </label>
                  <label className="span-2">
                    Description
                    <input onChange={(event) => setProductForm({ ...productForm, description: event.target.value })} value={productForm.description} />
                  </label>
                </div>
                <ResourceSummary items={product ? [`${product.name} (${product.status})`] : []} empty="No product selected." />
                <button className="primary-action" disabled={isBusy || !workspaceID} type="submit">
                  {busyAction === "product" ? <Loader2 className="spin" aria-hidden="true" size={18} /> : <PackagePlus aria-hidden="true" size={18} />}
                  Save product
                </button>
              </form>

              <form className="flow-card" onSubmit={submitOffer}>
                <PanelHeader done={Boolean(offer)} eyebrow="Step 3" title="Offer" />
                <div className="form-grid compact">
                  <label>
                    Title
                    <input onChange={(event) => setOfferForm({ ...offerForm, title: event.target.value })} value={offerForm.title} />
                  </label>
                  <label>
                    Price
                    <input inputMode="decimal" onChange={(event) => setOfferForm({ ...offerForm, price: event.target.value })} value={offerForm.price} />
                  </label>
                  <label>
                    Currency
                    <input onChange={(event) => setOfferForm({ ...offerForm, currency: event.target.value })} value={offerForm.currency} />
                  </label>
                </div>
                <ResourceSummary items={offer ? [`${offer.title || "Offer"} (${formatMoney(offer.price_cents ?? 0)})`] : []} empty="Create program and product first." />
                <button className="primary-action" disabled={isBusy || !program || !product} type="submit">
                  {busyAction === "offer" ? <Loader2 className="spin" aria-hidden="true" size={18} /> : <CheckCircle2 aria-hidden="true" size={18} />}
                  Save offer
                </button>
              </form>

              <form className="flow-card" onSubmit={submitLink}>
                <PanelHeader done={Boolean(link)} eyebrow="Step 4" title="Affiliate link" />
                <div className="form-grid compact">
                  <label className="span-2">
                    Destination URL
                    <input onChange={(event) => setLinkForm({ ...linkForm, destination_url: event.target.value })} value={linkForm.destination_url} />
                  </label>
                  <label className="span-2">
                    Label
                    <input onChange={(event) => setLinkForm({ ...linkForm, label: event.target.value })} value={linkForm.label} />
                  </label>
                </div>
                <ResourceSummary items={link ? [`${link.label || "Affiliate link"} -> ${link.destination_url}`] : []} empty="Create product and offer first." />
                <button className="primary-action" disabled={isBusy || !product || !offer} type="submit">
                  {busyAction === "link" ? <Loader2 className="spin" aria-hidden="true" size={18} /> : <Link2 aria-hidden="true" size={18} />}
                  Save link
                </button>
              </form>

              <form className="flow-card" onSubmit={submitShortLink}>
                <PanelHeader done={Boolean(shortLink)} eyebrow="Step 5" title="Short link" />
                <div className="form-grid compact">
                  <label className="span-2">
                    Slug prefix
                    <input onChange={(event) => setShortLinkSlug(event.target.value)} value={shortLinkSlug} />
                  </label>
                </div>
                <ResourceSummary items={shortLink ? [`/r/${shortLink.slug}`] : []} empty="Create affiliate link first." />
                <button className="primary-action" disabled={isBusy || !link} type="submit">
                  {busyAction === "shortLink" ? <Loader2 className="spin" aria-hidden="true" size={18} /> : <Link2 aria-hidden="true" size={18} />}
                  Create short link
                </button>
              </form>

              <section className="flow-card analytics-card">
                <PanelHeader done={Boolean(overview)} eyebrow="Step 6" title="Analytics overview" />
                <section className="metrics-grid" aria-label="Analytics overview">
                  <Metric label="Clicks" value={overview?.clicks ?? 0} />
                  <Metric label="Conversions" value={overview?.imported_conversions ?? 0} />
                  <Metric label="Gross" value={formatMoney(overview?.gross_amount_cents ?? 0)} />
                  <Metric label="Commission" value={formatMoney(overview?.commission_cents ?? 0)} />
                </section>
                <button className="primary-action" disabled={isBusy || !workspaceID} onClick={() => void refreshAnalytics()} type="button">
                  {busyAction === "analytics" ? <Loader2 className="spin" aria-hidden="true" size={18} /> : <RefreshCw aria-hidden="true" size={18} />}
                  Refresh analytics
                </button>
              </section>
            </div>
          </div>
        )}
      </section>
    </main>
  );
}

function PanelHeader({ done, eyebrow, title }: { done: boolean; eyebrow: string; title: string }) {
  return (
    <div className="panel-header">
      <div>
        <p className="eyebrow">{eyebrow}</p>
        <h3>{title}</h3>
      </div>
      <span className={`state-pill ${done ? "done" : ""}`}>{done ? "Ready" : "Pending"}</span>
    </div>
  );
}

function ResourceSummary({ empty, items }: { empty: string; items: string[] }) {
  return (
    <div className="resource-summary">
      {items.length ? items.map((item) => <span key={item}>{item}</span>) : <span className="muted">{empty}</span>}
    </div>
  );
}

function StatusBanner({ status }: { status: Status }) {
  if (status.type === "idle" || !status.message) {
    return null;
  }

  return (
    <div className={`status ${status.type}`} role={status.type === "error" ? "alert" : "status"}>
      {status.type === "loading" ? <Loader2 className="spin" aria-hidden="true" size={18} /> : null}
      <span>{status.message}</span>
    </div>
  );
}

function Metric({ label, value }: { label: string; value: string | number }) {
  return (
    <div className="metric">
      <span>{label}</span>
      <strong>{value}</strong>
    </div>
  );
}

function toCents(value: string) {
  const numeric = Number(value.replace(",", "."));
  if (!Number.isFinite(numeric)) {
    return 0;
  }
  return Math.round(numeric * 100);
}

function formatMoney(cents: number) {
  return new Intl.NumberFormat("pt-BR", {
    style: "currency",
    currency: "BRL"
  }).format(cents / 100);
}
