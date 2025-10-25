package resources

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SentryResourceModel describes the resource data model that is common to all sentries.
type SentryResourceModel struct {
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
	Sector      types.String `tfsdk:"sector"`
	Status      types.String `tfsdk:"status"`
	Enabled     types.Bool   `tfsdk:"enabled"`
	Config      types.Map    `tfsdk:"config"`
	Tags        types.Map    `tfsdk:"tags"`
	LastUpdated types.String `tfsdk:"last_updated"`
}
