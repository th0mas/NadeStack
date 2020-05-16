import React, { useState, useEffect } from 'react'
import { useParams } from 'react-router-dom'

export default () => {
  let { rune } = useParams()
  const [profileInfo, setProfileInfo] = useState({})
  const [isLoading, setIsLoading] = useState(true)
  const [error, setIsError] = useState({e: false, reason: ""})

  useEffect(() => {
    fetch(`/api/deeplink?rune=${rune}`)
        .then(r => r.status === 200 ? r : Promise.reject(r.status))
      .then(r => r.json())
      .then(data => {
        setProfileInfo(data)
        setIsLoading(false)
      })
      .catch((err) => {
        setIsError({e: true, reason: err})
        setIsLoading(false)
      })

  }, [rune])

  return (
    isLoading ? <h1>Loading.....</h1> :
      error.e ? <p>There was an error loading this link :( Reason: <code>{error.reason}</code></p> :
      <>
        <h2 className="subtitle">Hi <b>{profileInfo.User.DiscordNickname}</b>, please login with Steam to link up your Discord account. This should
    only have to be done once.</h2>

        <figure>
          <a href={profileInfo.Payload}>
            <img src="https://steamcdn-a.akamaihd.net/steamcommunity/public/images/steamworks_docs/english/sits_large_noborder.png" className="margin-auto"
    alt="steam login button"/>
          </a>
        </figure>

        <br/>
        <p className='is-size-7 has-text-grey-dark'>Your SteamID will be shared with NadeStack. This does <b>not</b> give NadeStack access to your steam account.</p>
      </>
  )
}