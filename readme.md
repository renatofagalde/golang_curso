## packages
1# pacote main para identificar quando for gerar algo exe
2# go run primeiro.go

## pacotes
para criar um modulo: go mod init modulo
um modulo é o equivalente ao pom ou package.json
go build, compila tudo que está na raiz para baixo

importando pacotes:
    ```go get github.com/badoux/checkmail```
<br/>
removendo pacotes não utilizados:
    ```go mod tidy```

## 
```
❯ pwd
/home/renatofagalde/desenvolvimento/golang_curso/31_testes_automatizados/enderecos
❯ go test ./... -v --cover
=== RUN   TestTipoDeEndereco
=== PAUSE TestTipoDeEndereco
=== RUN   TestTipoDeEnderecoErrado
=== PAUSE TestTipoDeEnderecoErrado
=== CONT  TestTipoDeEndereco
--- PASS: TestTipoDeEndereco (0.00s)
=== CONT  TestTipoDeEnderecoErrado
--- PASS: TestTipoDeEnderecoErrado (0.00s)
PASS
introducao-testes/enderecos     coverage: 100.0% of statements
ok      introducao-testes/enderecos     (cached)        coverage: 100.0% of statements
``

### gerando arquivo
go test --coverprofile cobertura.txt

### usando o arquivo gerado
go tool cover --func=cobertura.txt

go tool cover --html=cobertura.txt

```