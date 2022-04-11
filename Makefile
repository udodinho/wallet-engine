mock:
	mockgen -source=internals/ports/service.go -destination=internals/core/services/mock/wallet_mock.go -package=mock

test:
	go test -v ./...

run:
	go run opay-wallet-engine.go