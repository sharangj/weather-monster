FROM golang

RUN curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey | apt-key add - &&
ADD . /go/src/github.com/sharangj/weather_monster

WORKDIR /go/src/github.com/sharangj/weather_monster

RUN go install github.com/sharangj/weather_monster

EXPOSE 8080

ENTRYPOINT /go/bin/weather_monster
