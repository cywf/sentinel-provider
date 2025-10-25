package provider

import (
	"context"

	"github.com/cywf/sentinel-provider/internal/resources"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces
var _ provider.Provider = &SentinelProvider{}

// SentinelProvider defines the provider implementation.
type SentinelProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// SentinelProviderModel describes the provider data model.
type SentinelProviderModel struct {
	Endpoint types.String `tfsdk:"endpoint"`
	APIKey   types.String `tfsdk:"api_key"`
}

// Metadata returns the provider type name.
func (p *SentinelProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "sentinel"
	resp.Version = p.version
}

// Schema defines the provider-level schema for configuration data.
func (p *SentinelProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "The Sentinel Provider allows Terraform to manage AI Sentry resources for the Sentinel Project. " +
			"This provider enables deployment and configuration of specialized AI security sentries across " +
			"various critical infrastructure sectors including healthcare, energy, finance, and more.",
		Attributes: map[string]schema.Attribute{
			"endpoint": schema.StringAttribute{
				Description: "The Sentinel API endpoint URL. May also be provided via SENTINEL_ENDPOINT environment variable.",
				Optional:    true,
			},
			"api_key": schema.StringAttribute{
				Description: "The API key for authentication with the Sentinel API. May also be provided via SENTINEL_API_KEY environment variable.",
				Optional:    true,
				Sensitive:   true,
			},
		},
	}
}

// Configure prepares a Sentinel API client for data sources and resources.
func (p *SentinelProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var config SentinelProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Configuration values are now available for use in resources and data sources
	// In a production provider, you would create an API client here
	// For now, we'll pass the configuration to resources
	resp.DataSourceData = config
	resp.ResourceData = config
}

// DataSources defines the data sources implemented in the provider.
func (p *SentinelProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

// Resources defines the resources implemented in the provider.
func (p *SentinelProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		resources.NewApolloResource,
		resources.NewAresResource,
		resources.NewAthenaResource,
		resources.NewDemeterResource,
		resources.NewFenrirResource,
		resources.NewHermesResource,
		resources.NewJupiterResource,
		resources.NewLirResource,
		resources.NewLughResource,
		resources.NewMercuryResource,
		resources.NewMorriganResource,
		resources.NewOsirisResource,
		resources.NewPtahResource,
		resources.NewRaResource,
		resources.NewShivaResource,
		resources.NewSobekResource,
		resources.NewThothResource,
		resources.NewTycheResource,
	}
}

// New returns a new provider instance.
func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &SentinelProvider{
			version: version,
		}
	}
}

