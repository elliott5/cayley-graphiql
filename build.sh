cd static
gopherjs build -m static.go
cd ..
go generate
go run example/*.go
