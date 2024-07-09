FROM golang:1.22-alpine AS build

WORKDIR /usr/src/app/build

ARG GH_ACCESS_TOKEN

RUN apk add git && \
    git config --global url.https://$GH_ACCESS_TOKEN@github.com/.insteadOf https://github.com/

COPY go.mod ./
COPY go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go build -ldflags="-w -s" -o app .


FROM scratch
EXPOSE 50058
COPY --from=build /usr/src/app/build/app /usr/local/bin/

ENTRYPOINT ["app"]