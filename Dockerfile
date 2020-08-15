# build
FROM golang:1.14.4-alpine as builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build

# FROM node:12.18.2 as node
# COPY /web .
# RUN npm i
# RUN npm run build

FROM alpine:latest
copy --from=builder /app/vimlytics .
COPY /web/dist /web/dist
CMD ["./vimlytics"]

