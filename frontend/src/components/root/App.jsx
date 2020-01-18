import React, { Component } from 'react';
import { Route } from 'react-router-dom';
import { Switch } from 'react-router';
import { withRouter } from "react-router-dom";

import './App.css';

import { HelloPage } from '../hello-page/HelloPage';
import { MainChat } from '../main-chat/MainChat';
import { CodeShare } from '../code-share/CodeShare';

class App extends Component {
  render() {
    return (
      <div className="App">
        <Switch>
          <Route exact path='/' component={HelloPage} />
          <Route exact path='/mainchat' component={MainChat} />
          <Route exact path='/mainchat/:name' component={MainChat} />
          <Route path='/mainchat/:mainChannelName/:name' component={CodeShare} />
        </Switch>
      </div>
    );
  }
}

export default withRouter(App);
