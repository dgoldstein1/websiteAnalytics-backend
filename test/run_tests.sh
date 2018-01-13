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
chmod +x test/injectData.sh
chmod +x test/visits/run_test.sh
chmod +x test/visits-post/run_test.sh
chmod +x test/visits-ip/run_test.sh
chmod +x test/visits-filters/run_test.sh

# inject test data
test/injectData.sh ${1}

# wait one second for the test data to load
sleep 1

# testing queries to / and /visits
test/visits/run_test.sh ${1}
log_success_or_failure "retrieve all visits"

# post a new visit
test/visits-post/run_test.sh ${1}
log_success_or_failure "add a new visit"

# # get by ip
# test/visits-ip/run_test.sh ${1}
# log_success_or_failure "get visit by ip"

# # get by filter
# test/visits-filters/run_test.sh ${1}
# log_success_or_failure "filter by date"