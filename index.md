<h1 align="center">Discord Statistics Bot</h1><br>
<p align="center">
  <a href="https://open-source.hue.observer/pre-micro/">
    <img alt="Hue" title="Hue" src="https://i.imgur.com/ZEdQ3nF.png" width="550">
  </a>
</p>

<p align="center">
  Beautiful statistics. Built with Go.
</p>

<p align="center">
  <a href="https://open-source.hue.observer/pre-micro/">
    The story behind going open source
  </a>
|
  <a href="https://open-source.hue.observer/">
    Other open-source projects by us
  </a>
</p>

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Feedback](#feedback)
- [Build Process](#build-process)
- [Acknowledgments](#acknowledgments)

## Introduction

[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](http://makeapullrequest.com)
[![Discord](https://img.shields.io/badge/Chat_On-Discord-008080.svg?style=flat-square)](https://l.hue.observer/discord)

This bot is the old code from the original Hue bot that I decided to make open-source for everyone to edit and improve. The main function is to view server and user information within your server, using this bot. On 1st July 2019, this code will become 'live' is the aspect that it will be built after each new commit and placed under a single bot 'Hue Open Source' on Discord. This will be made available to all.

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

- Follow the [Golang install guide](https://golang.org/doc/install) for getting started building a project with native code.
- Clone or download the repo
- `go get github.com/BaileyJM02/Discord-Statistics-Bot` to install dependencies
- `go run main.go` to start the bot without building
- `go build main.go` to build the bot into a native build file
- `./<built file>` to run the bot using the native build file

**Development Keys**: The `token` in `config/` are for development purposes and do not represent the actual application keys.

## Acknowledgments

Thanks to [JetBrains](https://www.jetbrains.com) for supporting me with an [educational licence](https://www.jetbrains.com/student/).
