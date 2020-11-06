package app

const (
	eventInit                  = "wombat:init"
	eventError                 = "wombat:error"
	eventClientConnected       = "wombat:client_connected"
	eventClientStateChanged    = "wombat:client_state_changed"
	eventServicesSelectChanged = "wombat:services_select_changed"
	eventMethodInputChanged    = "wombat:method_input_changed"
	eventRPCStarted            = "wombat:rpc_started"
	eventRPCEnded              = "wombat:rpc_ended"
	eventInPayloadReceived     = "wombat:in_payload_received"
	eventInHeaderReceived      = "wombat:in_header_received"
	eventInTrailerReceived     = "wombat:in_trailer_received"
	eventStatBegin             = "wombat:stat_begin"
	eventStatOutHeader         = "wombat:stat_out_header"
	eventStatOutPayload        = "wombat:stat_out_payload"
	eventStatOutTrailer        = "wombat:stat_out_trailer"
	eventStatInHeader          = "wombat:stat_in_header"
	eventStatInPayload         = "wombat:stat_in_payload"
	eventStatInTrailer         = "wombat:stat_in_trailer"
	eventStatEnd               = "wombat:stat_end"
)
