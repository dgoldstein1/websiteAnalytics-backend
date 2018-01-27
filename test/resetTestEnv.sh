# resets server to environment before test

echo "$(tput bold) --- resetting environment ---$(tput sgr0)"
docker-compose down
# reset data from before tests
rm -r docker/mongodb/data/db
mkdir -p docker/mongodb/data/db
# TODO copy over temp data into docker/mongodb/data/db and be able to restart containers
# rm -rf .temp 
