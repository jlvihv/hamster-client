<div align=center>
<img src="./doc/logo-300.png" width=300" height="300" />
</div>
<div align=center>
<img src="https://img.shields.io/badge/golang-1.17-blue"/>
<img src="https://img.shields.io/badge/wails/v2-Beat3.0.0-lightBlue"/>
<img src="https://img.shields.io/badge/ant--design--vue-2.2.8-brightgreen"/>
<img src="https://img.shields.io/badge/go--libp2p-1.5.2-red"/>
</div>




[github](https://github.com/hamster-shared/hamster-provider): https://github.com/hamster-shared/hamster-client

hamster is a blockchain-based blockchain infrastructure service. Any computing device can easily access the Hamster network.

# Project Guidelines

## 1. Basic Introduction

### 1.1 Project Introduction

> Hamster Client is used to provide users with the ability to purchase and manage connections to their purchased resources . This includes the Marketplace, My Orders and My Resources modules.

### 1.2 Contributing Guide

Hi! Thank you for choosing Hamster.

Hamster is a blockchain that providers infrastructure service.

We are excited that you are interested in contributing to Hamster. Before submitting your contribution though, please make sure to take a moment and read through the following guidelines.

#### 1.2.1 Issue Guidelines

- Issues are exclusively for bug reports, feature requests and design-related topics. Other questions may be closed directly.

- Before submitting an issue, please check if similar problems have already been issued.

#### 1.2.2 Pull Request Guidelines

- Fork this repository to your own account. Do not create branches here.

- Commit info should be formatted as `[File Name]: Info about commit.` (e.g. `README.md: Fix xxx bug`)

- If your PR fixes a bug, please provide a description about the related bug.

- Merging a PR takes two maintainers: one approves the changes after reviewing, and then the other reviews and merges.

### 1.3 Version list

- main: 1.0.0 code, for prod
- develop: 2.0.0 dev code, for test

## 2. Getting started

### 2.1 Supported Platforms

- Windows 10/11 AMD64/ARM64
- MacOS 10.13+ AMD64
- MacOS 11.0+ ARM64
- Linux AMD64/ARM64

### 2.2 Dependencies

Wails has a number of common dependencies that are required before installation:

- Go 1.17+
- NPM (Node 15+)

### 2.3 wails Install

#### 2.3.1 Platform Specific Dependencies

- Windows

> Wails requires that the [WebView2](https://developer.microsoft.com/en-us/microsoft-edge/webview2/) runtime is installed. Some Windows installations will already have this installed. You can check using the `wails doctor` command (see below).

- Ubuntu

> Linux required the standard `gcc` build tools plus `libgtk3` and `libwebkit`. Rather than list a ton of commands for different distros, Wails can try to determine what the installation commands are for your specific distribution. Run `wails doctor` after installation to be shown how to install the dependencies. If your distro/package manager is not supported, please consult the[ Add Linux Distro](https://wails.io/docs/guides/linux-distro-support) guide.

- MacOS

> Wails requires that the xcode command line tools are installed. This can be done by running:
> `xcode-select --install`

#### 2.3.2 install command

Run `go install github.com/wailsapp/wails/v2/cmd/wails@latest` to install the Wails CLI.

#### 2.3.3 System Check

Running `wails doctor` will check if you have the correct dependencies installed. If not, it will advise on what is missing and help on how to rectify any problems.

If you can't find the command, please try to configure GOROOT, add the following statement to ~/.bash_profile: `export GOROOT=/usr/local/go`, and run `source ~/.bash_profile`

### 2.4 project compile

```bash
# use go mod And install the go dependency package
go mod tidy
# Compile
wails build
# open hamster
cd build
./hamster-client
```

### 2.5 development and debugging

```bash
# open frontend directory
cd frontend
# build frontend
npm run build
# go to root directory
cd ..
# use go mod And install the go dependency package
go mod tidy
# debugging 
wails dev
```



## 3. Technical selection

- Frontend: using [ant-design-vue](https://www.antdv.com/docs/vue/introduce/) based on [Vue](https://vuejs.org)，to code the page.
- Client: quickly write desktop applications with Wails[Wails](https://wails.io/docs/introduction)
- Backend: Use [gorm](https://gorm.io/docs/) to manipulate lightweight database [sqlite](https://www.sqlite.org/index.html)

## 4. Project Architecture

### 4.1 Project Layout

```
.
├── app				(Objects of the core)
├── build			(Compiled executable file)
├── config			(config)
├── ctx				(init app)
├── frontend			(frontend)
├── module
│   ├── account			(Managed Accounts)
│   ├── p2p			(p2p util)
│   ├── resource		(resource util)
│   └── wallet			(wallet util)
└── utils			(http util)
```

## 5. Features

- Resource Market: A trading market where computing power providers submit idle computing power to the market and configure prices. The client can choose the configuration and price resources to be purchased to form a transaction contract.
- My resources: A list of resources corresponding to the currently active orders, where you can view your purchased resources and renew them.
- Order List: List and details of all purchased resource orders, if a long time order is not processed, you have the option to cancel the order
- Link: Purchased resources can be connected here
- Setting: In the settings you can configure your account, public key, chain address, and gateway node

## 6. Knowledge base

### 6.1 Team blog
> https://github.com/hamster-shared

## 7. Maintainers
> ltyuanmu@gmail.com                                                       
                                                            
## 8. Contributors

Thank you for considering your contribution to hamster!

<a href="https://github.com/hamster-shared/hamster-client/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=hamster-shared/hamster-client" />
</a>

## 9. Commercial considerations

If you use this project for commercial purposes, please comply with the Apache2.0 agreement and retain the author's technical support statement.
