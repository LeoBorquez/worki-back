import React, { Component } from 'react';
import { Row, Container, Col } from 'react-bootstrap'
import styled from 'styled-components'


const Styles = styled.div`
.box {
    padding-top: 10px;
    border-radius: .25rem!important;
    box-shadow: 0 .125rem .25rem rgba(0,0,0,.075)!important;
    background-color: #fff;
}

`;

export class JobsList extends Component {
    render() {
        return (
            <Styles>
                <Container>
                    <Row >
                        <Col sm={8} className="box">
                        <a href="ver_aviso.html"><h4>Jardinero con Experiencia</h4></a>
                        <span className="fecha">Hace 7 Días</span>
                        <p>
                        Worki SA.</p>
                        <p><strong>Región Metropolitana</strong> > Peñalolen</p>
                        <p><strong>Clasificación:</strong> Jardineria & Paisajismo.</p>
                        <ul>
                            <li>Mantención de jardines</li>
                            <li>Arreglo de riego</li>
                            <li>
                            preferente con licencia de conducir
                            </li>
                        </ul>
                        <p>
                        necesita contratar jardineros con experiencia comprobada 
                        Horario lunes a viernes 08 a 18:00 hrs. Documentos al día. POR FAVOR SOLO EXPERIENCIA EN EL RUBRO.
                        Contacto: +56 9 7112201
                        </p>
                        <span className="badge badge-secondary">JARDINERIA</span>
                        <span className="badge badge-secondary">REMODELACION</span>
                        <span className="badge badge-secondary">CARPINTERIA</span>
                        <hr></hr>
                        </Col>
                        <Col sm={4}>sm=4</Col>
                    </Row>
                </Container>
            </Styles>
        );
    }
}

export default JobsList;