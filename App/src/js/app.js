import React, { Component } from 'react';
import { render } from 'react-dom';
import grpc from 'grpc';

const PROTO_PATH =   `${__dirname}/../proto/hello.proto`;
const protoHello = grpc.load(PROTO_PATH).hello;
const Talk = new protoHello.Talk('localhost:50050', grpc.credentials.createInsecure())


export default class App extends Component {
  constructor(props) {
    super(props);
    this.state = {message: ''};
  }

  componentDidMount() {
    Talk.SendChat({message: 'yo what up'}, (err, response) => {
      this.setState({message: response.message});
    });
  }

  render() {

    return (
      <div>
        {`React loaded.\n${this.state.message}`}
      </div>
    );
  }
}

render(<App />, document.getElementById('app'));