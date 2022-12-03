FROM golang:1.19

WORKDIR /app

COPY ./kp-backend/go.mod ./kp-backend/go.sum ./

RUN go mod download

COPY ./kp-backend ./

RUN go build

EXPOSE 10532

# ENTRYPOINT [ "go" ]

CMD [ "./kimchi", "-env=deploy" ]