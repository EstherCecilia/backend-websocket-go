# 🕸️ Projeto WebSocket com Go e React

Este projeto demonstra a implementação de um servidor WebSocket em **Go** e um cliente WebSocket em **React**, onde o servidor pode enviar mensagens para usuários específicos conectados.

## 🚀 Funcionalidades

- Gerenciamento de múltiplos usuários conectados via WebSocket.
- Envio de mensagens específicas para um usuário identificado pelo `userID`.
- Mensagens de broadcast para todos os usuários conectados.
- Exemplo prático de interação em tempo real.

## 🛠️ Tecnologias Utilizadas

- **Backend**: Go (Golang) com WebSocket.
- **Frontend**: React.js para comunicação com o servidor WebSocket.



## 📦 Como Rodar o Projeto

### Backend (Servidor Go)

1. **Clone o repositório**:
```bash
git clone https://github.com/seu-usuario/websocket-go-react.git
cd websocket-go-react/server
 ```
2. **Inicie o servidor: Certifique-se de ter o Go instalado e execute**:
 ```bash
go run main.go
 ```

O servidor estará disponível em http://localhost:8080.

3. **Frontend (Cliente React)**
Instale as dependências:

 ```bash
cd ../frontend
npm install
npm start
 ```


O cliente estará disponível em http://localhost:3000.
