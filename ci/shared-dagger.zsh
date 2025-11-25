go build ./...

dagger -m ./ci/dagger functions
dagger -m github.com/jczz/shared/ci/dagger call deploy