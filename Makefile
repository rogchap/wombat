
.PHONY setup
setup:
	go mod vendor
	git clone https://github.com/therecipe/env_darwin_amd64_513.git vendor/github.com/therecipe/env_darwin_amd64_513
	qtsetup

moc:
	qtmoc desktop ./internal/...
