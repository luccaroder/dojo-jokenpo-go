# dojo-jokenpo-go

A iddeia desse Dojo é a utilização de REST API em GO.

## Problema

Em uma partida de Jokenpo o sistema deve ser capaz de identificar entre dois jogadores
qual foi o vencedor ou se houve algum empate: https://dojopuzzles.com/problems/jokenpo/

## ideia de resolução

```
GET /jokenpo?player1=pedra&player2=tesoura
{ winner: player1, description: pedra ganha de tesoura }

GET /jokenpo?player1=papel&player2=papel
{ description: papel empata com papel }
```
Módulo externo https://github.com/gorilla/mux