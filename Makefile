SRC	 := ./cmd
BIN  := ./build/bin/go-tool

DEBUG 	:= -gcflags='all=-N -l'
RELEASE := -ldflags='-w -s'

all:
	go build ${DEBUG} -o ${BIN} ${SRC}

clean:
	rm -rf ${BIN}

.PHONY: all clean

.DEFAULT_GOAL: all