# Sistema de Consulta de Clima por CEP

Este projeto tem como objetivo desenvolver um sistema em Go que recebe um CEP, identifica a cidade correspondente e retorna as temperaturas atuais em três unidades de medida: Celsius, Fahrenheit e Kelvin. O sistema é hospedado no Google Cloud Run.

## Funcionalidade

O sistema possui as seguintes funcionalidades:

- **Receber um CEP válido de 8 dígitos**
- **Pesquisar a localização correspondente ao CEP utilizando a API ViaCEP**
- **Consultar a temperatura atual utilizando a API WeatherAPI**
- **Retornar a temperatura em Celsius, Fahrenheit e Kelvin**
- **Responder adequadamente nos seguintes cenários:**

### Respostas de Sucesso
- **Código HTTP:** 200
- **Response Body:**
```json
{
  "temp_C": 28.5,
  "temp_F": 83.3,
  "temp_K": 301.65
}
```

### Respostas de Erro

### CEP inválido (formato incorreto):
- **Código HTTP:** 422
- **Mensagem:** `invalid zipcode`

### CEP não encontrado:
- **Código HTTP:** 404
- **Mensagem:** `can not find zipcode`

## Requisitos

- O sistema deve validar o formato do CEP e garantir que seja um número válido de 8 dígitos.
- O sistema deve integrar com as APIs ViaCEP e WeatherAPI para obter as informações necessárias.

## Fórmulas de Conversão

### Celsius para Fahrenheit:
\[
F = C * 1.8 + 32
\]

### Celsius para Kelvin:
\[
K = C + 273
\]

## Tecnologias Utilizadas

- **Linguagem:** Go
- **APIs:**
  - ViaCEP (https://viacep.com.br/)
  - WeatherAPI (https://www.weatherapi.com/)
- **Hospedagem:** Google Cloud Run
- **Docker:** Para containerização da aplicação


### Testes Automatizados
O código inclui testes automatizados para verificar o funcionamento das funcionalidades principais.

```make test``` ou ```make test-cover```

### Deploy no Google Cloud Run
O sistema foi implementado e deployado no Google Cloud Run. Você pode acessar a versão hospedada através do link:
https://cep-weather-service-806797558058.us-central1.run.app/cep/11111111