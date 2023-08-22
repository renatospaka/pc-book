package serializer

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ProtobufToJSON converts protocol buffer message to JSON string
func ProtobufToJSON(message proto.Message) ([]byte, error) {
	options := protojson.MarshalOptions{
		UseEnumNumbers:    false,
		EmitUnpopulated:   true,
		UseProtoNames:     true,
		Indent:            "	",
		Multiline:         true,
	}

	return options.Marshal(message)
}
