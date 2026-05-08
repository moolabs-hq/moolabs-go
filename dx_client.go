// Unified Moolabs SDK facade for Go.
//
// This file is hand-written DX layer that sits ON TOP of the openapi-generator
// output. It is copied into the generated tree by `generate.sh` post-codegen
// when `moolabs-go.yaml` sets `dx_dir: sdks/dx/go`.
//
// Two top-level groups customers see:
//   - client.Cls.*   → routes to api.moolabs.com (BFF, direct)
//   - client.Meter.* → routes to meter.moolabs.com (Meter, direct)
//
// Token: ONE APIKey, generated in the customer UI, valid against both backends
// (each backend validates the same token independently — no proxying).
//
// Usage:
//
//	import "github.com/moolabs/moolabs-go"
//
//	client := moolabs.NewMoolabs(moolabs.Config{APIKey: "moo_live_..."})
//
//	// CLS (BFF-routed)
//	wallet, _, err := client.Cls.Wallets.CreateWallet(ctx).WalletCreate(in).Execute()
//	grants, _, err := client.Cls.Grants.ListGrants(ctx).Execute()
//
//	// Meter (Meter-routed)
//	_, err = client.Meter.Events.IngestEvents(ctx).Body(events).Execute()
//	meters, _, err := client.Meter.Meters.ListMeters(ctx).Execute()
//
// Collision note: BFF and Meter both have Portal / Subscriptions tags. The
// stitch step renames Meter's side to MeterPortal / MeterSubscriptions /
// MeterBilling so openapi-generator emits separate Go classes. Each backend
// gets its own *APIClient, and each *Cls / *Meter struct holds pointers to
// the right service from the right APIClient.

package moolabs

const (
	DefaultClsBaseURL   = "https://api.moolabs.com"
	DefaultMeterBaseURL = "https://meter.moolabs.com"
)

// Config configures a Moolabs client.
type Config struct {
	// APIKey is the customer's UI-issued Moolabs API key. The same key
	// authenticates against both CLS (BFF) and Meter backends.
	APIKey string

	// ClsBaseURL overrides the default https://api.moolabs.com.
	// Leave empty for production; set for staging/private deployments.
	ClsBaseURL string

	// MeterBaseURL overrides the default https://meter.moolabs.com.
	MeterBaseURL string
}

// Moolabs is the top-level SDK client. Use NewMoolabs to construct.
//
// Field types are named ClsAPI / MeterAPI (suffixed) to avoid collision
// with the generated Meter resource model that also lives in this package.
// The exposed field NAMES are still Cls / Meter for clean customer code.
type Moolabs struct {
	// Cls groups all operations that route to the CLS / BFF backend.
	Cls *ClsAPI
	// Meter groups all operations that route to the Meter backend.
	Meter *MeterAPI
}

// Cls holds pointers to BFF-routed API services.
type ClsAPI struct {
	Wallets       *WalletsAPIService
	Grants        *GrantsAPIService
	Ledger        *LedgerAPIService
	Alerts        *AlertsAPIService
	AutoTopup     *AutoTopupAPIService
	RateCards     *RateCardsAPIService
	Rating        *RatingAPIService
	FxRates       *FxRatesAPIService
	Portal        *PortalAPIService        // BFF /v1/portal/*
	Subscriptions *SubscriptionsAPIService // BFF /v1/subscriptions/*
}

// Meter holds pointers to Meter-routed API services. Tags renamed at stitch
// time from Portal/Subscriptions/Billing to MeterPortal/MeterSubscriptions/
// MeterBilling so the generator emits distinct Go classes.
type MeterAPI struct {
	Events         *EventsAPIService
	Meters         *MetersAPIService
	Customers      *CustomersAPIService
	Subscriptions  *MeterSubscriptionsAPIService // /api/v1/subscriptions/*
	Billing        *MeterBillingAPIService       // /api/v1/billing/*
	Entitlements   *EntitlementsAPIService
	Notifications  *NotificationsAPIService
	Apps           *AppsAPIService
	Portal         *MeterPortalAPIService // /api/v1/portal/*
	ProductCatalog *ProductCatalogAPIService
	Subjects       *SubjectsAPIService
}

// NewMoolabs constructs a Moolabs client backed by two underlying APIClient
// instances (one per backend host). Both carry the same APIKey as a Bearer
// token in the Authorization header.
//
// Panics if cfg.APIKey is empty.
func NewMoolabs(cfg Config) *Moolabs {
	if cfg.APIKey == "" {
		panic("moolabs: Config.APIKey is required")
	}
	if cfg.ClsBaseURL == "" {
		cfg.ClsBaseURL = DefaultClsBaseURL
	}
	if cfg.MeterBaseURL == "" {
		cfg.MeterBaseURL = DefaultMeterBaseURL
	}

	clsCfg := NewConfiguration()
	clsCfg.Servers = ServerConfigurations{{URL: cfg.ClsBaseURL}}
	clsCfg.AddDefaultHeader("Authorization", "Bearer "+cfg.APIKey)
	clsClient := NewAPIClient(clsCfg)

	meterCfg := NewConfiguration()
	meterCfg.Servers = ServerConfigurations{{URL: cfg.MeterBaseURL}}
	meterCfg.AddDefaultHeader("Authorization", "Bearer "+cfg.APIKey)
	meterClient := NewAPIClient(meterCfg)

	return &Moolabs{
		Cls: &ClsAPI{
			Wallets:       clsClient.WalletsAPI,
			Grants:        clsClient.GrantsAPI,
			Ledger:        clsClient.LedgerAPI,
			Alerts:        clsClient.AlertsAPI,
			AutoTopup:     clsClient.AutoTopupAPI,
			RateCards:     clsClient.RateCardsAPI,
			Rating:        clsClient.RatingAPI,
			FxRates:       clsClient.FxRatesAPI,
			Portal:        clsClient.PortalAPI,
			Subscriptions: clsClient.SubscriptionsAPI,
		},
		Meter: &MeterAPI{
			Events:         meterClient.EventsAPI,
			Meters:         meterClient.MetersAPI,
			Customers:      meterClient.CustomersAPI,
			Subscriptions:  meterClient.MeterSubscriptionsAPI,
			Billing:        meterClient.MeterBillingAPI,
			Entitlements:   meterClient.EntitlementsAPI,
			Notifications:  meterClient.NotificationsAPI,
			Apps:           meterClient.AppsAPI,
			Portal:         meterClient.MeterPortalAPI,
			ProductCatalog: meterClient.ProductCatalogAPI,
			Subjects:       meterClient.SubjectsAPI,
		},
	}
}
