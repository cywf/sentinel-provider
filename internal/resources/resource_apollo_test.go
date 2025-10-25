package resources

import (
"testing"

"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestApolloResourceMetadata(t *testing.T) {
resource := NewApolloResource()
apolloResource, ok := resource.(*ApolloResource)
if !ok {
t.Fatal("Expected resource to be of type *ApolloResource")
}

if apolloResource == nil {
t.Fatal("Expected resource to be non-nil")
}
}

func TestSentryResourceModel(t *testing.T) {
model := SentryResourceModel{
ID:          types.StringValue("test-id"),
Name:        types.StringValue("test-sentry"),
Description: types.StringValue("test description"),
Sector:      types.StringValue("Healthcare"),
Status:      types.StringValue("active"),
Enabled:     types.BoolValue(true),
}

if model.ID.ValueString() != "test-id" {
t.Errorf("Expected ID to be 'test-id', got '%s'", model.ID.ValueString())
}

if model.Name.ValueString() != "test-sentry" {
t.Errorf("Expected Name to be 'test-sentry', got '%s'", model.Name.ValueString())
}

if model.Sector.ValueString() != "Healthcare" {
t.Errorf("Expected Sector to be 'Healthcare', got '%s'", model.Sector.ValueString())
}

if !model.Enabled.ValueBool() {
t.Error("Expected Enabled to be true")
}
}
