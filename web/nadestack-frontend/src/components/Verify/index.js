import React, { useState, useEffect } from 'react'

export default () => {
  var queryString = window.location.search

  const [isLoading, setIsLoading] = useState(true)
  const [authState, setAuthState] = useState({})

  useEffect(() => {
    fetch("http://localhost:8080/api/auth/callback" + queryString)
    .then(resp => resp.json())
    .then((data) => {
      setIsLoading(false)
      setAuthState(data)
      console.log(data)
    })
  }, [queryString])


  return (
    isLoading
      ?
      <h1>Authenticating with Steam please wait...</h1>
      : authState.success ? <>
      <h1 className="subtitle">Successfully authenticated SteamID: {authState.steamID}</h1>
      <h1 className="subtitle bold">You may now close this tab</h1>
      </>
      : <h1>Unable to authenticate steam account: <code>{authState.error}</code></h1>
  )
}