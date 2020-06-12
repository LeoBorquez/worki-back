import React, { Component } from 'react';
import {  Container, Col, Row } from 'react-bootstrap'
import styled from 'styled-components'

const Styles = styled.div`
.btn{
    background-color: #ff4417;
    color: white;
  }

.btn:hover {
    box-shadow: 0 .125rem .25rem rgba(0,0,0,.3)!important;
  }

  .footer {
    height: 50px;
    background-color: #e2e2e2;
    margin-top: 40px;
    padding-top: 30px;
    padding-bottom: 30px;
  }
`;


export class Footer extends Component {
    render() {
        return (
            <Styles>
                <div className="footer">
                <Container>
                    <Row>
                        <Col sm={3}>
                            Herramientas
                        </Col>
                        <Col sm={3}>
                            Compañia
                        </Col>
                        <Col sm={3}>
                            Contacto
                        </Col>
                        <Col sm={3}>
                            Empleadores
                        </Col>
                    </Row>
                </Container>
                </div>
            </Styles>
        );
    }
}

export default Footer;