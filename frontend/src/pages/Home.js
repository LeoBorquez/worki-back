import React, { Component } from 'react';
import '../App.css';
import { JobsList } from '../components/JobsList'
import { SearchBar } from '../components/SearchBar'
import { NavigationBar } from '../components/NavigationBar'
import {  Container } from 'react-bootstrap'
import Footer from '../components/Footer';
import Default from '../pages/Default'

export class Home extends Component {

        state = { usedSearch: false, results: []}

        _handleResults = (results) => {
          this.setState({results, usedSearch: true})
        }
      
        _renderResult () {
          return this.state.results.length === 0
          ? <div className="box">Sin Resultados <span role="img" aria-label="sad">☹️</span></div>
          : <JobsList jobs={this.state.results} /> 
        }

        _renderDefault () {
          return this.state.results.length === 0
          ? <Default />
          : <div className="box">Error <span role="img" aria-label="eyes">👀</span></div> 
        }
        

        render() {
                return (
                <div>
                <SearchBar  onResults={this._handleResults}/>
                {/*<NavigationBar />*/}
                <Container>
                {this.state.usedSearch
                ? this._renderResult()
                : this._renderDefault()}
                </Container>
                <Footer />
                </div>
                )
        }
}

export default Home;