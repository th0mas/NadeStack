import React from 'react'
import {
  BrowserRouter as Router,
  Switch,
  Route,
} from "react-router-dom"

import Login from './../LogIn'
import Verify from './../Verify'
import Home from './../Home'
import Logo from './../Logo'

export default () => {
  return (
    <div>
      <Router>
        <Switch>
        <Route path='/' exact><Home /></Route> {/* It works here to render the full home page. idk why */}
          <div className='columns is-centered is-vcentered full-height'>
            <div className='column is-one-third'>
              <div className='box'>
                <Logo />

                <Route path='/verify/:rune'><Verify /></Route>
                <Route path='/:rune' exact><Login /></Route>

              </div>
            </div>
          </div>
         
        </Switch>
      </Router>
    </div>
  )
}