# run_test.sh

local_location="test/visits-filters"

# init
> "${local_location}/output"

# get by ip address
TESTING_IP="71.55.154.101"
curl -s "http://localhost:${1}/visits?ip=${TESTING_IP}" >> "${local_location}/output"
# all visits from last week
curl -s "http://localhost:${1}/visits?ip=${TESTING_IP}&from=-7" >> "${local_location}/output"
# all visits from 7 seven days to 3 days ago
curl -s "http://localhost:${1}/visits?ip=${TESTING_IP}&from=-7&to=-3" >> "${local_location}/output"
# bad arguments
# from > to
curl -s curl -s "http://localhost:${1}/visits?from=-7&to=-10" >> "${local_location}/output"


if cmp -s "${local_location}/output" "${local_location}/expected_output"
then
    exit 0
else
    exit 1
fi