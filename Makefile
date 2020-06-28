
.PHONY: setup
setup:
	go mod vendor
	git clone https://github.com/therecipe/env_darwin_amd64_513.git vendor/github.com/therecipe/env_darwin_amd64_513
	qtsetup
	
mod:
	mkdir _temp
	mv vendor/github.com/therecipe/qt _temp/qt
	mv vendor/github.com/therecipe/env_darwin_amd64_513 _temp/env_darwin_amd64_513
	go mod vendor
	mv _temp/env_darwin_amd64_513 vendor/github.com/therecipe/env_darwin_amd64_513
	mv _temp/qt vendor/github.com/therecipe/qt
	rm -rf _temp

clean-moc:
	find ./internal -name 'moc*' -delete
