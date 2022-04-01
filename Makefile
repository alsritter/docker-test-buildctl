build:
	go build -o dctest 
	buildctl --debug \
		--addr tcp://localhost:1234 build \
		--frontend dockerfile.v0 \
		--local context=. \
		--local dockerfile=. \
		--output type=image,name=alsritter.com/zhujl04/ci-test:01-master,push=false