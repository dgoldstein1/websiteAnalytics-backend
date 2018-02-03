# run_test.sh

local_location="test/visits-query_type-and"

# init
> "${local_location}/output"

curl -s "http://localhost:${1}/visits?country_code=US&zip_code=22301&query_type=and" >> "${local_location}/output"
# default should be and (same output as above)
curl -s "http://localhost:${1}/visits?country_code=US&zip_code=22301" >> "${local_location}/output"
# test bad input doesn't work
curl -s "http://localhost:${1}/visits?country_code=US&zip_code=22301&query_type=ansdfd" >> "${local_location}/output"


if cmp -s "${local_location}/output" "${local_location}/expected_output"
then
    exit 0
else
    exit 1
fi