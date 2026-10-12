# Encomendas + Gremlins

Exemplo pequeno com duas condições, um cliente VIP e contas de desconto e portes.
Os valores do código estão em cêntimos.

- Clientes VIP recebem 10% de desconto quando o subtotal é pelo menos 100 €.
- Encomendas com subtotal abaixo de 50 € pagam 5 € de portes.
- O desconto usa divisão inteira e arredonda para baixo ao cêntimo.

Os três testes iniciais verificam uma encomenda de 30 €, uma de 120 € para VIP
e outra de 120 € sem VIP. Executam todas as instruções da função, mas não
verificam exatamente os limites de 50 € e 100 €.

```sh
go test -v -cover ./...
./bin/gremlins unleash
```

Depois, descomenta `TestLimiteDescontoVIP` e `TestLimitePortesGratis` em
`main_test.go` e executa os comandos novamente. Esses casos detetam as mutações
que alteram a inclusão dos valores nos limites.

A configuração `.gremlins.yaml` exclui os exemplos guardados na pasta
`apresentacao/`. O PowerPoint mantém-se separado deste exemplo.
