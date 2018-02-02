# run_test.sh

local_location="test/visits-ip"

# init
> "${local_location}/output"

curl -s "http://localhost:${1}/visits?ip=100.15.234.229" >> "${local_location}/output"


if cmp -s "${local_location}/output" "${local_location}/expected_output"
then
    exit 0
else
    exit 1
fi