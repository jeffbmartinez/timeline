# How to setup a new machine

## Spin up a new machine

Use digital ocean, spin up a new Ubuntu image.

## Create non-root user

- Log in as root (`ssh root@[new-machine-ip]`)
- Follow instructions on how to [create new user](https://www.digitalocean.com/community/tutorials/how-to-add-and-delete-users-on-an-ubuntu-14-04-vps), don't forget/skip the bit about giving the new user sudo access if necessary
- log out of root (just `exit`), log in as your new user to verify it works

## Install useful utilities

- Run `sudo apt-get update` to update apt-get with the latest packages
- Install stuff you'll need/want (`sudo apt-get install build-essential git emacs tree`)
  - Add `-y` to say automatically say yes to prompts if a script is made out of this
- Add whatever other packages you need/want

## Install golang

The apt-get golang package is outdated as of 2015.05.13, so install manually.

The following assume you are in the user's home directory:

- `wget https://storage.googleapis.com/golang/go1.4.2.linux-amd64.tar.gz`
- `sha1sum go1.4.2.linux-amd64.tar.gz` should print out `5020af94b52b65cc9b6f11d50a67e4bae07b0aff  go1.4.2.linux-amd64.tar.gz`
  - You can see if there are newer versions at [the golang download page](https://golang.org/dl/)
  - If you want install a new version, don't forget to remove the old version first to be safe. If you've followed instructions here to install, uninstalling should just be `rm -rf /usr/local/go`
- `sudo tar -C /usr/local -xzf go1.4.2.linux-amd64.tar.gz`
- `mkdir -p ~/go/bin ~/go/pkg ~/go/src`
- Add appropriate golang environment variables to `/etc/profile`
  - `sudo bash -c "echo -e '\nexport PATH=\$PATH:/usr/local/go/bin:~/go/bin\nexport GOPATH=\$HOME/go' >> /etc/profile"`
- Your current session won't have the updated golang environment variables, so log out and log in, or just add them manually for this session:
  - `export PATH=$PATH:/usr/local/go/bin`
  - `export GOPATH=$HOME/go`

## Install influxdb

- `wget http://get.influxdb.org/influxdb_0.9.0-rc30_amd64.deb`
  - Maybe want to check for a newer version at [the influxdb download page](http://influxdb.com/download/). **Scroll down** to see the section "Major Releases".
- `sudo dpkg -i influxdb_0.9.0-rc30_amd64.deb`. 0.9.0 rc30 is the newest release as of 2015.05.15.

## Install grafana

- `...`

## install node

- `...`

## Run the software

Timeline and grafana both depend on influxdb, so start influxdb first.

### Influxdb

- `sudo service influxdb start`

### Timeline

- `...`

### Grafana

- `...`
