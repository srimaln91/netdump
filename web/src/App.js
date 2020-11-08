import React from 'react';
import './App.css';
import Controller from './components/Controller';
import Logger from './components/Logger';
class App extends React.Component {

  render() {
    return (
      <React.Fragment>
        <div class="container-fluid">
          <div class="row">
            <aside class="col-2 px-3 py-5 fixed-top" id="left">

              <div class="list-group w-100">
                <Controller />
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
                <div class="col-12 py-4" style={{ height: 900 }}>
                  <Logger />
                </div>
              </div>
            </main>
          </div>
        </div>
      </React.Fragment>
    )
  }
}

export default App;
