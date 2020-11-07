import React from 'react';
import { LazyLog, ScrollFollow } from 'react-lazylog';
class LoggerComponent extends React.Component {

    url = 'http://localhost:5050';

    constructor(props) {    super(props);    this.state = {      value: null,    };  }
    render() {
      return (
        <ScrollFollow
        startFollowing={true}
        render={({ follow, onScroll }) => (
          <LazyLog url={this.url} stream enableSearch follow={follow} onScroll={onScroll} extraLines />
        )}
      />
      );
    }
  }

export default LoggerComponent