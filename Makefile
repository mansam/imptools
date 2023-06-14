.PHONY: all sav unpac

all: sav unpac

sav:
	go build -o ./sav sav/cmd/main.go

unpac:
	go build -o ./unpac unpac/cmd/main.go