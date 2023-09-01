Para adicionar SSL à sua API Go, você precisa gerar um certificado SSL e uma chave privada. Existem várias maneiras de fazer isso, mas uma maneira fácil é usar o OpenSSL.

1. Certifique-se de ter o OpenSSL instalado na sua máquina. Se você estiver usando o Ubuntu, pode instalá-lo usando o seguinte comando:

```
sudo apt-get install openssl
```

2. Em seguida, você precisa gerar uma chave privada:

```
openssl genrsa -out server.key 2048
```

Este comando irá gerar uma chave privada de 2048 bits e salvá-la em um arquivo chamado "server.key".

3. Em seguida, você precisa gerar um certificado SSL. Você pode fazer isso usando o seguinte comando:

```
openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650
```

Este comando irá gerar um certificado SSL auto-assinado e salvá-lo em um arquivo chamado "server.crt".

4. Agora, você precisa configurar sua API para usar o certificado SSL e a chave privada. Você pode fazer isso usando o seguinte código:

```go
package main

import (
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello, world!"))
    })

    err := http.ListenAndServeTLS(":443", "server.crt", "server.key", nil)
    if err != nil {
        panic(err)
    }
}
```

Este código configura um manipulador para a raiz da sua API e, em seguida, inicia o servidor HTTP com a opção `ListenAndServeTLS`, que especifica o certificado SSL e a chave privada a serem usados.

5. Certifique-se de que sua API esteja ouvindo na porta 443, que é a porta padrão para HTTPS. Se você estiver executando sua API em uma máquina local, precisará acessá-la em https://localhost. Se você estiver executando sua API em um servidor remoto, precisará acessá-la em https://seu-domínio.com.

Observação: lembre-se de que o certificado SSL autoassinado não é confiável para uso em produção. Para um ambiente de produção, você deve obter um certificado SSL assinado por uma autoridade de certificação confiável.