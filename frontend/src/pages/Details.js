import React, { Component } from 'react';
import '../App.css'; 
import PropTypes from 'prop-types'  //es bueno añadir proptyes para validar los datos q esta recibiendo 
import {  Container, Col, Row, Button } from 'react-bootstrap'
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

.btn{
    background-color: #ff4417;
    color: white;
  }

.btn:hover path {
    fill: red;
  }
.btn:hover {
    box-shadow: 0 .125rem .25rem rgba(0,0,0,.3)!important;
}


  .btn-secundary{
    background-color: #ebebeb; 
    color: #000;
  }

`;

export class Details extends Component {

    static propTypes = {
        match: PropTypes.shape({
            params: PropTypes.object,
            isExact: PropTypes.bool,
            path: PropTypes.string,
            url: PropTypes.string
        })
    }
 
    state = { movie: {}}
    
    _fetchMovie({ id }){
        fetch(`https://www.omdbapi.com/?apikey=${API_KEY}&i=${id}`)
    .then(res => res.json())
    .then(movie => {
        
        console.log({ movie })
        this.setState({ movie })
    })
    }

    _goBack (){
        window.history.back()
    }

    componentDidMount(){ // cuando el componente se ha renderizado al menos una vez
        console.log(this.props)
        const { id } = this.props.match.params
        this._fetchMovie({id})
    }

    render() {

        const { Title, Poster, Actors, Metascore, Plot } = this.state.movie

        return (
            <Styles>
                <div className="jumbo">
                <Container>
                <Row>
                    <Col>
                    <h5>

<svg onClick={this._goBack} style={{fill: '#fff'}} class="bi bi-arrow-left-square-fill" width="1em" height="1em" viewBox="0 0 16 16" fill="currentColor" xmlns="http://www.w3.org/2000/svg">
  <path fill-rule="evenodd" d="M2 0a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V2a2 2 0 0 0-2-2H2zm6.354 10.646a.5.5 0 0 1-.708.708l-3-3a.5.5 0 0 1 0-.708l3-3a.5.5 0 1 1 .708.708L6.207 7.5H11a.5.5 0 0 1 0 1H6.207l2.147 2.146z"/>
</svg> 
                    </h5>
                    </Col>
                    <Col align="right">
                    👨‍🚀 Iniciar Sesión
                    </Col>
                </Row>
               
                </Container>
                <Container>
                <Row>
                <Col sm={8} >
                    <div className="box">
                    <h1>{Title}</h1>
                    <strong>{Actors}</strong>
                    <hr></hr>
                    <p>
                    <img src={Poster} width="40%"/>
                    </p>
                    <p>
                    An interesting client is looking for a Full Stack Developer to work within a team on interesting products & projects.

What you’ll bring:

***BSc or BEng or not
***3-5 years experience of creating features & functionality writing code in Python or Java or similar
***Work on Python with Django
***Solid JavaScript working in React & Node.js
***CSS/ SCSS
***PostgreSQL
***Enjoy a DevOps approach utilising AWS
                    </p>
                    <hr></hr>
                    <p>
                    <a href="#">Be careful</a> Don’t provide your bank or credit card details when applying for jobs. If you see something suspicious <a href="#">report this job ad.</a>
                    </p>
                    
                        </div>
                    </Col>
                    
                    <Col sm={4} >
                    <div className="box">
                        <Row>
                            <Col>
                            <Button variant="default" size="lg" block>
                                Ofertar
                            </Button>
                            <hr></hr>
                            </Col>
                        </Row>
                        <Row>
                            <Col><Button variant="default" size="lg" className="btn-secundary" block>
                            <svg class="bi bi-heart-fill" width="1em" height="1em" viewBox="0 0 16 16" fill="currentColor" xmlns="http://www.w3.org/2000/svg">
  <path fill-rule="evenodd" d="M8 1.314C12.438-3.248 23.534 4.735 8 15-7.534 4.736 3.562-3.248 8 1.314z"/>
</svg> Guardar
                            </Button></Col>
                            <Col><Button variant="default" size="lg" className="btn-secundary" block>
                            <svg class="bi bi-envelope-fill" width="1em" height="1em" viewBox="0 1.5 16 16" fill="currentColor" xmlns="http://www.w3.org/2000/svg">
  <path fill-rule="evenodd" d="M.05 3.555A2 2 0 0 1 2 2h12a2 2 0 0 1 1.95 1.555L8 8.414.05 3.555zM0 4.697v7.104l5.803-3.558L0 4.697zM6.761 8.83l-6.57 4.027A2 2 0 0 0 2 14h12a2 2 0 0 0 1.808-1.144l-6.57-4.027L8 9.586l-1.239-.757zm3.436-.586L16 11.801V4.697l-5.803 3.546z"/>
</svg> Enviar
                            </Button></Col>
                        </Row>
                    </div>

                    <div className="box">

                   <p><strong>Publicado: </strong> 4 Jun 2020</p>
                   <p><strong>Ubicación: </strong>
Santiago,
Peñalolen</p>
<p><strong>Presupuesto: </strong>
$50.000 - $100.000</p>
<p><strong>Tipo:</strong> Por Hora</p>
<p><strong>Clasificaión: </strong>
Información & Comunicación Tecnologica</p>


                    </div>
                    </Col>

                </Row>
                
                
                <Row>
                <Col>
                <div className="box">
                <h4>Propuetas</h4>


                </div>
                </Col>
                </Row>
                </Container>

                </div>
            </Styles>
        );
    }
}

export default Details;