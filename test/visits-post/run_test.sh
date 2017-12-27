# run_test.sh

local_location="test/visits-post"

# init
> "${local_location}/output"

# post a new visit
curl -s -H \
 "Content-Type: application/json" \
 -X POST -d '{"ipAddress": "123.456.789.1", "location": "TEST_LOCATION"}' \
 "http://localhost:${1}/visits" >> "${local_location}/output"

if cmp -s "${local_location}/output" "${local_location}/expected_output"
then
    exit 0
else
    exit 1
fi