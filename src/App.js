// export default App;
import React from 'react';
import TrafficLight from 'react-trafficlight';
import './App.css';

const App = () => (
  <div class="App">
    <TrafficLightContainer />
  </div>
);

class TrafficLightContainer extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      redOn: false,
      yellowOn: false,
      greenOn: false,
    }
  }

  render() {
    return (
      <TrafficLight
        Size="150"
        onRedClick={() => this.setState({ redOn: !this.state.redOn })}
        onYellowClick={() => this.setState({ yellowOn: !this.state.yellowOn })}
        onGreenClick={() => this.setState({ greenOn: !this.state.greenOn })}

        RedOn={this.state.redOn}
        YellowOn={this.state.yellowOn}
        GreenOn={this.state.greenOn}
      />
    )
  }
}

export default App;
