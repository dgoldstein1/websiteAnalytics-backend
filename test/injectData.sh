# injectData.sh

# Created by David Goldstein on 1.13.18
# Injects test data into server

echo "loading test data..."

# post a new visit
curl -s -H \
 "Content-Type: application/json" \
 -X POST -d '{ "ip" : "TEST_IP_A", "city" : "TEST_CITY_A", "country_code" : "US", "country_name" : "United States", "latitude" : 44.854698181152344, "longitude" : -93.785400390625, "metro_code" : 613, "region_code" : "MN", "time_zone" : "America/Chicago", "zip_code" : "99999"}' \
 "http://localhost:${1}/visits"

 # post a new visit
curl -s -H \
 "Content-Type: application/json" \
 -X POST -d '{ "ip" : "TEST_IP_B", "city" : "TEST_CITY_B", "country_code" : "US", "country_name" : "United States", "latitude" : 44.854698181152344, "longitude" : -93.785400390625, "metro_code" : 613, "region_code" : "MN", "time_zone" : "America/Chicago", "zip_code" : "11111"}' \
 "http://localhost:${1}/visits"


echo "done"