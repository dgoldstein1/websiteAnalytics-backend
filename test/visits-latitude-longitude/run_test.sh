# run_test.sh

local_location="test/visits-latitude-longitude"

# init
> "${local_location}/output"

curl -s "http://localhost:${1}/visits?latitude=38.818599700927734&longitude=-77.0625" >> "${local_location}/output"


if cmp -s "${local_location}/output" "${local_location}/expected_output"
then
    exit 0
else
    exit 1
fi