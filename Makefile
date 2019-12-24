TARGET=app

all: clean build

clean:
	rm -rf $(TARGET)

build:
	go build -o $(TARGET) main.go

rpi:
