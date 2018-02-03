# run_tests.sh

# this script runs all integration tests
# before it can be run, you should run test/launchTestServer.sh from the root directory

# usage : run_tests {port}

# logs success / failure of last command
# param name of command
log_success_or_failure() {
    if [ $? -eq 0 ]
    then
        echo "$(tput setab 2 )--- SUCCESS ---$(tput sgr0) ${1} "
    else
        echo "$(tput setab 1 )--- FAILURE ---$(tput sgr0) ${1} "
    fi
}

# add authorization to all test scripts
chmod +x test/initTestEnv.sh
chmod +x test/visits/run_test.sh
chmod +x test/visits-post/run_test.sh
chmod +x test/visits-ip/run_test.sh
chmod +x test/visits-city/run_test.sh
chmod +x test/visits-country_code/run_test.sh
chmod +x test/visits-country_name/run_test.sh
chmod +x test/visits-latitude-longitude/run_test.sh
chmod +x test/visits-metro_code/run_test.sh
chmod +x test/visits-region_code/run_test.sh
chmod +x test/visits-time_zone/run_test.sh
chmod +x test/visits-zip_code/run_test.sh
chmod +x test/resetTestEnv.sh

# inject test data
test/initTestEnv.sh

# give the server and mongo a second to load
echo "$(tput bold) --- waiting for dev server to load ---$(tput sgr0)"
sleep 5

echo "$(tput bold) --- running tests ---$(tput sgr0)"
# testing queries to / and /visits
test/visits/run_test.sh ${1}
log_success_or_failure "retrieve all visits"

# get by ip
test/visits-ip/run_test.sh ${1}
log_success_or_failure "get visit by ip"
# get by city
test/visits-city/run_test.sh ${1}
log_success_or_failure "get visit by city"
# get by country_code
test/visits-country_code/run_test.sh ${1}
log_success_or_failure "get visit by country_code"
# get by country_name
test/visits-country_name/run_test.sh ${1}
log_success_or_failure "get visit by country_name"
# get by latitude-longitude
test/visits-latitude-longitude/run_test.sh ${1}
log_success_or_failure "get visit by latitude-longitude"
# get by metro_code
test/visits-metro_code/run_test.sh ${1}
log_success_or_failure "get visit by metro_code"
# get by region_code
test/visits-region_code/run_test.sh ${1}
log_success_or_failure "get visit by region_code"
# get by time_zone
test/visits-time_zone/run_test.sh ${1}
log_success_or_failure "get visit by time_zone"
# get by zip_code
test/visits-zip_code/run_test.sh ${1}
log_success_or_failure "get visit by zip_code"

# post a new visit
test/visits-post/run_test.sh ${1}
log_success_or_failure "add a new visit"


test/resetTestEnv.sh