# Clone Eco360

Projeto para aprendizado, clonando o app Eco360

Rodando o projeto
1. Realizar o download do Go [Download](https://go.dev/dl/)
2. Para saber se a instalação foi realizada com sucesso, basta rodar o seguinte comando , dentro do CMD
3. `go version`
4. Abrir o projeto dentro do VsCode, rodar o seguinte comando
5. `go mod tidy`
6. Logo depois será feito o download das bibliotecas necessárias para o funcionamento
7. Depois de concluir o download, rodar o projeto com o seguinte comando:
8. go run main.go


 Com o projeto rodando, pode ser feito o teste nas seguintes rotas

    localhost:3000/
Deverá exibir uma mensagem 

    Pagina de teste
## Outras rotas

### Lista todos os clientes

    localhost:3000/cliente
Lista todos os clientes

### Listando apenas um cliente (codigo)
Filtro pelo código do cliente

    localhost:3000/cliente/{codigo}



