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
go test ./test -v -run TestConvertFileVersion7
go test ./test -v -run TestConvertFileVersion8
```

#### Testes do executável

```
./bin/atlas_to_iramuteq.exe C:/Teste/atlas_to_iramuteq/questionario_version7.txt C:/Teste/atlas_to_iramuteq/escala.csv C:/Teste/atlas_to_iramuteq/resultado_version7.txt ;
./bin/atlas_to_iramuteq.exe C:/Teste/atlas_to_iramuteq/questionario_version8.txt C:/Teste/atlas_to_iramuteq/escala.csv C:/Teste/atlas_to_iramuteq/resultado_version8.txt ;
```

### Build

```
go build -o ./bin/atlas_to_iramuteq.exe ./*.go
```