null  :=
space := $(null) #
comma := ,

PKGSLIST = tables util
COVERPKGS= $(subst $(space),$(comma),$(strip $(foreach i,$(PKGSLIST),github.com/sudachen/go-tables/$(i))))

build:
	go build ./...

run-tests:
	#ln -sf $$(pwd) go-tables
	cd tests && go test -o ../tests.test -c -covermode=atomic -coverprofile=c.out -coverpkg=../...
	./tests.test -test.v=true -test.coverprofile=c.out
	cp c.out gocov.txt
	sed -i -e 's:github.com/sudachen/go-tables/::g' c.out

run-cover:
	go tool cover -html=gocov.txt

run-cover-tests: run-tests run-cover


