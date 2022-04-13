# README

## Requerimientos
- Go
- Terminal ubicada en el directorio donde se encuentra este documento

## Pedido `GET` a Sitio de la UH

El código de `uhclient.go` realiza un pedido `GET` a `www.uh.cu` en el puerto 80 (HTTP) e imprime la respuesta del servidor. Para ejecutarlo hacemos

```console
$ go run uhclient.go
```

## Patrón *Request/Reply*
Los archivos `client.go` y `server.go` contienen el código necesario para ejemplificar el patrón *Request/Reply*.

Cuando el cliente le envía un nombre al servidor, este lo registra de ser necesario en un archivo y le responde un mensaje de bienvenida al cliente.

Primero ejecutamos el servidor:

```console
$ go run server.go
```

y luego de que este informe que se encuentra escuchando, ejecutamos el cliente, pasándole un nombre:

```console
$ go run client.go andy
```