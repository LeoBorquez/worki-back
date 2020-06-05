import React from 'react';
import { Nav, Navbar, NavbarBrand } from 'react-bootstrap'
import styled from 'styled-components'

const Styles = styled.div`

.navbar {
    background-color: #222;
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
            <Navbar.Brand href="/"><div className="logo">OTTO</div></Navbar.Brand>
            <Navbar.Toggle aria-controls="basic-navbar-nav" />
            <Navbar.Collapse id="basic-navbar-nav">
            <Nav className="ml-auto">
                <Nav.Item><Nav.Link href="/">HOME</Nav.Link></Nav.Item>
                <Nav.Item><Nav.Link href="/dropzone">Drag and Drop</Nav.Link></Nav.Item>
                <Nav.Item><Nav.Link href="/about">ABOUT</Nav.Link></Nav.Item>
                <Nav.Item><Nav.Link href="/table">Table</Nav.Link></Nav.Item>
            </Nav>
            </Navbar.Collapse>
        </Navbar>
    </Styles>
)