#!/bin/bash
# https://github.com/protocolbuffers/protobuf/releases
# go get -u github.com/golang/protobuf/protoc-gen-go
# go get -u github.com/gogo/protobuf/protoc-gen-gogofast

protoc -I . task.proto --gogofast_out=plugins=grpc,\
Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/struct.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types:.

protoc -I . task.proto --descriptor_set_out=task.protoset --include_imports

protoc -I . task.proto --encode main.Task <<< 'text: "test", done: true' | \
protoc -I . task.proto --decode main.Task

# protoc --decode_raw < data.pbf

# hexdump -c data.pbf
