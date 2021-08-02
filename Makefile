#
# github.com/MichaelDarr/stack
#

.PHONY: proto
proto:
	cd $(CURDIR)
	protoc --go_out=./ --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		--js_out=import_style=commonjs:. \
		--grpc-web_out=import_style=typescript,mode=grpcwebtext:. \
		proto/**/*.proto
