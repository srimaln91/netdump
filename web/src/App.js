import React from 'react'

import {Form, Button} from 'react-bootstrap'
import LoggerComponent from './components/LoggerComponent'

import './App.css';
function App() {

  const url = 'http://localhost:5050';

  
  return (
    <React.Fragment>
   <div class="container-fluid">
    <div class="row">
        <aside class="col-2 px-3 py-5 fixed-top" id="left">

            <div class="list-group w-100">
            <Form>
              
              <Form.Group controlId="exampleForm.ControlSelect1">
                <Form.Label>Network Interface</Form.Label>
                <Form.Control as="select">
                  <option>1</option>
                  <option>2</option>
                  <option>3</option>
                  <option>4</option>
                  <option>5</option>
                </Form.Control>
              </Form.Group>
              <Form.Group controlId="exampleForm.ControlInput1">
                <Form.Label>Pattern</Form.Label>
                <Form.Control type="text" placeholder="None" />
              </Form.Group>
              <Form.Group controlId="exampleForm.ControlTextarea1">
                <Button variant="primary" class="pull-right">Apply</Button>
              </Form.Group>
            </Form>
            </div>

        </aside>
        <main class="col-10 invisible">

        </main>
        <main class="col offset-2 h-100">
            <div class="row bg-light">
                <div class="col-12 py-4">
                    Log output
                </div>
            </div>
            <div class="row bg-white">
                <div class="col-12 py-4" style={{height:900}}>
                    <LoggerComponent/>
                </div>
            </div>
        </main>
    </div>
</div>
    </React.Fragment>

  );
}

export default App;
