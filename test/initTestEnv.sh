# injectData.sh

# Created by David Goldstein on 1.13.18
# Injects test data into server

echo "$(tput bold) --- downing containers ---$(tput sgr0)"

# down containers if neccesary
docker-compose down

echo "$(tput bold) --- building container ---$(tput sgr0)"
docker build -t dgoldstein1/websiteanalytics-backend . --quiet

echo "$(tput bold) --- loading test data ---$(tput sgr0)"
# make a temp directory of current data
# mkdir -p .temp
# cp -r docker/mongodb/data/db/* .temp 

# copy over test data into mounted docker container directory
cp -r test/data/db/* docker/mongodb/data/db

# start the dev server
echo "$(tput bold) --- starting dev server ---$(tput sgr0)"
docker-compose up -d
