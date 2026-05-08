# Moolabs Go SDK

Unified Go SDK for the Moolabs platform. One module, one client, one auth flow — covers both billing (CLS) and usage (Meter) operations.

```go
import (
    "context"
    "github.com/moolabs/moolabs-go"
)

ctx := context.Background()
client := moolabs.NewMoolabs(moolabs.Config{APIKey: "moo_live_..."})

// Billing / wallets / grants — routed to api.moolabs.com (CLS)
wallet, _, err := client.Cls.Wallets.CreateWallet(ctx).WalletCreate(in).Execute()
grants, _, err := client.Cls.Grants.ListGrants(ctx).Execute()

// Usage events / meters / subscriptions — routed to meter.moolabs.com
_, err = client.Meter.Events.IngestEvents(ctx).Body(events).Execute()
meters, _, err := client.Meter.Meters.ListMeters(ctx).Execute()
```

## Install

```bash
go get github.com/moolabs/moolabs-go@latest
```

Requires Go 1.18+.

For a specific version:

```bash
go get github.com/moolabs/moolabs-go@v0.1.0
```

## Authentication

Generate an API key in your Moolabs dashboard. The same key authenticates against both backends; the SDK handles routing internally.

```go
client := moolabs.NewMoolabs(moolabs.Config{
    APIKey: "moo_live_...",
})
```

For staging or private deployments, override the base URLs:

```go
client := moolabs.NewMoolabs(moolabs.Config{
    APIKey:       "moo_test_...",
    ClsBaseURL:   "https://staging-api.moolabs.com",
    MeterBaseURL: "https://staging-meter.moolabs.com",
})
```

## Two namespaces

The SDK exposes two top-level groups via the `*Moolabs` struct.

### `client.Cls.*` — CLS / billing operations

Routes all calls to `https://api.moolabs.com`. Fields:

| Field | Type | Purpose |
|---|---|---|
| `Cls.Wallets` | `*WalletsAPIService` | Wallet lifecycle, balances, transfers |
| `Cls.Grants` | `*GrantsAPIService` | Credit grants and grant policies |
| `Cls.Ledger` | `*LedgerAPIService` | Ledger entries, transfers, audit |
| `Cls.Alerts` | `*AlertsAPIService` | Balance / threshold alert subscriptions |
| `Cls.AutoTopup` | `*AutoTopupAPIService` | Auto-topup rules |
| `Cls.RateCards` | `*RateCardsAPIService` | Rate card definitions |
| `Cls.Rating` | `*RatingAPIService` | Rate-event scoring |
| `Cls.FxRates` | `*FxRatesAPIService` | FX rate lookups |
| `Cls.Portal` | `*PortalAPIService` | Portal token issuance (BFF flavor) |
| `Cls.Subscriptions` | `*SubscriptionsAPIService` | Subscription bindings (BFF flavor) |

### `client.Meter.*` — Usage / metering operations

Routes all calls to `https://meter.moolabs.com`. Fields:

| Field | Type | Purpose |
|---|---|---|
| `Meter.Events` | `*EventsAPIService` | Ingest usage events |
| `Meter.Meters` | `*MetersAPIService` | Define and query meters |
| `Meter.Customers` | `*CustomersAPIService` | Customer entity management |
| `Meter.Subscriptions` | `*MeterSubscriptionsAPIService` | Meter-side subscriptions |
| `Meter.Billing` | `*MeterBillingAPIService` | Meter-side billing primitives |
| `Meter.Entitlements` | `*EntitlementsAPIService` | Entitlement checks |
| `Meter.Notifications` | `*NotificationsAPIService` | Webhook + notification channels |
| `Meter.Apps` | `*AppsAPIService` | App registrations |
| `Meter.Portal` | `*MeterPortalAPIService` | Portal tokens (Meter flavor) |
| `Meter.ProductCatalog` | `*ProductCatalogAPIService` | Product catalog |
| `Meter.Subjects` | `*SubjectsAPIService` | Subjects (e.g., users, accounts) |

## Quickstart — typical usage flows

### Ingest usage events

```go
ctx := context.Background()

events := []moolabs.Event{
    {
        Id:      "evt_unique_id",
        Type:    "api.request",
        Subject: "user_42",
        Time:    time.Now().UTC(),
        Data:    map[string]interface{}{"endpoint": "/v1/predict", "tokens": 1500},
    },
}

_, err := client.Meter.Events.IngestEvents(ctx).Body(events).Execute()
if err != nil {
    log.Fatalf("ingest: %v", err)
}
```

### Check entitlement

```go
result, _, err := client.Meter.Entitlements.CheckEntitlement(ctx).
    Subject("user_42").
    Feature("ai-credits").
    Execute()

if err != nil {
    log.Fatalf("entitlement: %v", err)
}

if result.Granted {
    // serve the request
}
```

### Create a wallet + grant credits

```go
wallet, _, err := client.Cls.Wallets.CreateWallet(ctx).
    WalletCreate(moolabs.WalletCreate{
        TenantId:  "tenant_xyz",
        PoolId:    "pool_abc",
        OwnerType: "USER",
        OwnerId:   "user_42",
    }).
    Execute()

if err != nil {
    log.Fatalf("create wallet: %v", err)
}

_, _, err = client.Cls.Grants.CreateGrant(ctx).
    GrantCreate(moolabs.GrantCreate{
        WalletId:  wallet.Id,
        Amount:    "100.00",
        ExpiresAt: time.Now().Add(365 * 24 * time.Hour),
    }).
    Execute()
```

## Error handling

The Go SDK returns errors via the standard `error` interface. Use `errors.As` to inspect typed errors from the underlying HTTP response:

```go
wallet, httpResp, err := client.Cls.Wallets.CreateWallet(ctx).Execute()
if err != nil {
    var apiErr *moolabs.GenericOpenAPIError
    if errors.As(err, &apiErr) {
        // structured error with response body + status
        log.Printf("status=%d body=%s", httpResp.StatusCode, apiErr.Body())
    }
    return err
}
```

Common status codes:

| Code | Meaning |
|---|---|
| 401 | API key invalid or expired |
| 403 | Forbidden — key valid but lacks the required scope |
| 404 | Resource not found |
| 422 | Validation error — inspect response body for field details |
| 429 | Rate limited — back off and retry |
| 5xx | Server-side error — safe to retry with exponential backoff |

## Module path note

The published Go module path is `github.com/moolabs/moolabs-go`. Import in your code as:

```go
import "github.com/moolabs/moolabs-go"

// Usage
client := moolabs.NewMoolabs(moolabs.Config{...})
```

The package name is `moolabs` even though the repo is `moolabs-go` — Go convention is to omit language suffixes from the import alias.

## Source + issues

- Source mirror: https://github.com/moolabs-hq/moolabs-go
- Issues + feature requests: https://github.com/moolabs-hq/moolabs-go/issues

This module is auto-generated from the Moolabs OpenAPI specs on every release. Direct edits to the mirror repo will be overwritten on next publish.
