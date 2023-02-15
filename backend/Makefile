build:
	go build

run: build
	./myb serve

watch:
	ulimit -n 1000 #increase the file watch limit, might required on MacOS
	reflex -s -r '\.go$$' go run . serve