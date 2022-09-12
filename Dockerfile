FROM golang:latest as builder
COPY ./backend/vendor ./backend/go.mod ./backend/go.sum /server
WORKDIR /server
COPY ./backend .
RUN go build ./cmd/main.go

FROM node:alpine as node_builder
COPY ./frontend/ /frontend
WORKDIR /frontend
RUN npm install && npm run build

FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN apk --no-cache add libc6-compat
WORKDIR /app
COPY --from=builder /server/main /server/config ./
COPY --from=node_builder /frontend/dist ./frontend
RUN chmod +x ./main
EXPOSE 8081
CMD ./main
