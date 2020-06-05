import React, { Component } from 'react'
import { Form, FormControl, InputGroup, Button, Container,Jumbotron as Jumbo } from 'react-bootstrap'
import styled from 'styled-components'



const Styles = styled.div`
.jumbo {
    position: relative;
	width: 100%;
    height: 110px;
    background-color: #0d3880;
    color: #fff;
    padding: 25px;
}

.btn-search{
    background-color: #ff4417;
    border-color: #ff4417;
    color: #fff;
}
`;

export class SearchBar extends Component {
    render() {
        return (
            <Styles>
                <div className="jumbo">
                <Container>
                <h5>Bienvenido</h5>
                <Form>
                <InputGroup >
                    <FormControl
                    placeholder="Buscar Trabajo"
                    aria-label="Buscar Trabajo"
                    aria-describedby="buscar-trabajo"
                    />
                    <InputGroup.Append>
                    <Button variant="outline-secondary" className="btn-search">Buscar</Button>
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