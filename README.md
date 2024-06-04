# Do-list app

This app uses [Bun](https://bun.sh/) [Vite](https://bun.sh/guides/ecosystem/vite) [SWC](https://swc.rs/) for work so Unix-like environment reqired

## Hardware requirements

- CPU >= 4 core
- RAM >= 16Gb

## System requirements

- __Linux:__ (Prefered) Ubuntu >= 22.04, Debian 12 or later
- __MacOS:__ 10.14 or later
- __Windows:__ (10, 11) + WSL2 (Work supported ONLY inside WSL2)

## Software

- __Docker [Download](https://docs.docker.com/engine/install/)__ Docker Desktop or Docker Engine latest
- __Golang [Download](https://go.dev/dl/)__ >= 1.21

## For Windows users

1. [Install Docker](https://docs.docker.com/desktop/install/windows-install/)

2. [Install WSL2](https://learn.microsoft.com/ru-ru/windows/wsl/install) (Optional, no need to do that if Docker installed)

3. [Install Ubuntu 22.04 in WSL](https://apps.microsoft.com/store/detail/ubuntu-22042-lts/9PN20MSR04DW)

4. [Create Access Token in Gitlab](https://gitlab.tehzor.com/-/profile/personal_access_tokens) with this type of rights `read_api, read_repository, write_repository, read_registry, write_registry`

5. Run Ubuntu in terminal:

```ps1
ubuntu2204.exe
```

6. Run commands from [Fast start](#Fast-start-(ubuntu-family-amd_x86_64-and-WSL2-(ubuntu22.04)-only)) paragraph

## Fast start (ubuntu family amd_x86_64 and WSL2 (ubuntu22.04) only)

For fast install deps (except Docker) use this script

```sh
sudo apt update && sudo apt upgrade
sudo apt install curl tar git make

# Install Go v1.21.1
curl -LO https://go.dev/dl/go1.21.1.linux-amd64.tar.gz
#sudo rm -rf ~/go
tar -C ~/ -xzf go1.21.1.linux-amd64.tar.gz
echo 'export PATH=$PATH:~/go/bin' >> ~/.bashrc
rm go1.21.1.linux-amd64.tar.gz
source ~/.bashrc

# WARNING replace TOKEN_NAME and TOKEN_PASS with real ones created in Gitlab
# config
git config --global url."https://TOKEN_NAME:TOKEN_PASS@gitlab.com/".insteadof "https://gitlab.com/"
go env -w GOPRIVATE=gitlab.com/
echo "machine gitlab.com login MY_TOKEN_NAME password TOKEN_PASS" > ~/.netrc
```

### Before start

Create `.env` in working dir using `.env.example`:

```sh
cp ./.env.example ./.env
ln -s ./.env ./backend/.env
```

### Building

To build a binary file, run the command in the console in the project directory:

```shell
make build
```

#### Docker infrasructure for local development

```shell
# To run infrastructure and services:
make up
# To stop infrastructure and services:
make down
```
