package vpn

type VPNProvider string

// VPNPending defines the Data of CloudEvent with type=dev.knative.vpn.pending
type VPNPending struct {

	// Provider holds the provider name of VPNs
	Provider VPNProvider `json:"vpn_provider"`
}

// VPNReady defines the Data of CloudEvent with type=dev.knative.vpn.ready
type VPNReady struct {

	// Provider holds the provider name of VPNs
	Provider VPNProvider `json:"vpn_provider"`
}
