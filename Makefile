CURRENT_DIR=${PWD}
BIN_DIR=${CURRENT_DIR}/bin
DISTR_MAC=darwin
DISTR_WIN=windows
ARCH=amd64
OS := $(shell uname)

all: clean build run

clean:
	@echo
	@echo Cleaning...
	rm -rf ${BIN_DIR}

build:
	@echo
	@echo Building...
ifeq ($(OS), Darwin)
	export GOOS=${DISTR_MAC}
	export GOARCH=${ARCH}
	go build -v -o ${BIN_DIR}/child ./child
	go build -v -o ${BIN_DIR}/parent ./parent
else ifeq ($(OS), Windows)
	set GOOS=${DISTR_WIN}
	set GOARCH=${ARCH}
	go build -v -o ${BIN_DIR}/child.exe ./child
	go build -v -o ${BIN_DIR}/parent.exe ./parent
else
	@echo Failed to build. OS $(OS) not supported
endif

run:
	@echo
	@echo Executing...
	${BIN_DIR}/parent