# Atlas to Iramuteq

## Parâmetros

Lista de parâmetros a serem passados na linha de comando:

* pathAtlasFile = Caminho do arquivo com o questionário do Atlas
* pathCsvFile = Caminho do arquivo CSV com o dicionário de labels
* pathResultFile = Caminho do arquivo com o resultado do processamento
* separatorCsv = Separador utilizado no CSV

## Comandos

### Teste

#### Testes Unitários

```
go test ./test -v -run TestConvertFile
```

### Build

```
go build -o ./bin/atlas_to_iramuteq.exe ./*.go
```