FROM golang:1.18-alpine

WORKDIR /go/src
COPY . .

RUN apk upgrade --update && \
    apk add --no-cache gcc libc-dev && \
    apk --no-cache add git

RUN go mod download

CMD ["go", "run", "github.com/cosmtrek/air"]
