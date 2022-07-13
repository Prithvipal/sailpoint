FROM golang:alpine

WORKDIR /bin

COPY server .

ENTRYPOINT ["server"]