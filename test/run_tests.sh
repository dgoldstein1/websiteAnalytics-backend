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

# testing queries to / and /visits
chmod +x test/visits/run_test.sh
test/visits/run_test.sh ${1}
log_success_or_failure "retrieve all visits"
