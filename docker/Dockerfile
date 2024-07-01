FROM golang:1.22.4 as build

WORKDIR /usr/src/app/build

ENV GOPRIVATE=github.com/materials-resources/microservices-proto
ARG GH_ACCESS_TOKEN
RUN git config --global url.https://$GH_ACCESS_TOKEN@github.com/.insteadOf https://github.com/

COPY go.mod ./
COPY go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /usr/local/bin/prophet .


FROM scratch as production
EXPOSE 50058
COPY --from=build /usr/local/bin/prophet /usr/local/bin/prophet

ENTRYPOINT ["prophet"]
CMD ["serve"]