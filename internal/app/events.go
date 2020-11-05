package app

const (
	eventClientConnected       = "wombat:client_connected"
	eventClientStateChanged    = "wombat:client_state_changed"
	eventServicesSelectChanged = "wombat:services_select_changed"
	eventMethodInputChanged    = "wombat:method_input_changed"
	eventRPCStarted            = "wombat:rpc_started"
	eventRPCEnded              = "wombat:rpc_ended"
	eventInPayloadReceived     = "wombat:in_payload_received"
	eventInHeaderReceived      = "wombat:in_header_received"
	eventInTrailerReceived     = "wombat:in_trailer_received"
)
