FROM golang:1.25-alpine AS build

WORKDIR /src/backend

COPY backend/go.mod backend/go.sum ./
RUN go mod download

COPY backend/ ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /out/affiliate-api ./cmd/api

FROM alpine:3.22

RUN addgroup -S affiliate && adduser -S affiliate -G affiliate

WORKDIR /app
COPY --from=build /out/affiliate-api /app/affiliate-api

USER affiliate
EXPOSE 8080

CMD ["/app/affiliate-api"]
