# run_test.sh

# init
cd test/visits
> output

# make a call to / and /visits and log the output
curl -s "http://localhost:${1}/visits" >> output

if cmp -s "output" "expected_output"
then
    exit 0
else
    exit 1
fi