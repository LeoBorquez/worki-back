import React from 'react';
import { Nav, Navbar, Container } from 'react-bootstrap'
import styled from 'styled-components'

const Styles = styled.div`

.navbar {
    background-color: #0e1724;
}   

.navbar-brand, .navbar-nav .nav-link {
    color: #bbb;

    &:hover {
        color:white;
    }
}
`;

export const NavigationBar = () => (

    <Styles>
        
        <Navbar expand="lg">
        <Container>
            <Navbar.Brand href="/"><div className="logo">Dashboard</div></Navbar.Brand>
            <Navbar.Toggle aria-controls="basic-navbar-nav" />
            <Navbar.Collapse id="basic-navbar-nav">
                <Container>
                <Nav className="lg-auto">
                    <Nav.Link href="#home">Mis Proyectos</Nav.Link>
                    <Nav.Link href="#features">Inbox</Nav.Link>
                    <Nav.Link href="#pricing">Feedback</Nav.Link>
                </Nav>
                </Container>
            </Navbar.Collapse> 
            <Navbar.Collapse className="justify-content-end">
                <Navbar.Text style={{color: '#fff'}}>
                Signed in as: <a href="#login" style={{color: '#f1f1f1'}}>Mark Otto</a>
                </Navbar.Text>
            </Navbar.Collapse>
            </Container>
        </Navbar>
        
    </Styles>
)