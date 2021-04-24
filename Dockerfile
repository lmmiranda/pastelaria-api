FROM golang

WORKDIR $GOPATH/pastelaria-api

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

EXPOSE 8080

CMD ["pastelaria-api"]

