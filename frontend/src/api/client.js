let ws = new WebSocket("ws://localhost:8080/ws");

let connect = () => {
  ws.onopen = () => {
    console.log("Connected to WebSocket server");
  };

  ws.onmessage = (msg) => {
    console.log(msg);
  };

  ws.onclose = () => {
    console.log("Connection closed");
  };

  ws.onerror = (error) => {
    console.log("Socket error: ", error);
  };
};

let sendMsg = (msg) => {
  console.log("Sending message: ", msg);
  ws.send(msg);
};

export { connect, sendMsg };
