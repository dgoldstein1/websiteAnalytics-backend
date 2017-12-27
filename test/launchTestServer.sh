# launchTestServer.sh

# Created by David Goldstein on 12.27.2017
# Launches a testDB 

# usage : {port}

# load in data
echo "loading in data..."
cp test/sampleData.db bolt.db
echo "done"

export TEST_PORT=${1}
echo "launching project on port ${TEST_PORT} ..."
go install
heroku local -p ${TEST_PORT}

echo "cleaning data..."
> bolt.db
echo "done"

