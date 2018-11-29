# Atlas to Iramuteq

## Parâmetros

Lista de parâmetros a serem passados na linha de comando:

* pathAtlasFile = Caminho do arquivo com o questionário do Atlas
* pathCsvFile = Caminho do arquivo CSV com o dicionário de labels
* separatorCsv = Separador utilizado no CSV

## Comandos

### Teste

#### Testes Unitários

```
go test ./test/service -v -run TestLoad
```

### Build

```
go build -o ./bin/etl_siconv.exe ./*.go
```