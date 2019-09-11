# Sistema de Comunicação Distribuida
Tarefa solicitada dentro do escopo da graduação de Engenharia da Computação. 
# Pré-requisitos
- Linux (Debian foi utilizado por sua alta estabilidade)
- Golang e seus complementos
- RabbitMQ - **servidor**
-  AMQP como **cliente** RabbitMQ para GO
# Golang e complementos
Go foi definido como a linguagem desse projeto.
A versão atual do pacote GO (1.7~5), nos repositórios Debian, se encaixa perfeitamente no projeto visto que muitos requisitos futuros não têm suporte às últimas versões.

	$ apt install golang golang-go-tools
**Configurações:**
É necessário configurar/informar os caminhos dos binários ao sistema.

	$ export GOROOT=/usr/lib/go-1.7;
E, principalmente, exportar as variáveis $GOPATH e $GOBIN.
Há um script Shell de exemplo na pasta *_aux*, onde se pode entender facilmente quais os arquivos que devem ser modificados e como devem ser modificados para que funcione perfeitamente.

**Tutoriais utilizados:**
- [https://comoinstalar.info/go-golang-en-linux/](https://comoinstalar.info/go-golang-en-linux/)
- [https://github.com/golang/go/wiki](https://github.com/golang/go/wiki)
- [https://golang.org/doc/install](https://golang.org/doc/install)
	
