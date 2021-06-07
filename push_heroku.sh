export $(cat VERSION | grep VERSION)
docker build . -t dgoldstein1/websiteanalyitcs-backend:$VERSION

docker login --username=_ --password=$(heroku auth:token) registry.heroku.com

docker tag \
 	dgoldstein1/websiteanalyitcs-backend:$VERSION \
 	registry.heroku.com/quiet-brushlands-26130/web

docker push registry.heroku.com/quiet-brushlands-26130/web

heroku container:release web --app quiet-brushlands-26130

heroku open -a quiet-brushlands-26130
heroku logs --tail -a quiet-brushlands-26130