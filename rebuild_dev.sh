set -euxo pipefail

go build -o main

docker build . -t dgoldstein1/websiteanalytics-backend

docker-compose up -d

docker-compose logs -f server