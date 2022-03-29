build:
	buildctl build \
		--frontend dockerfile.v0 \
		--local context=. \
		--local dockerfile=. \
		--opt network=bridge
		--output type=image,name=docker.io/username/image,push=false