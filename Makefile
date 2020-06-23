
.PHONY: setup
setup:
	go mod vendor
	git clone https://github.com/therecipe/env_darwin_amd64_513.git vendor/github.com/therecipe/env_darwin_amd64_513
	qtsetup

.PHONY: moc
moc: clean-moc
	qtmoc desktop ./internal/app
	qtmoc desktop ./internal/model


clean-moc:
	find ./internal -name 'moc*' -delete
