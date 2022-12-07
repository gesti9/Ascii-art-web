FROM golang:alpine

WORKDIR /app
COPY . . 





RUN go build -o main .

LABEL version="0.1"
LABEL name="gesti9"
EXPOSE 8080


CMD ["./main"]