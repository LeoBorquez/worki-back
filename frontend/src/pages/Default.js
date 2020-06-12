import React, { Component } from 'react';
import '../App.css'; 
import { Job } from '../components/Job'
import { Row, Col, Button } from 'react-bootstrap'
import styled from 'styled-components'

const Styles = styled.div`
.btn{
    background-color: #ff4417;
    color: white;
  }

.btn:hover {
    box-shadow: 0 .125rem .25rem rgba(0,0,0,.3)!important;
  }
`;

export class Default extends Component {

  constructor(props) {
    super(props);
 
    this.state = {
      Search: [],
      isLoading: false,
      error: null,
    };
  }
 
  componentDidMount() {
    fetch('https://www.omdbapi.com/?apikey=ecc70863&s=iron+man')
    .then(response => {
      if (response.ok) {
        return response.json();
      } else {
        throw new Error('Something went wrong ...');
      }
    })
      .then(data => this.setState({ Search: data.Search, isLoading: false }))
      .catch(error => this.setState({ error, isLoading: false }));
  }
 

    render() {
      const { Search, isLoading, error } = this.state;
      if (error) {
        return <p>{error.message}</p>;
      }

      if (isLoading) {
        return <p>Loading ...</p>;
      }
      return (
        <Styles>
          <marquee width="100%" direction="left">
          <strong>#Black Lives Matter</strong> 
          </marquee>
            <Row>
                <Col sm={8} >
                {
                Search.map(job=> {
                    
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

export default Default;