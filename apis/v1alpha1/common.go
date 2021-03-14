package v1alpha1

// Profile is common required information to extend provider-alibaba
type Profile struct {
	// ZoneID is Zone ID of a region
	ZoneID string `json:"zoneID,omitempty"`
	// EnvType TODO(zzxwill) needs further confirmation
	EnvType string `json:"envType,omitempty"`
	// ServiceAccount is service account name
	ServiceAccount string `json:"serviceAccount,omitempty"`
	// ResoureGroup is resource group
	ResourceGroup string `json:"resourceGroup,omitempty"`
	// VersionID TODO(zzxwill) needs further confirmation
	VersionID int `json:"versionID,omitempty"`
}
