go test -c  ---- создаст тестовый двоичный файл под названием packageName.test
go test strings -count=1 отключение кэширования
export MYENV=BAR && go test
go test ./... -cover
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
go test ./... -coverprofile profile
