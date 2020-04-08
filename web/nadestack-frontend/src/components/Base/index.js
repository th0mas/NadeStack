import React from 'react'
import {
  BrowserRouter as Router,
  Switch,
  Route,
  Link
} from "react-router-dom"

import Login from './../LogIn'
import Verify from './../Verify'

export default () => {
  return (
    <div>
      <div className='columns is-centered is-vcentered full-height'>
        <div className='column is-one-third'>
      <div className='box'>
         <h1 className='title'>NadeStack <span role="img">🧨</span></h1>
        <Router>
          <Route path='/verify'><Verify /></Route>
          <Route path='/'><Login /></Route>
        </Router>
      </div>
      </div>
      </div>
    </div>
  )
}