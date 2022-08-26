VERSION = v1.3.0
macos:
	wails build
	cd build/bin && tar -czvf hamster-client-$(VERSION)-darwin-amd64.tar.gz ./hamster-client.app

linux:
	wails build
	cd build/bin && tar -czvf hamster-client-$(VERSION)-linux-amd64.tar.gz ./hamster-client

windows:
	wails build
	cd build/bin && tar -czvf hamster-client-$(VERSION)-windows-amd64.tar.gz ./hamster-client.exe
