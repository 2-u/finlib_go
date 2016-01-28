# finlib_go
Financial calculations library written in Go

# Make sure you set your GOPATH correctly
cd $GOPATH/src

# If you make changes, please run gofmt
gofmt -w github.com/mraym/finlib_go/*.go

# Build the library
go build github.com/mraym/finlib_go

# Run the test to make sure that there is output.
# You should see the mortgage payment and you should see FAIL.
# That's ok, I just want to see the mortgage payment output.
go test github.com/mraym/finlib_go
