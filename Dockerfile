FROM golang:1.17 AS build

WORKDIR /app
COPY . .
RUN make build

EXPOSE 80
CMD ["./build/app", "serve", "--port", "80"]