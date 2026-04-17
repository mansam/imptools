.PHONY: all exe sav repac unpac fleetwriter viewer

all: exe sav repac unpac fleetwriter viewer

exe:
	go build -o ./bin/exe sav/exe/main.go

sav:
	go build -o ./bin/sav sav/cmd/main.go

repac:
	go build -o ./bin/repac pac/repac/main.go

unpac:
	go build -o ./bin/unpac pac/unpac/main.go

fleetwriter:
	go build -o ./bin/fleetwriter sav/writer/main.go

viewer:
	go build -o ./bin/viewer viewer/cmd/main.go
