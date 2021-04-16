FROM golang:latest as build
ENV APP_HOME /go/testservice

COPY ./ $APP_HOME

WORKDIR $APP_HOME

RUN CGO_ENABLED=0 GOOS=linux go build -o service_app

FROM scratch as image

COPY --from=build /go/testservice/service_app .

CMD ["/service_app"]
