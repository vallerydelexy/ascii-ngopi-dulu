GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o build/ascii-darwin-arm64 ascii.go
GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o build/ascii-darwin-amd64 ascii.go
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o build/ascii-linux-amd64 ascii.go
GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o build/ascii-windows-amd64.exe ascii.go

# Compress with UPX
upx --best --ultra-brute build/ascii-darwin-arm64
upx --best --ultra-brute build/ascii-darwin-amd64
upx --best --ultra-brute build/ascii-linux-amd64
upx --best --ultra-brute build/ascii-windows-amd64.exe
