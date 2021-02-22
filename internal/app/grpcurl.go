package app

import (
	"flag"
	"fmt"
	"strings"

	"github.com/google/shlex"
)

type multiString []string

func (s *multiString) String() string {
	return strings.Join(*s, ",")
}

func (s *multiString) Set(value string) error {
	*s = append(*s, value)
	return nil
}

// GrpcurlArguments parsed grpcurl command arguments
type GrpcurlArguments struct {
	Target string `json:"target"`
	Method string `json:"method"`
	Metadata headers `json:"metadata"`
	Data string `json:"data"`
}

// ParseGrpcurlCommand parse grpcurl command
func ParseGrpcurlCommand(command string) (*GrpcurlArguments, error) {
	args, _ := shlex.Split(command)
	for index, arg := range args {
		if strings.TrimSpace(arg) == "" {
			args = append(args[:index], args[index+1:]...)
		}
	}

	if len(args) == 0 {
		return nil, fmt.Errorf("Empty grpcurl command. ")
	}

	if strings.ToLower(args[0]) != "grpcurl" {
		return nil, fmt.Errorf("Invalid grpcurl command. ")
	}

	flags := flag.NewFlagSet(args[0], flag.ExitOnError)

	_ = flags.Bool("help", false, "")
	_ = flags.Bool("version", false,"")
	_ = flags.Bool("plaintext", false, "")
	_ = flags.Bool("insecure", false, "")
	_ = flags.String("cacert", "", "")
	_ = flags.String("cert", "", "")
	_ = flags.String("key", "", "")
	_ = flags.Bool("expand-headers", false, "")
	_ = flags.String("authority", "", "")
	_ = flags.String("user-agent", "", "")
	data := flags.String("d", "", "")
	format := flags.String("format", "json", "")
	_ = flags.Bool("allow-unknown-fields", false, "")
	_ = flags.Float64("connect-timeout", 0, "")
	_ = flags.Bool("format-error", false, "")
	_ = flags.Float64("keepalive-time", 0, "")
	_ = flags.Float64("max-time", 0, "")
	_ = flags.Int("max-msg-sz", 0, "")
	_ = flags.Bool("emit-defaults", false, "")
	_ = flags.String("protoset-out", "", "")
	_ = flags.Bool("msg-template", false, "")
	_ = flags.Bool("v", false, "")
	_ = flags.Bool("vv", false, "")
	_ = flags.String("servername", "", "")
	_ = flags.Bool("use-reflection", false, "")

	if *format != "" && *format != "json" {
		return nil, fmt.Errorf("Data format must be json. ")
	}

	var protoset      multiString
	var protoFiles    multiString
	var importPaths   multiString
	var addlHeaders   multiString
	var rpcHeaders    multiString
	var reflHeaders   multiString
	flags.Var(&addlHeaders, "H", "")
	flags.Var(&rpcHeaders, "rpc-header", "")
	flags.Var(&reflHeaders, "reflect-header", "")
	flags.Var(&protoset, "protoset", "")
	flags.Var(&protoFiles, "proto", "")
	flags.Var(&importPaths, "import-path", "")

	err := flags.Parse(args[1:])
	if err != nil {
		return nil, err
	}

	grpcurlArgs := flags.Args()
	if len(grpcurlArgs) != 2 {
		return nil, fmt.Errorf("Invalid grpcurl arguments. ")
	}

	var metadata headers

	for _, headerValue := range append(addlHeaders, rpcHeaders...) {
		headerData := strings.Split(headerValue, ":")
		
		if len(headerData) != 2 {
			continue
		}
		metadata = append(metadata, header {
			Key: headerData[0],
			Val: headerData[1],
		})
	}

	return &GrpcurlArguments {
		Target: grpcurlArgs[0],
		Method: grpcurlArgs[1],
		Data: *data,
		Metadata: metadata,
	}, nil
}