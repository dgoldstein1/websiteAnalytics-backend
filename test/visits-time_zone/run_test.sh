# run_test.sh

local_location="test/visits-time_zone"

# init
> "${local_location}/output"

curl -s "http://localhost:${1}/visits?time_zone=Asia/Jerusalem" >> "${local_location}/output"


if cmp -s "${local_location}/output" "${local_location}/expected_output"
then
    exit 0
else
    exit 1
fi