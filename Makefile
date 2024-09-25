# Makefile

# Variables
IMAGE_NAME = ivancurkovic046/metrics-demo
TAG = latest

# Build the Docker image
build: build-sbom scout

# Build the Docker image
build-sbom: 
	docker buildx build --sbom=true -t $(IMAGE_NAME):$(TAG) .

# scout quickview
scout:
	docker scout quickview $(IMAGE_NAME):$(TAG)

# Push the Docker image to Docker Hub
push: build
	docker push $(IMAGE_NAME):$(TAG)

# Clean up local Docker images
clean:
	docker rmi $(IMAGE_NAME):$(TAG)

# Run tests
test:
	go test -v

# Build, push, and test[]
all: build push test