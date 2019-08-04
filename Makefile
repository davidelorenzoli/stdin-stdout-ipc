build:
	go build -o ./bin/child ./child
	go build -o ./bin/parent ./parent

clean:
	rm -r ./bin

run:
	./bin/parent