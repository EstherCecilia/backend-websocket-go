import React, { useState, useEffect } from 'react';

function App() {
  const [message, setMessage] = useState('');

  useEffect(() => {
    // Conectar-se ao servidor WebSocket
    const socket = new WebSocket('ws://localhost:8080/ws?userID=user1');

    // Função chamada sempre que uma nova mensagem é recebida
    socket.onmessage = (event) => {
      setMessage(event.data); // Atualiza o estado com a mensagem recebida
    };

    // Fechar a conexão WebSocket quando o componente for desmontado
    return () => {
      socket.close();
    };
  }, []);

  return (
    <div className="App">
      <h1>Notificação de Atualizações</h1>
      <p>{message}</p>
    </div>
  );
}

export default App;
