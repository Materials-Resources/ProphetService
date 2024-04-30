FROM golang:1.22.1-alpine as build

WORKDIR /usr/src/app/build

COPY go.mod ./
COPY go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go build -v -o /usr/local/bin/prophet .


FROM alpine as production
EXPOSE 50058
COPY --from=build /usr/local/bin/prophet /usr/local/bin/prophet

ENTRYPOINT ["prophet"]
CMD ["serve"]