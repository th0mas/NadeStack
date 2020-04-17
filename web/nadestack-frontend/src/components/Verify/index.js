import React, { useState, useEffect } from 'react'
import { useParams } from 'react-router-dom'

export default () => {
  var queryString = window.location.search
  let { rune } = useParams()

  const [isLoading, setIsLoading] = useState(true)
  const [authState, setAuthState] = useState({})

  useEffect(() => {
    fetch("/api/auth/callback" + queryString + `&rune=${rune}`)
    .then(resp => resp.json())
    .then((data) => {
      setIsLoading(false)
      setAuthState(data)
      console.log(data)
    })
  }, [queryString, rune])


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
