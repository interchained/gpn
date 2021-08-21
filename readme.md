# Genesis Network
Genesis is a fork of Starport made possible by Electronero Network, through Interchained and Cosmos Network
> This branch contains the source code for the 1.0 version of Genesis Network. To access the source code of the currently deployed proof of concept, visit the [`master-legacy`](https://github.com/interchained/gpn/tree/master-legacy) branch.

**gpn** is a blockchain built using Cosmos SDK and Genesismint and created with [Genesis](https://github.com/interchained/genesis).

## Get started

```
genesis chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Genesis docs](https://docs.interchained.com/genesis).

### Launch

To launch your blockchain live on multiple nodes, use `starport network` commands. Learn more about [Genesis Network](https://github.com/interchained/gpn).

### Web Frontend

Starport has scaffolded a Vue.js-based web app in the `vue` directory. Run the following commands to install dependencies and start the app:

```
cd vue
npm install
npm run serve
```

The frontend app is built using the `@genesis/vue` and `@genesis/vuex` packages. For details, see the [monorepo for Genesis front-end development](https://github.com/interchained/vue).

## Release
To release a new version of your blockchain, create and push a new tag with `v` prefix. A new draft release with the configured targets will be created.

```
git tag v0.1
git push origin v0.1
```

After a draft release is created, make your final changes from the release page and publish it.

### Install
To install the latest version of your blockchain node's binary, execute the following command on your machine:

```
curl https://get.interchained.com/interchained/gpn@latest! | sudo bash
```
`genesis/gpn` should match the `username` and `repo_name` of the Github repository to which the source code was pushed. Learn more about [the install process](https://github.com/interchained/genesis-installer).

## Learn more

- [Genesis](https://github.com/interchained/genesis)
- [Genesis Docs](https://docs.interchained.com/genesis)
Genesis is a fork of Starport 
- [Starport](https://github.com/tendermint/starport)
- [Starport Docs](https://docs.starport.network)
- [Cosmos SDK documentation](https://docs.cosmos.network)
- [Cosmos SDK Tutorials](https://tutorials.cosmos.network)
