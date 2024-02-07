.PHONY: build

build:
	sam build
	sam local invoke
