dagger -m
dagger -m github.com/jczz/shared/ci/dagger call \
  shared-index-html contents

dagger -m ./ci/dagger call shared-index-html contents

go build ./...

# dagger
dagger -m ./ci/dagger functions