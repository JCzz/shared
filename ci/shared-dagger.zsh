dagger -m
dagger -m github.com/jczz/shared/ci/dagger call deploy

go build ./...

# dagger
dagger -m ./ci/dagger functions