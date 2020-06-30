// export default App;
import React from 'react';
import TrafficLight from 'react-trafficlight';
import './App.css';

const App = () => (
  <div className="App">
    <GPIOTrafficLight red="17" yellow="27" green="22" />
  </div>
);

class GPIOTrafficLight extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      redOn: false,
      yellowOn: false,
      greenOn: false,
    }
  }

  componentDidMount() {
    for (const color in this.props) {
      const pin = this.props[color];

      fetch('/gpio/' + pin)
      .then((response) => response.json())
      .then((json) => {
        this.updateLight(color, json);
      })
      .catch((error) => console.error(error))
    }
  }

  render() {
    return (
      <TrafficLight
        Size="150"
        onRedClick={() => this.toggleLight('red')}
        onYellowClick={() => this.toggleLight('yellow')}
        onGreenClick={() => this.toggleLight('green')}

        RedOn={this.state.redOn}
        YellowOn={this.state.yellowOn}
        GreenOn={this.state.greenOn}
      />
    )
  }

  updateLight(color, json) {
    const button = color + 'On';
    this.setState({ [button]: (json.value === '0') });
  }

  toggleLight(color) {
    const pin = this.props[color];
    const button = color + 'On';
    const value = this.state[button] ? '1' : '0';

    fetch('/gpio/' + pin, {
      method: 'POST',
      headers: {
        Accept: 'application/json',
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({ value: value })
    })
    .then((response) => response.json())
    .then((json) => {
      this.updateLight(color, json);
    })
    .catch((error) => console.error(error));
  }
}

export default App;
