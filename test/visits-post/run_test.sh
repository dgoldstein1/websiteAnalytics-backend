# run_test.sh

local_location="test/visits-post"

# init
> "${local_location}/output"

# post a new visit
curl -s -H \
 "Content-Type: application/json" \
 -X POST -d '{ "ip" : "TEST_IP_1", "city" : "TEST_CITY_1", "country_code" : "US", "country_name" : "United States", "latitude" : 44.854698181152344, "longitude" : -93.785400390625, "metro_code" : 613, "region_code" : "MN", "time_zone" : "America/Chicago", "zip_code" : "55387"}' \
 "http://localhost:${1}/visits" >> "${local_location}/output"

if cmp -s "${local_location}/output" "${local_location}/expected_output"
then
    exit 0
else
    exit 1
fi