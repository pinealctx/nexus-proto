.PHONY: generate clean lint

# ─── Protobuf Code Generation ────────────────────────────────

generate:
	cd proto && buf generate --template buf.gen.go.yaml
	cd proto && buf generate --template buf.gen.connect.yaml

clean:
	rm -rf gen/go/api
	rm -rf gen/go/shared

lint:
	cd proto && buf lint
