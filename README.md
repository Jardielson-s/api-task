go install github.com/swaggo/swag/cmd/swag@latest
export PATH=$PATH:$(go env GOPATH)/bin
source ~/.bashrc  # ou source ~/.zshrc
swag init --generalInfo cmd/main.go --output docs
Gorm