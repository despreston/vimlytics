# build Go binary
FROM golang:1.14.4-alpine as builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build
RUN ls

# build UI
FROM node:12.18.3-alpine as node
COPY /web /web
WORKDIR /web
RUN npm i
RUN npm run build

# run binary
FROM alpine:latest
COPY --from=builder /app/vimlytics .
COPY --from=node /web/dist /web/dist
CMD ["./vimlytics"]

