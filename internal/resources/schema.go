package resources

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

// GetCommonSentrySchema returns the common schema attributes for all sentry resources
func GetCommonSentrySchema(sectorName, description string) schema.Schema {
	return schema.Schema{
		Description: description,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "The unique identifier for this sentry resource.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Description: "The name of the sentry instance.",
				Required:    true,
			},
			"description": schema.StringAttribute{
				Description: "A description of the sentry instance and its purpose.",
				Optional:    true,
			},
			"sector": schema.StringAttribute{
				Description: "The critical infrastructure sector this sentry protects.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"status": schema.StringAttribute{
				Description: "The current operational status of the sentry (e.g., active, inactive, maintenance).",
				Computed:    true,
			},
			"enabled": schema.BoolAttribute{
				Description: "Whether the sentry is enabled and actively monitoring.",
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(true),
			},
			"config": schema.MapAttribute{
				Description: "Configuration parameters specific to this sentry.",
				ElementType: schema.StringAttribute{}.GetType(),
				Optional:    true,
			},
			"tags": schema.MapAttribute{
				Description: "A map of tags to assign to the sentry resource.",
				ElementType: schema.StringAttribute{}.GetType(),
				Optional:    true,
			},
			"last_updated": schema.StringAttribute{
				Description: "Timestamp of the last update to this resource.",
				Computed:    true,
			},
		},
	}
}
