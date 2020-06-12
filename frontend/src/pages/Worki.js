import React, { Component } from 'react';
import logo from '../logo-worki.png';

export class Worki extends Component {
    render() {
        return (
            <div className="App">
              <header className="App-header">
                <img src={logo} className="App-logo" alt="logo" />
                <p>
                  PRONTO
                </p>
                <p>Portal de empleos para personas <code>independientes</code>
                </p>
                
                <a
                  className="App-link"
                  href="https://forms.gle/YDw5u5r7yXrBhNGJ6"
                  target="_blank"
                  rel="noopener noreferrer"
                >
                  Más información
                </a>
              </header>
            </div>
        );
    }
}

export default Worki;