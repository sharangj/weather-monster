FROM golang

ADD . /go/src/github.com/sharangj/weather_monster

WORKDIR /go/src/github.com/sharangj/weather_monster

RUN go install github.com/sharangj/weather_monster

EXPOSE 8080

ENTRYPOINT /go/bin/weather_monster
