# Atlas to Iramuteq
Programa para a geração de arquivo para o programa Iramuteq, a partir de arquivos usados em questionários do programa Atlas.

## Download

### Download da versão atual (1.0.0):
* Windows: [atlas_to_iramuteq.exe](https://github.com/vagnerpraia/atlas_to_iramuteq/releases/download/1.0.0/atlas_to_iramuteq)
* Linux: [atlas_to_iramuteq](https://github.com/vagnerpraia/atlas_to_iramuteq/releases/download/1.0.0/atlas_to_iramuteq.exe)

## Parâmetros

Lista de parâmetros e sua respectiva ordem de utilização:

* pathQuizFile: Caminho do arquivo com o questionário do Atlas (default: questionario.txt)
* pathCsvFile: Caminho do arquivo CSV com o dicionário de labels (default: escala.csv)
* pathResultFile: Caminho do arquivo com o resultado do processamento (default: resultado.txt)
* separatorCsv: Separador utilizado no CSV (default: ;)
* quoteCsv: Indicador de fronteira dos dados utilizado no CSV (default: ")

## Comandos

### Testes

#### Testes unitários

```
go test ./test/control -v -run TestRouterVersion7
go test ./test/control -v -run TestRouterVersion78
```

#### Testes do executável

```
bin/atlas_to_iramuteq.exe C:/Teste/atlas_to_iramuteq/questionario_version7.txt C:/Teste/atlas_to_iramuteq/escala.csv C:/Teste/atlas_to_iramuteq/resultado_version7.txt "," """
bin/atlas_to_iramuteq.exe C:/Teste/atlas_to_iramuteq/questionario_version8.txt C:/Teste/atlas_to_iramuteq/escala.csv C:/Teste/atlas_to_iramuteq/resultado_version8.txt "," """
```

## Build

Com debug ativado:
```
GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -a -ldflags '-extldflags "-static"' -o bin/atlas_to_iramuteq.exe *.go
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -ldflags '-extldflags "-static"' -o bin/atlas_to_iramuteq *.go
```

Com debug desativado:
```
GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -a -ldflags '-extldflags "-static" -s -w' -o bin/atlas_to_iramuteq.exe *.go
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -ldflags '-extldflags "-static" -s -w' -o bin/atlas_to_iramuteq *.go
```
