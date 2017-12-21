echo "GET /VISTS"
curl http://localhost:5000/visits
echo ""

echo "POST /VISTS"
curl -H \
 "Content-Type: application/json" \
 -X POST -d '{"ipAddress": "127.0.0.1", "location": "minneapolis"}' \
 http://localhost:5000/visits
echo ""


echo "GET /VISTS"
curl http://localhost:5000/visits
echo ""
