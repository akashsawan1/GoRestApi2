FROM golang:1.22

WORKDIR /akash

COPY . /akash/

RUN go build /akash

EXPOSE 8080

ENTRYPOINT [ "./rest-api" ]
