package resources

import (
"context"
"fmt"
"time"

"github.com/hashicorp/terraform-plugin-framework/path"
"github.com/hashicorp/terraform-plugin-framework/resource"
"github.com/hashicorp/terraform-plugin-framework/types"
"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure the implementation satisfies the expected interfaces.
var (
_ resource.Resource                = &LughResource{}
_ resource.ResourceWithConfigure   = &LughResource{}
_ resource.ResourceWithImportState = &LughResource{}
)

// NewLughResource is a helper function to simplify the provider implementation.
func NewLughResource() resource.Resource {
return &LughResource{}
}

// LughResource is the resource implementation for the Lugh Sentry (Postal & Shipping sector).
type LughResource struct{}

// Metadata returns the resource type name.
func (r *LughResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
resp.TypeName = req.ProviderTypeName + "_lugh"
}

// Schema defines the schema for the resource.
func (r *LughResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
resp.Schema = GetCommonSentrySchema(
"Postal & Shipping",
"Manages a Lugh Sentry resource. Lugh is specialized for protecting the Postal & Shipping sector.",
)
}

// Configure adds the provider configured client to the resource.
func (r *LughResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
// Provider configuration is available in req.ProviderData
// In a production provider, you would configure an API client here
}

// Create creates the resource and sets the initial Terraform state.
func (r *LughResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
var plan SentryResourceModel
diags := req.Plan.Get(ctx, &plan)
resp.Diagnostics.Append(diags...)
if resp.Diagnostics.HasError() {
return
}

// Generate a unique ID for the resource
plan.ID = types.StringValue(fmt.Sprintf("lugh-%s-%d", plan.Name.ValueString(), time.Now().Unix()))
plan.Sector = types.StringValue("Postal & Shipping")
plan.Status = types.StringValue("active")
plan.LastUpdated = types.StringValue(time.Now().Format(time.RFC3339))

tflog.Info(ctx, "Creating Lugh sentry", map[string]interface{}{
"id":   plan.ID.ValueString(),
"name": plan.Name.ValueString(),
})

// In a production provider, you would make API calls to create the resource
// For this skeleton, we just set the state

diags = resp.State.Set(ctx, plan)
resp.Diagnostics.Append(diags...)
if resp.Diagnostics.HasError() {
return
}
}

// Read refreshes the Terraform state with the latest data.
func (r *LughResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
var state SentryResourceModel
diags := req.State.Get(ctx, &state)
resp.Diagnostics.Append(diags...)
if resp.Diagnostics.HasError() {
return
}

tflog.Info(ctx, "Reading Lugh sentry", map[string]interface{}{
"id": state.ID.ValueString(),
})

// In a production provider, you would read the resource from the API
// For this skeleton, we just use the existing state

diags = resp.State.Set(ctx, &state)
resp.Diagnostics.Append(diags...)
if resp.Diagnostics.HasError() {
return
}
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *LughResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
var plan SentryResourceModel
diags := req.Plan.Get(ctx, &plan)
resp.Diagnostics.Append(diags...)
if resp.Diagnostics.HasError() {
return
}

plan.LastUpdated = types.StringValue(time.Now().Format(time.RFC3339))

tflog.Info(ctx, "Updating Lugh sentry", map[string]interface{}{
"id": plan.ID.ValueString(),
})

// In a production provider, you would make API calls to update the resource

diags = resp.State.Set(ctx, plan)
resp.Diagnostics.Append(diags...)
if resp.Diagnostics.HasError() {
return
}
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *LughResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
var state SentryResourceModel
diags := req.State.Get(ctx, &state)
resp.Diagnostics.Append(diags...)
if resp.Diagnostics.HasError() {
return
}

tflog.Info(ctx, "Deleting Lugh sentry", map[string]interface{}{
"id": state.ID.ValueString(),
})

// In a production provider, you would make API calls to delete the resource
}

// ImportState imports an existing resource into Terraform.
func (r *LughResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
