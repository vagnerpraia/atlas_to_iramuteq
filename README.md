# Atlas to Iramuteq
Programa para a geração de arquivo para o programa Iramuteq, a partir de arquivos usados em questionários do programa Atlas.

## Download

### Versão atual (1.0.0):
* Windows: [atlas_to_iramuteq.exe](https://github.com/vagnerpraia/atlas_to_iramuteq/releases/download/1.0.0/atlas_to_iramuteq)
* Linux: [atlas_to_iramuteq](https://github.com/vagnerpraia/atlas_to_iramuteq/releases/download/1.0.0/atlas_to_iramuteq.exe)

## Utilização

A utilização do programa é por meio de linha de comando, onde o arquivo executável do programa deve ser chamado. Opcionalmente, podem ser passados parâmetros de configuração da execução do programa. Estes parâmetros devem ser passados em ordem após o nome do execuável. Quando os parâmetros não forem passados, serão utilizados valores padrões.

A lista de parâmetros, seus valores padrões e sua respectiva ordem é a seguinte:

* Parâmetro 1: pathQuiz
    * Descrição: Caminho do arquivo com o questionário do Atlas.
    * Valor padrão: questionario.txt
* Parâmetro 2: pathCsv
    * Descrição: Caminho do arquivo CSV com a escala dos dados.
    * Valor padrão: escala.csv
* Parâmetro 3: pathResult
    * Descrição: Caminho do arquivo com o resultado do processamento.
    * Valor padrão: resultado.txt
* Parâmetro 4: separatorCsv
    * Descrição: Caractere utilizado como separador de dados no CSV.
    * Valor padrão: ;
* Parâmetro 5: quoteCsv
    * Descrição: Caractere utilizado como limitador de dados no CSV.
    * Valor padrão: "

## Testes

### Testes unitários

```
go test ./test/control -v -run TestRouterVersion7
go test ./test/control -v -run TestRouterVersion8
```

### Testes do executável

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