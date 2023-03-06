FROM golang:alpine as development
# Add a work directory
WORKDIR $GOPATH/src/go_fiber
# Cache and install dependencies
COPY go.mod go.sum ./
RUN go mod download
# Copy app files
COPY . .
# Install Reflex for development (For hot reload)
RUN go install github.com/cespare/reflex@latest
# Expose port
EXPOSE 8000
# Start app
CMD reflex -g '*.go' go run main.go --start-service
