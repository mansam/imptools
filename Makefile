.PHONY: all sav unpac

all: sav repac unpac

sav:
	go build -o ./bin/sav sav/cmd/main.go

repac:
	go build -o ./bin/repac pac/repac/main.go

unpac:
	go build -o ./bin/unpac pac/unpac/main.go