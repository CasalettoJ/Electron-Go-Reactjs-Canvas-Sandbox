import React, { Component } from 'react';
import { render } from 'react-dom';

export default class App extends Component {
  render() {
    return (
      <div>
        React loaded.
      </div>
    );
  }
}

render(<App />, document.getElementById('app'));