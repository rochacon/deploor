get: get-test get-prod

get-test:
	@/bin/echo "Installing test dependencies... "
	@go list -f '{{range .TestImports}}{{.}} {{end}}' ./... | tr ' ' '\n' |\
		grep '^.*\..*/.*$$' | grep -v 'github.com/rochacon/git-hooks-to-run-fabric' |\
		sort | uniq | xargs go get >/dev/null 2>&1
	@/bin/echo "ok"

get-prod:
	@/bin/echo "Installing production dependencies... "
	@go list -f '{{range .Imports}}{{.}} {{end}}' ./... | tr ' ' '\n' |\
		grep '^.*\..*/.*$$' | grep -v 'github.com/rochacon/git-hooks-to-run-fabric' |\
		sort | uniq | xargs go get >/dev/null 2>&1
	@/bin/echo "ok"

test:
	@go test -i ./...
	@go test ./...

all:
	@mkdir -p dist/hooks
	@go build -o dist/hooks/post-receive hooks/post-receive.go
	@go build -o dist/hooks/update hooks/update.go
	@chmod +x dist/hooks/*

clean:
	@rm -r dist/
