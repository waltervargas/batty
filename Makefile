##
# batty
#
# @file
# @version 0.1

integration-test:
	go test -tags=integration -v

test:
	go test -v

build:
	go build cmd/batty


# end
