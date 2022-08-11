FROM golang:1.18-bullseye

WORKDIR /code

COPY ./ /code/

RUN go mod download

# RUN go build -o /docker-gs-ping

# EXPOSE 8080

# CMD [ "/docker-gs-ping" ]
