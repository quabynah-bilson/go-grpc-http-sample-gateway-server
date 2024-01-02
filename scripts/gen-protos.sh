# Description: Generate the protobuf files

# variables
PROTO_PATH=protos
OUR_DIR=server/proto_gen

# create the gen directory
mkdir -p "$OUR_DIR"
mkdir -p "$OUR_DIR"/openapi

# remove the old generated files
rm -rf "$OUR_DIR/*.go"

# generate the new files
protoc -I=$PROTO_PATH --go_out=$OUR_DIR --go_opt=paths=source_relative \
  --go-grpc_out=$OUR_DIR --go-grpc_opt=paths=source_relative \
   --grpc-gateway_opt=paths=source_relative \
   --grpc-gateway_out=$OUR_DIR \
   --openapiv2_out=$OUR_DIR/openapi \
  $(find $PROTO_PATH -name '*.proto')
