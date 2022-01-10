# API de carrinho

## Regras
Após receber a requisição deverão ser aplicadas as seguintes regras:

### Regra número 1
Para cada produto você precisará calcular a porcentagem de desconto e isso deve ser feito consumindo um serviço gRPC fornecido por nós para auxiliar no seu teste. Utilize a imagem Docker para subir esse serviço de desconto e o arquivo proto para gerar o cliente na linguagem escolhida. Você pode encontrar como gerar um cliente gRPC nas documentações oficiais da ferramenta e em outros guias encontrados na internet.

### Regra número 2
Caso o serviço de desconto esteja indisponível o endpoint de carrinho deverá continuar funcionando porém não vai realizar o cálculo com desconto.

### Regra número 3
Deverá ser verificado se é black friday e caso seja, você deve adicionar um produto brinde no carrinho. Lembrando que os produtos brindes possuem a flag is_gift = true e não devem ser aceitos em requisições para adicioná-los ao carrinho (em uma loja virtual, esse produto não deveria ir para "vitrine"). A data da Black Friday fica a seu critério.

### Regra número 4
Deverá existir apenas uma entrada de produto brinde no carrinho.

### Regra número 5
É mandatório o uso de Docker para rodar a aplicação, de preferência utilizar docker-compose, pois facilitará o processo de correção.

## Pré-requisito
* Go 1.17: https://go.dev/dl/
* Docker: https://docs.docker.com/engine/install/
* Docker compose: https://docs.docker.com/compose/install/

## Configurações
Variáveis de ambiente:
```sh
#Indica o mês-dia do evento Black Friday. Por exemplo, 11-23 (23 de Novembo)
export BLACK_FRIDAY_DAY=12-30
#Modo "release" do framework gin (serviços REST). Além de "release", os outros valores possíveis são "debug" ou "test"
export GIN_MODE=release
#Porta de internet que o sistema ficará disponível localmente. Por exemplo, http://localhost:8080
export PORT=:8080
#Endereço do servidor de descontos (serviço de descontos)
export DISCOUNT_SERVER=new-backend-challenge_discount_1:50051
```

## Construir aplicação
Local:
- Na raiz do projeto faça:
```sh
cd src
go mod download
cd cmd
go build -o ./new-backend-challenge
```
---
Docker:
- Na raiz do projeto faça:
```sh
docker-compose build
```
ou
```sh
docker pull hashorg/hash-mock-discount-service
docker build -t new-backend-challenge:latest -f Dockerfile .
```

## Iniciar aplicação
Local:
- Na raiz do projeto faça:
```sh
cd src/cmd
./new-backend-challenge
```
---
Docker:
- Na raiz do projeto faça:
```sh
docker-compose up
```
ou
```sh
docker run -it --expose 50051 -p50051:50051 hashorg/hash-mock-discount-service:latest&
docker run -it --expose 8080 -p8080:8080 new-backend-challenge:latest&
```

## Executar testes
- Na raiz do projeto faça:
```sh
cd src
go test  ./... -v
```

## Endpoints
Observação: A aplicação local roda em http://localhost:8080/

### POST /carts/checkout
Criar carrinho com produtos (ID) e quantidades.


Requisição:
```json
{
    "products": [
        {
            "id": 1, //ID do produto
            "quantity": 2 // Quantidade a ser comprada do produto
        }
    ]
}
```

Resposta:
```json
{
    "total_amount": 20000, // Valor total da compra sem desconto
    "total_amount_with_discount": 19500, // Valor total da compra com desconto
    "total_discount": 500, // Valor total de descontos
    "products": [
        {
            "id": 1, //ID do produto
            "quantity": 2, //Quantidade
            "unit_amount": 10000, // Preço do produto em centavos
            "total_amount": 20000, // Valor total na compra desse produto em centavos
            "discount": 500, // Valor total de desconto em centavos
            "is_gift": false // É brinde?
        },
        {
            "id": 3, //ID do produto
            "quantity": 1, //Quantidade
            "unit_amount": 0, // Preço do produto em centavos
            "total_amount": 0, // Valor total na compra desse produto em centavos
            "discount": 0, // Valor total de desconto em centavos
            "is_gift": true // É brinde?
        }
    ]
}
```

### GET /products
Obter produtos (não brindes).


Resposta:
```json
[
    {
        "id": 1, //ID do produto
        "title": "Ergonomic Wooden Pants", //Nome do produto
        "description": "Deleniti beatae porro.", //Descrição do produto
        "amount": 15157, // Preço do produto em centavos
        "is_gift": false // É brinde?
    },
    {
        "id": 2, //ID do produto
        "title": "Ergonomic Cotton Keyboard", //Nome do produto
        "description": "Iste est ratione excepturi repellendus adipisci qui.", //Descrição do produto
        "amount": 93811, // Preço do produto em centavos
        "is_gift": false // É brinde?
    }
]
```


### GET /products/1
Obter produto pelo ID (não brinde).


Resposta:
```json
{
    "id": 1, //ID do produto
    "title": "Ergonomic Wooden Pants", //Nome do produto
    "description": "Deleniti beatae porro.", //Descrição do produto
    "amount": 15157, // Preço do produto em centavos
    "is_gift": false // É brinde?
}
```

## Outras documentações

* [Swagger](/docs/swagger.yaml)

