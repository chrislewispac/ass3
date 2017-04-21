# Strassen Assignment 3

## I'm a TF, how do I compile this thing?

#### First, Does your grading server have go installed?

```
go version
//go version go1.8 darwin/amd64
```

#### If that is blank, let's just install go real quick....

```
sudo wget https://storage.googleapis.com/golang/go1.8.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.8.linux-amd64.tar.gz
sudo export PATH=$PATH:/usr/local/go/bin
```

#### Cool, Go is installed. Now everything should work normally.

```
make //?? Awesome!
make test //?? Oh boy! But wait. you have to have this -> https://github.com/smartystreets/goconvey
```
