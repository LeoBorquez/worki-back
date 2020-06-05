import React from 'react';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom' //se agrega react router
import { Layouts } from './components/Layouts' 
import './App.css';
import { Home } from './pages/Home'
import { Worki } from './pages/Worki'
import { NoMatch } from './pages/NoMatch'


function App() {
  return (
    <React.Fragment>
      <Layouts>
      <Router>
        <Switch>
          <Route exact path="/" component={Worki} />
          <Route  path="/home" component={Home} />
          <Route component={NoMatch} />
        </Switch>
      </Router>
      </Layouts>
      </React.Fragment>
    
  );
}

export default App;
