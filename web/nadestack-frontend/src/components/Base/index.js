import React from 'react'
import {
  BrowserRouter as Router,
  Switch,
  Route,
} from "react-router-dom"

import Login from './../LogIn'
import Verify from './../Verify'

export default () => {
  return (
    <div>
      <div className='columns is-centered is-vcentered full-height'>
        <div className='column is-one-third'>
      <div className='box'>
         <h1 className='title'>NadeStack <span role="img" aria-label="dynamite">🧨</span> <p className='is-size-7 has-text-grey-dark'>btec popflash</p></h1>
        <Router>
          <Switch>
          <Route path='/verify/:rune'><Verify /></Route>
          <Route path='/:rune'><Login /></Route>
          </Switch>
        </Router>
      </div>
      </div>
      </div>
    </div>
  )
}