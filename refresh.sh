CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o refresh_amd64
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o refresh_arm64

# Create a universal binary
lipo -create -output refresh refresh_amd64 refresh_arm64
