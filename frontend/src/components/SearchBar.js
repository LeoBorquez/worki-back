import React, { Component } from 'react'
import { Form, FormControl, InputGroup, Button, Container, Row, Col } from 'react-bootstrap'
import styled from 'styled-components'

const API_KEY = 'ecc70863'

const Styles = styled.div`
.jumbo {
    position: relative;
	width: 100%;
    height: 110px;
    background-color: #0d3880;
    color: #fff;
    padding-top: 25px;
}

.btn-search {
    background-color: #ff4417;
    border-color: #ff4417;
    color: #fff;
}
`;

export class SearchBar extends Component {

    state = {
        inputMovie: ''
        }
        
        _handleChange = (e) => {
        this.setState({inputMovie: e.target.value})
        }
        
        _handleSubmit = (e) => {
        e.preventDefault()
        const {inputMovie} = this.state
        fetch(`https://www.omdbapi.com/?apikey=${API_KEY}&s=${inputMovie}`)
        .then(res => res.json())
        .then(results => {
        const { Search = [], totalResults = 0 } = results
        console.log({Search, totalResults})
        this.props.onResults(Search)
        })
        }

        _goBack (){
            window.history.back()
        }

    render() {
        return (
            <Styles>
                <div className="jumbo">
                <Container>
                <Row>
                    <Col><h5>Bienvenido</h5></Col>
                    <Col align="right" onClick={this._goBack}>👨‍🚀 Iniciar Sesión</Col>
                </Row>
                
                <Form onSubmit={this._handleSubmit}>
                <InputGroup >
                    <FormControl
                    placeholder="Buscar Trabajo"
                    aria-label="Buscar Trabajo"
                    aria-describedby="buscar-trabajo"
                    onChange={this._handleChange}
                    />
                    <InputGroup.Append>
                    <Button variant="outline-secondary" className="btn-search" type="submit">Search</Button>
                    </InputGroup.Append>
                </InputGroup>
                </Form>
                </Container>
                </div>
            </Styles>
        );
    }
}

export default SearchBar;