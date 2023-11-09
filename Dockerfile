FROM golang:1.20.4-alpine
WORKDIR /App
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /gathering-app
EXPOSE 1212
CMD ["/gathering-app", "server"]