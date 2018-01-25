# injectData.sh

# Created by David Goldstein on 1.13.18
# Injects test data into server

echo " --- downing containers... --- "
# down containers if neccesary
docker-compose down
echo "--- done ---"

echo "--- loading test data... ---"
# make a temp directory of current data
# mkdir -p .temp
# cp -r docker/mongodb/data/db/* .temp 

# copy over test data into mounted docker container directory
cp -r test/data/db/* docker/mongodb/data/db
echo "--- done ---"

# start the dev server
echo "--- starting dev server ---"
docker-compose up -d
echo "--- done ---"