# delve debugging in root mode for Ubuntu 22 04 

This repo shows how you can debug files in root mode

## Settings up $PATH and installing delve for root user 

First thing first you need `go` and `dlv` commands available for your root user. You can do it like this
```
# log in root user 
sudo su 
# add to $PATH paths to go binary, go install command install things in ~/go/bin folder even for root user
# command below assumes that golang is installed at "/usr/local/go" path
printf "\nexport PATH=$PATH:/usr/local/go/bin:~/go/bin\n" >> ~/.bashrc
# reload session 
exit
sudo su 
# install dlv 
go install github.com/go-delve/delve/cmd/dlv@latest
```
