# run_test.sh

local_location="test/visits-query_type-or"

# init
> "${local_location}/output"

curl -s "http://localhost:${1}/visits?country_code=US&zip_code=22301&query_type=or" >> "${local_location}/output"


if cmp -s "${local_location}/output" "${local_location}/expected_output"
then
    exit 0
else
    exit 1
fi