FROM golang:1.12.3


WORKDIR $GOPATH/src/github.com/daniellockard/OperatorBot
COPY . .

ENV GO111MODULE on

RUN go get -d -v ./...

RUN go build .
RUN go build -buildmode=plugin -o echo.so plugins/echo.go

CMD ["./OperatorBot"]
