import React, { Component } from 'react';
import PropTypes from 'prop-types'
import { Row, Col, Button } from 'react-bootstrap'
import styled from 'styled-components'
import { Job } from './Job'

const Styles = styled.div`
.btn{
    background-color: #ff4417;
    color: white;
  }

.btn:hover {
    box-shadow: 0 .125rem .25rem rgba(0,0,0,.3)!important;
  }
`;

export class JobsList extends Component {

    static propTypes = {
        jobs: PropTypes.array
    }

    render() {
        const { jobs } = this.props
        return (
        <Styles>
            <Row>
                <Col sm={8} >
                {
                jobs.map(job=> {
                    
                    return (
                            <div key={job.imdbID}>
                                <Job 
                                    id={job.imdbID}
                                    type={job.Type}
                                    title={job.Title}
                                    year={job.Year}
                                    poster={job.Poster}
                                />
                            </div>
                            )
                    })
                }
                </Col>
                <Col sm={4}>
                <div className="box">
                <center>
                <strong>No estas en Worki? <span role="img" aria-label="sorprendido">😮</span></strong>
                <br></br><br></br>
                <Button variant="light" className="btn">Crear cuenta</Button>
                </center>
                </div>
                </Col>
            </Row>
         </Styles>
        );
    }
}

export default JobsList;