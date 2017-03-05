FROM golang:latest


# Set GOPATH/GOROOT environment variables
RUN mkdir -p /go
ENV GOPATH /go
ENV PATH $GOPATH/bin:$PATH

# Set up app
ADD . $GOPATH/src/api
WORKDIR $GOPATH/src/api

# install glide
RUN curl https://glide.sh/get | sh
RUN glide install

# go get all of the dependencies
# RUN go get github.com/labstack/echo
# RUN go get github.com/labstack/echo/middleware
# RUN go get github.com/Sirupsen/logrus
# RUN go get github.com/go-sql-driver/mysql
# RUN go get github.com/gocraft/dbr

# Set Env
# see https://medium.com/statuscode/golang-docker-for-development-and-production-ce3ad4e69673#.y0afkspxv
# ARG echo_env
# ENV ECHO_ENV $echo_env

EXPOSE 3000

# CMD ["go", "run", "/app/main.go"]
