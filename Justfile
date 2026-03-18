set dotenv-load := true

out := "build/goaler.bot"
version := `git rev-parse --short HEAD`
flags := "-ldflags=\"-X 'main.version=" + version + "' -w -s\""

all: build

[group('run')]
buildx: build exec

[group('build')]
build:
    go build -o {{out}} .

[group('build')]
buildf:
    @echo "Building with flags"
    go build {{flags}} -o {{out}} .

[group('run')]
exec:
    @echo "Executing binary:"
    ./{{out}}

[group('run')]
run:
    go run .

clean:
    rm {{out}}

[group('test')]
test *args:
    go test -v . {{args}}

[group('test')]
check:
    staticcheck .
