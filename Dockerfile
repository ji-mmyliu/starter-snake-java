# FROM scratch
FROM ubuntu:20.04

COPY --from=golang:1.18 /usr/local/go/ /usr/local/go/
ENV PATH="/usr/local/go/bin:${PATH}"

RUN apt-get update -y && apt-get upgrade -y
# RUN apt-get install golang-go -y

RUN apt-get install default-jdk -y

WORKDIR /app
COPY . .
RUN go build -v -o ./app ./server

RUN ["javac", "Main.java"]

CMD ["./app"]
# CMD ["go", "run", "./server"]