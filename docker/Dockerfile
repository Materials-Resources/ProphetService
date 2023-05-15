FROM golang:1.20.4-bullseye as build

WORKDIR /usr/src/app/build

COPY go.mod ./
COPY go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go build -v -o /usr/local/bin/app .


FROM alpine as production
EXPOSE 50058
COPY --from=build /usr/local/bin/app /usr/local/bin/app

CMD ["app"]