PREFIX=$(shell pwd)
OUTPUT_DIR=${PREFIX}/bin
OUTPUT_FILE=${OUTPUT_DIR}/course

build:
	@echo "+ build"
	go build -o ${OUTPUT_FILE}