FROM golang:1.23

WORKDIR /go/project/nbrb

ADD go.mod go.sum main.go ./
ADD internal ./internal

EXPOSE 8080

CMD ["go", "run", "main.go"]