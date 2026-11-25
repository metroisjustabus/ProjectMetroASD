# ---- Stage 1: build the React frontend ----
FROM node:20-alpine AS frontend
WORKDIR /frontend
COPY web/package.json ./package.json
RUN npm install
COPY web/ ./ 
RUN npm run build

# ---- Stage 2: build the Go backend ----
FROM golang:1.23-alpine AS backend
WORKDIR /backend
COPY services/core/go.mod services/core/go.sum* ./
RUN go mod download
COPY services/core/ ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /out/metro .

# ---- Stage 3: minimal runtime ----
FROM alpine:3.20
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=backend /out/metro ./metro
COPY --from=frontend /frontend/dist ./static

EXPOSE 5000

HEALTHCHECK --interval=30s --timeout=3s --start-period=5s \
  CMD wget -qO- http://localhost:5000/health || exit 1

CMD ["./metro"]