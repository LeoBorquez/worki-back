import React, { Component } from 'react';
import PropTypes from 'prop-types'
import {Link } from 'react-router-dom'
import { Badge } from 'react-bootstrap'
import styled from 'styled-components'

const Styles = styled.div`
.box {
    margin-top: 10px;
    padding: 15px;
    padding-left: 15px;
    border-radius: .25rem!important;
    box-shadow: 0 .125rem .25rem rgba(0,0,0,.09)!important;
    background-color: #fff;
}

.box:hover {
    box-shadow: 0 .125rem .25rem rgba(0,0,0,.3)!important;
}

h4 {
    color: #2765cf;
 }
 
 .fecha{
     text-align: right;
     float: right;
     top: -32px;
     position: relative;
   }
`;

export class Job extends Component {
    static propTypes = {
        id: PropTypes.string,
        type: PropTypes.string,
        title: PropTypes.string,
        year: PropTypes.string,
        poster: PropTypes.string
    }
    render() {

        const { id, type, poster, title, year } = this.props

        return (
            <Styles>
            <div className="box">
                <Link to={`/detail/${id}`}><h4>{title}</h4></Link>
                <span className="fecha">Hace 7 Días</span>
                <p>{type}</p>
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
                {year}
                </p>
                <Badge variant="secondary">JARDINERIA</Badge>{' '}
                <Badge variant="secondary">REMODELACION</Badge>{' '}
                <Badge variant="secondary">CARPINTERIA</Badge>{' '}
                <hr></hr>
                </div>
                </Styles>

        );
    }
}
