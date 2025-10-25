package provider

import (
"testing"

"github.com/hashicorp/terraform-plugin-framework/providerserver"
"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

// testAccProtoV6ProviderFactories are used to instantiate a provider during
// acceptance testing. The factory function will be invoked for every Terraform
// CLI command executed to create a provider server to which the CLI can
// reattach.
var testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
"sentinel": providerserver.NewProtocol6WithError(New("test")()),
}

func TestProvider(t *testing.T) {
// This is a basic test to ensure the provider can be instantiated
provider := New("test")()
if provider == nil {
t.Fatal("Expected provider to be non-nil")
}
}

func TestProviderMetadata(t *testing.T) {
provider := New("test")()
sentinelProvider, ok := provider.(*SentinelProvider)
if !ok {
t.Fatal("Expected provider to be of type *SentinelProvider")
}

if sentinelProvider.version != "test" {
t.Errorf("Expected version to be 'test', got '%s'", sentinelProvider.version)
}
}
