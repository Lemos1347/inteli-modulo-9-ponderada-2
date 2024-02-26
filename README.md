# Teste automatizados do simulador de dispositivos IoT

## Introdução

Este documento descreve os testes automatizados implementados para o simulador de dispositivos IoT. Os testes visam validar a funcionalidade do simulador, garantindo que a comunicação entre publicador e assinante ocorra corretamente e dentro do tempo esperado. Para facilitar a criação dos testes, juntei os dois módulos que estavam separados em um único módulo, o `simulation` que contém o código do publisher e do subscriber.

## Demonstração

https://github.com/Lemos1347/inteli-modulo-9-ponderada-2/assets/99190347/fc8860e5-dfc5-429b-b607-d09098dcb62d

## Testes Implementados

Em todos os testes eu inicio um publisher e um subscriber individualmente em goroutines separadas da principal do teste.

### TestSendingMessages

Verifica se os dados estão chegando corretamente no broker. O teste publica uma mensagem em um tópico específico e valida se a mensagem é recebida, por exemplo se o valor enviado não é algo vazio.

### TestMessageAcertivity

Testa a acurácia das mensagens enviadas e recebidas. Compara o valor enviado pelo publicador com o valor recebido pelo assinante, garantindo que os dados transmitidos sejam consistentes e precisos. Para esse caso, o teste utiliza-se de um channel para garantir um informação precisa para realizar a comparação.

### TestSendingMessagesTime

Avalia o tempo de envio das mensagens. O teste garante que as mensagens sejam enviadas e recebidas dentro de um limite de tempo aceitável, validando assim a eficiência temporal da comunicação. O simulador foi projeto para enviar mensagens entre 1 a 5 segundos. Foi colocado então um limite de 6 segundo entre as mensagens, caso ultrapasse esse limite, o teste falha.

## Como Executar

Para executar os testes automatizados, utilize os seguintes comandos:

> [!IMPORTANT]
> É necessário ter o [Go](https://go.dev/doc/install) instalado na máquina para executar os testes e qualquer broker MQTT rodando na máquina (recomendo o [Mosquitto](https://mosquitto.org/download/)).

Primeiro clone o projeto:

```bash
git clone https://github.com/Lemos1347/inteli-modulo-9-ponderada-2
```

Navegue até a pasta do projeto:

```bash
cd inteli-modulo-9-ponderada-2
```

Em seguida, rode um broker MQTT na máquina (aqui demonstrarei como rodar o Mosquitto):

```bash
mosquitto -c mosquitto.conf
```

Em seguida baixe as dependências do projeto:

```bash
cd simulation ; go mod tidy
```

E depois execute os testes:

```bash
go test -v
```

Esse comando ira executar todos os testes da forma mais verbosa possível, exibindo detalhes sobre a execução de cada teste.
