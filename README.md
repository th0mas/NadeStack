# NadeStack 🧨
btec popflash rip off

A discord bot to provision CSGO Games. Allows users to link their steam accounts and directly provision CS:GO games.

## Requirements
* A discord bot user & associated token
* A postgres database (although you could modify `models/db.go` to use any database driver)

## Setup:

The frontend JavaScript needs to be built, e.g.
```bash
cd web/nadestack-frontend
yarn install
yarn build
```

Then the server can be compiled:
```bash
go build
```

Config vars need to be set, either through a `config.yml` file (see `example_config.yml`) or through envinronment vars (should be capitalized versions of the vars in `example_config`. These can be used together, with the environment variables taking precedence over the ones defined in a config file.
