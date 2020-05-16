import React from "react";
import GitHubButton from 'react-github-btn'
import Nav from "./nav"
import Embed from './embed.png'
import Footer from '../Footer'

export default () => {
  return <div>
    <Nav />
    <section className="hero is-dark is-bold">
      <div className="hero-body">
        <div className="container">
          <div className="columns">
            <div className="column is-half">
              
              <h1 class="title">
                Simple CSGO Matches on Discord.
              </h1>
              <h2 class="subtitle">
                1v1s up to 5v5s - no configuration needed.
              </h2>
            </div>
            <div className="column">
              <div className="box fill-height">
                <h2 className="is-size-4 has-text-dark has-text-weight-bold">Host your own</h2>
                <p>NadeStack is fully open source, so you can run your own bot and make your own customizations. </p><p>
                  We use DatHost to run CSGO servers, so setup the bot, enter your DatHost details and you're ready to go.
                </p><br />
                <p>Currently in development. Expect big bugs.</p>
                <br></br>
                <GitHubButton href="https://github.com/th0mas/NadeStack" data-size="large" aria-label="Get th0mas/NadeStack on GitHub">Get on GitHub</GitHubButton>
              </div>
            </div>
            <div className="column">
              <div className="box fill-height">
              <h2 className="is-size-4 has-text-dark has-text-weight-bold">Host for me</h2>
              <p>Just add our Bot to your discord and start playing! We manage updates, new features and running servers.</p>

              <button class="button" title="Coming Soon" disabled>Coming Soon(ish)</button>
              </div>
            </div>
            </div>
            </div>
          </div>
    </section>

    <section className="section">
      <div className="container">
        <div className="columns">
          <div className="column is-one-third">
        <figure className="image">
        <img src={Embed} alt="Discord embed example"></img>
        </figure>
        </div>
        <div className="column">
          <h3 className="is-size-4 has-text-dark is-bold">10 man? Just <code>/play de_mirage</code></h3>
          <p>NadeStack will recognise teams and run a map veto before launching a local 128-tick server.</p>
          <br />
          <h3 className="is-size-4 has-text-dark is-bold">Warmup? Just <code>/dm de_cache</code></h3>
          <p>Run a 128tick deathmatch server to warmup on. On any map - including workshop maps.</p>
          <br />
          <h3 className="is-size-4 has-text-dark is-bold">Hotel? Trivago</h3>
          <p>haha yes</p>
        </div>
        </div>
      </div>
    </section>
    <Footer />
  </div>
}