# # FROM wingrunr21/alpine-heroku-cli
# FROM heroku/heroku:16
# MAINTAINER David Goldstein <dgoldstein01@gmail.com>

# COPY . /app/src/github.com/dgoldstein1/websiteAnalytics-backend
# WORKDIR /app/src/github.com/dgoldstein1/websiteAnalytics-backend

# ENV HOME /app
# ENV GOVERSION=1.9.2
# ENV GOROOT $HOME/.go/$GOVERSION/go
# ENV GOPATH $HOME
# ENV PATH $PATH:$HOME/bin:$GOROOT/bin:$GOPATH/bin
# RUN mkdir -p $HOME/.go/$GOVERSION
# RUN cd $HOME/.go/$GOVERSION; curl -s https://storage.googleapis.com/golang/go$GOVERSION.linux-amd64.tar.gz | tar zxf -

# RUN go install -v github.com/dgoldstein1/websiteAnalytics-backend
# RUN wget -qO- https://cli-assets.heroku.com/install-ubuntu.sh | sh

# RUN heroku --version

# build stage
FROM golang:latest 
RUN mkdir -p /go/src/github.com/dgoldstein1/websiteAnalytics-backend 
ADD . /go/src/github.com/dgoldstein1/websiteAnalytics-backend 
WORKDIR /go/src/github.com/dgoldstein1/websiteAnalytics-backend 
RUN go build -o main . 

CMD ["/go/src/github.com/dgoldstein1/websiteAnalytics-backend/main"]

