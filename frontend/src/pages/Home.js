import React, { Component } from 'react';
import '../App.css';
import { SearchBar } from '../components/SearchBar'
import { JobsList } from '../components/JobsList'
import styled from 'styled-components'


const Styles = styled.div`

.home{
  background-color: #f1f1f1;
  min-height: 100vh;
}

`;

export const Home = () => (
            <Styles>
            <div className="home">
                <SearchBar />
                <br/>
                <JobsList />
            </div>
            </Styles>
            
          
        );

