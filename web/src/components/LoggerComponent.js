
class LoggerComponent extends React.Component {
    constructor(props) {    super(props);    this.state = {      value: null,    };  }
    render() {
      return (
        <button className="square" onClick={() => alert('click')}>
          {this.props.value}
        </button>
      );
    }
  }
