# run_test.sh

local_location="test/visits-ip"

# init
> "${local_location}/output"

# post a new visit
TESTING_IP="71.55.154.101"
curl -s "http://localhost:${1}/visits/ip/${TESTING_IP}" >> "${local_location}/output"
curl -s "http://localhost:${1}/visits/ip/" >> "${local_location}/output"
curl -s "http://localhost:${1}/visits/ip/badIpAddress" >> "${local_location}/output"


if cmp -s "${local_location}/output" "${local_location}/expected_output"
then
    exit 0
else
    exit 1
fi