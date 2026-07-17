# 1. Load the official ko extension from Tilt's registry
load('ext://ko', 'ko_build')

# 2. Point Tilt to your Docker Compose file
docker_compose('docker-compose.yml')

# 3. Tell Tilt to build the images using ko instead of Docker
# The first argument MUST match the image name in docker-compose.yml
ko_build('distributed-go/test-api', './cmd/TestAPI')