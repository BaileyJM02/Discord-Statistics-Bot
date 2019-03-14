<h1 align="center"> Hue </h1> <br>
<p align="center">
  <a href="https://open-source.hue.observer/pre-micro/">
    <img alt="Hue" title="Hue" src="url" width="450">
  </a>
</p>

<p align="center">
  Beautiful statistic collecton. Built with Go.
</p>

<p align="center">
  <a href="https://open-source.hue.observer/pre-micro/">
    The story behind Hue
  </a>
|
  <a href="https://open-source.hue.observer/">
    Other open-source projects
  </a>
</p>

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Feedback](#feedback)
- [Build Process](#build-process)
- [Acknowledgments](#acknowledgments)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Introduction

[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](http://makeapullrequest.com)
[![Discord](https://img.shields.io/badge/Chat_On-Discord-008080.svg?style=flat-square)](https://l.hue.observer/discord)

View server and user information within your server. 

## Features

A few of the things you can do with Hue:

* Connect to the Discord API
* See the ammount of messages sent from a user
* See the ammount of links sent by a user
* See the ammount of members joining your server
* More coming soon...

## Feedback

Feel free to send us feedback on [Twitter](https://twitter.com/huediscord) or [file an issue](https://github.com/baileyjm02/Discord-Statistics-Bot/issues/new). Feature requests are always welcome. If you wish to contribute, please take a quick look at the [guidelines](./CONTRIBUTING.md)!

If there's anything you'd like to chat about, please feel free to join our [Discord Server](https://l.hue.observer/discord)!

## Build Process

- Follow the [React Native Guide](https://facebook.github.io/react-native/docs/getting-started.html) for getting started building a project with native code. **A Mac is required if you wish to develop for iOS.**
- Clone or download the repo
- `yarn` to install dependencies
- `yarn run link` to link react-native dependencies
- `yarn start:ios` to start the packager and run the app in the iOS simulator (`yarn start:ios:logger` will boot the application with [redux-logger](<https://github.com/evgenyrodionov/redux-logger>))
- `yarn start:android` to start the packager and run the app in the the Android device/emulator (`yarn start:android:logger` will boot the application with [redux-logger](https://github.com/evgenyrodionov/redux-logger))

Please take a look at the [contributing guidelines](./CONTRIBUTING.md) for a detailed process on how to build your application as well as troubleshooting information.

**Development Keys**: The `token` in `config/` are for development purposes and do not represent the actual application keys.

## Acknowledgments

Thanks to [JetBrains](https://www.jetbrains.com) for supporting me with an [educational licence](https://www.jetbrains.com/student/).
