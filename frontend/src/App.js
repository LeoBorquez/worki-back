import React from 'react';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom' //se agrega react router
import { Layouts } from './components/Layouts' 
import './App.css'; 
import { Home } from './pages/Home'
import { Worki } from './pages/Worki'
import { NoMatch } from './pages/NoMatch'
import { Details } from './pages/Details'
import { Registro } from './pages/Registro'


function App() {
  return (
    <React.Fragment>
      <Layouts>
      <Router>
        <Switch>
          <Route exact path="/" component={Worki} />
          <Route  path="/home" component={Home} />
          <Route path='/detail/:id' component={Details} />
          <Route path='/registro' component={Registro} />
          <Route component={NoMatch} />
        </Switch>
      </Router>
      </Layouts>
      </React.Fragment>
     
  );
}

export default App;
