build:
	packr2.exe build
	go build -tags=jsoniter .
swagger:
	swag init