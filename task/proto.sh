#!/bin/bash

# go get -u github.com/gogo/protobuf/types
# go get -u github.com/gogo/protobuf/protoc-gen-gofast

protoc task.proto -I . \
--gofast_out=plugins=grpc,paths=source_relative,\
Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/struct.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types:.

# go get -u github.com/golang/protobuf/protoc-gen-go
# go get -u github.com/golang/protobuf/ptypes

# protoc task.proto -I . --go_out=plugins=grpc,paths=source_relative:.

protoc -I . task.proto --descriptor_set_out=task.protoset --include_imports
protoc -I . task.proto --encode task.Task <<< 'text: "test", done: true' > task.pbf
protoc -I . task.proto --decode task.Task < task.pbf
protoc --decode_raw < task.pbf
hexdump -c task.pbf
hexdump task.pbf

# npm i protobufjs -g
pbjs -p . task.proto -t static-module -w es6 -o task.pb.mjs
