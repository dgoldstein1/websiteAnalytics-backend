# url=https://salty-sierra-82911.herokuapp.com/visits
url=http://localhost:5000/visits

echo "GET /VISTS"
curl "${url}"
echo ""

echo "POST /VISTS"
curl -H \
 "Content-Type: application/json" \
 -X POST -d '{"ipAddress": "127.0.0.1", "location": "minneapolis"}' \
 "${url}"
echo ""


echo "GET /VISTS"
curl "${url}"
echo ""
