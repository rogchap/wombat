package app

import "google.golang.org/grpc/stats"

type initData struct {
	Version   string `json:"version"`
	BuildMode string `json:"build_mode"`
}

type workspaceState struct {
	CurrentID string
}

type protos struct {
	Files []string `json:"files"`
	Roots []string `json:"roots"`
}

type options struct {
	ID      string `json:"id"`
	Addr    string `json:"addr"`
	Reflect bool   `json:"reflect"`
	Protos  protos `json:"protos"`

	Insecure   bool   `json:"insecure"`
	Plaintext  bool   `json:"plaintext"`
	Rootca     string `json:"rootca"`
	Clientcert string `json:"clientcert"`
	Clientkey  string `json:"clientkey"`
}

type methodSelect struct {
	FullName     string `json:"full_name"`
	Name         string `json:"name"`
	ClientStream bool   `json:"client_stream"`
	ServerStream bool   `json:"server_stream"`
}

type methodsSelect []methodSelect

type serviceSelect struct {
	FullName string        `json:"full_name"`
	Methods  methodsSelect `json:"methods"`
}

type servicesSelect []serviceSelect

type fieldDesc struct {
	Name     string       `json:"name"`
	FullName string       `json:"full_name"`
	Kind     string       `json:"kind"`
	Repeated bool         `json:"repeated"`
	MapKey   *fieldDesc   `json:"map_key"`
	MapValue *fieldDesc   `json:"map_value"`
	Oneof    []fieldDesc  `json:"oneof"`
	Enum     []string     `json:"enum"`
	Message  *messageDesc `json:"message"`
}

type messageDesc struct {
	Name     string      `json:"name"`
	FullName string      `json:"full_name"`
	Fields   []fieldDesc `json:"fields"`
}

type methodInput struct {
	FullName string       `json:"full_name"`
	Message  *messageDesc `json:"message"`
}

type header struct {
	Key string `json:"key"`
	Val string `json:"val"`
}

type headers []header

type rpcStart struct {
	ClientStream bool `json:"client_stream"`
	ServerStream bool `json:"server_stream"`
}

type rpcEnd struct {
	Status     string `json:"status"`
	StatusCode int32  `json:"status_code"`
	Duration   string `json:"duration"`
}

type errorMsg struct {
	Title   string `json:"title"`
	Message string `json:"msg"`
}

type rpcStatOutHeader struct {
	*stats.OutHeader
	Header string
}

type rpcStatOutPayload struct {
	*stats.OutPayload
	Data string
}

type rpcStatOutTrailer struct {
	*stats.OutTrailer
	Trailer string
}

type rpcStatInHeader struct {
	*stats.InHeader
	Header string
}

type rpcStatInPayload struct {
	*stats.InPayload
	Data string
}

type rpcStatInTrailer struct {
	*stats.InTrailer
	Trailer string
}

type rpcStatEnd struct {
	*stats.End
	Error string
}

type releaseInfo struct {
	OldVersion string `json:"old_version"`
	NewVersion string `json:"new_version"`
	URL        string `json:"url"`
}

type commands struct {
	Grpcurl string `json:"grpcurl"`
}
