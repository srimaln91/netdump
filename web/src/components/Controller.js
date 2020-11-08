import axios from 'axios';
import React from 'react';
import { Button, Form } from 'react-bootstrap';

class Controller extends React.Component {

  constructor() {
    super()
    this.handleChange = this.handleChange.bind(this)
    this.applyConfig = this.applyConfig.bind(this)
  }

  state = {
    isLoaded: false,
    netInterfaces: [],
    netInterfaceSelected: ""
  }

  handleChange(event) {
    this.setState({ netInterfaceSelected: event.target.value })
  }

  applyConfig() {
    axios.post("http://localhost:5050/apply_config", { "interface": this.state.netInterfaceSelected }).then(() => {
      console.log("Posted")
    }).catch(err => {
      console.log(err)
    })
  }

  componentDidMount() {
    axios.get("http://localhost:5050/interfaces").then(
      result => {
        this.setState({
          isLoaded: true,
          netInterfaces: result.data.data
        });
      },
      // Note: it's important to handle errors here
      // instead of a catch() block so that we don't swallow
      // exceptions from actual bugs in components.
      error => {
        this.setState({
          isLoaded: true,
          error
        });
      }
    );
  }

  render() {
    return (
      <Form>
        <Form.Group controlId="exampleForm.ControlSelect1">
          <Form.Label>Network Interface</Form.Label>
          <Form.Control as="select" onChange={this.handleChange} value={this.state.netInterfaceSelected}>
            {this.state.netInterfaces.map(int => (
              <option>{int}</option>
            ))
            }
          </Form.Control>
        </Form.Group>
        <Form.Group controlId="exampleForm.ControlInput1">
          <Form.Label>Pattern</Form.Label>
          <Form.Control type="text" placeholder="None" />
        </Form.Group>
        <Form.Group controlId="exampleForm.ControlTextarea1">
          <Button variant="primary" class="pull-right" onClick={this.applyConfig}>Apply</Button>
        </Form.Group>
      </Form>
    )
  }
}

export default Controller
