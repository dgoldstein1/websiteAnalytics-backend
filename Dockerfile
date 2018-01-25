
# build stage
FROM golang:latest 
RUN mkdir -p /go/src/github.com/dgoldstein1/websiteAnalytics-backend 
ADD . /go/src/github.com/dgoldstein1/websiteAnalytics-backend 
WORKDIR /go/src/github.com/dgoldstein1/websiteAnalytics-backend 
RUN go build -o main . 

CMD ["/go/src/github.com/dgoldstein1/websiteAnalytics-backend/main"]

