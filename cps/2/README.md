# README

[Este enlace](https://pkg.go.dev/net/rpc) fue toda la documentaci&oacute;n consultada.

Para ejecutar el servidor, hacemos

```console
$ go run server.go
```

y para el cliente

```console
$ go run client.go NOMBRE_DE_USUARIO
```

El `NOMBRE_DE_USUARIO` se usar&aacute; para llamar a la funci&oacute;n `Iam`.

## Contenedor para el Servidor
En el archivo [rpc-server.dockerfile](../../rpc-server.dockerfile) del directorio raíz del proyecto se encuentran las instrucciones de *docker* para construir una imagen que corra a [server.go](server.go).

> Para ejecutar los siguientes comandos, debe abrir una terminal en el directorio ra&iacute;z del repo.

Para construir la imagen, hacemos: (**OJO: se va a descargar la imagen de *GO* que pesa m&aacute;s de 300 MB**)
```console
$ docker build -t distri-systems/rpc-server -f rpc-server.dockerfile .
```

y para levantar el contenedor:
```console
$ docker run --rm -it -p 1234:1234 distri-systems/rpc-server
```

Luego, de hacer esto, podemos ejecutar el cliente en el *host*.