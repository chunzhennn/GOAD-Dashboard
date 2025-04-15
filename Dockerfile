#syntax=docker/dockerfile:1.7-labs

FROM node:22 AS frontend-builder
WORKDIR /app/frontend
COPY ui/package*.json ./
RUN npm install
COPY ui/ ./
RUN npm run build

FROM golang:1.24-alpine AS backend-builder
WORKDIR /app/backend
COPY go.mod go.sum ./
RUN go mod download
COPY --exclude=ui . ./
RUN mkdir -p ./ui ./build
COPY --from=frontend-builder /app/frontend/dist ./ui/dist
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags="-s -w" -o build/GOAD-Dashboard 

FROM alpine:latest
WORKDIR /app
COPY --from=backend-builder /app/backend/build/GOAD-Dashboard .

ENV PORT=8080 \
    GIN_MODE=release
EXPOSE 8080

CMD ["/app/GOAD-Dashboard"]

