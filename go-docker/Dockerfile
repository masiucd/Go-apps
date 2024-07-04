FROM golang:1.22


WORKDIR /app


COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /go-docker-mcd


EXPOSE 3000

CMD ["/go-docker-mcd"]