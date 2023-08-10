# Build stage
FROM golang:1.20-alpine3.17 AS build

WORKDIR /go/src/trimly
COPY . .
COPY sample.env .

RUN go mod download
RUN go build -o /go/src/trimly/main

# Final stage
FROM alpine:3.17
WORKDIR /app
COPY --from=build /go/src/trimly/main .
COPY --from=build /go/src/trimly/sample.env .

ENTRYPOINT ["./main"]