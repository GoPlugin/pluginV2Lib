.PHONY: gomodtidy
gomodtidy:
	go mod tidy
	cd ./ops && go mod tidy

.PHONY: godoc
godoc:
	go install golang.org/x/tools/cmd/godoc@latest
	# http://localhost:6060/pkg/github.com/Goplugin/pluginV2-relay/
	godoc -http=:6060
