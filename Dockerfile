FROM golang:1.21-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN GOOS=linux go build -a -o youtube-studio-v2 .

FROM scratch
WORKDIR /root/
COPY --from=build /app/youtube-studio-v2 .
CMD ["./youtube-studio-v2"]
