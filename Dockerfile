FROM golang:alpine
WORKDIR /usr/app
COPY . .
RUN go build -o main -ldflags "-s -w" .
CMD [ "./main" ]