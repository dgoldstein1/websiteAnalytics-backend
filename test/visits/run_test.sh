# run_test.sh

local_location="test/visits"

# init
> "${local_location}/output"

# make a call to / and /visits and log the output
curl -s "http://localhost:${1}/visits" >> "${local_location}/output"

if cmp -s "${local_location}/output" "${local_location}/expected_output"
then
    exit 0
else
    exit 1
fi