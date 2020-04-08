import React, { useState, useEffect } from 'react'

export default () => {
  var id = 'TEST'
  const [profileInfo, setProfileInfo] = useState({})

  useEffect(() => {
    fetch(`/api/deeplink?id=${id}`)
      .then((response) => setProfileInfo(response.json()))
      .then((data) => console.log(data))
  })
  
  return (
    <>
    <h2 class="subtitle">Hi <b>t0mh</b>, please login with Steam to link up your Discord account. This should
    only have to be done once.</h2>
    
    <figure>
    <img src="https://steamcdn-a.akamaihd.net/steamcommunity/public/images/steamworks_docs/english/sits_large_noborder.png" className="margin-auto"></img>
    </figure>

    <br></br>
    <p className='is-size-7 has-text-grey-dark'>Your SteamID will be shared with NadeStack. This does <b>not</b> give NadeStack access to your steam account.</p>
    </>
  )
}