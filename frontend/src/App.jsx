import React, { Component } from "react";
import "./App.css";
import { connect, sendMsg } from "./api/client.js";

class App extends Component {
  componentDidMount() {
    connect();
  }

  send() {
    console.log("hello");
    sendMsg("hello");
  }

  render() {
    return (
      <div className="App">
        <button onClick={this.send}>Hit</button>
      </div>
    );
  }
}

export default App;
