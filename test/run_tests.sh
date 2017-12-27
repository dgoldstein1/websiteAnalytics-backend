# run_tests.sh

# this script runs all integration tests
# before it can be run, you should run test/launchTestServer.sh from the root directory

# usage : run_tests {port}

# logs success / failure of last command
# param name of command
log_success_or_failure() {
    if [ $? -eq 0 ]
    then
        echo "--- SUCCESS --- ${1}"
    else
        echo "--- FAILURE --- ${1}"
    fi
}
# add authorization to all test scripts
chmod +x test/visits/run_test.sh
chmod +x test/visits-post/run_test.sh

# testing queries to / and /visits
test/visits/run_test.sh ${1}
log_success_or_failure "retrieve all visits"

# post a new visit
test/visits-post/run_test.sh ${1}
log_success_or_failure "add a new visit"
