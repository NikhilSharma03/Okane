FROM golang:1.18-alpine3.15

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./

RUN cd server/cmd && go build -v -o server

EXPOSE 8000

# Run the executable
CMD ["/app/server/cmd/server"]