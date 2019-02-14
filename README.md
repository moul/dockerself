# dockerself
:whale: runtime dockerizer

![](https://manfredtouron.com/uploads/2019/02/14/dockerself.png)

This is a PoC of a go program that calls the Docker API, inject itself (the binary) and switch to Docker for the execution.

It's like `syscall.Exec` within Docker.

:warning: the injected binary will run within Docker, so it must be build for the same architecture and should be build statically. (_a.k.a. it works well on Linux_).

## Usage

```golang
import (
        "flag"
        "fmt"

        "moul.io/dockerself"
)

func main() {
        dockerize := flag.Bool("dockerize", false, "dockerize")
        flag.Parse()
	    if *dockerize {
		        dockerself.Dockerize("ubuntu")
                return
		}
        if dockerize.WithinDocker() {
                fmt.Println("I'm inside the matrix (docker)")
        } else {
                fmt.Println("I'm in the real world :(")
        }
}
```

```console
$ ./example
I'm in the real world :(
$ ./example --dockerize
I'm inside the matrix (docker)
```

## Example

See [dockerself-example](./dockerself-example/).

```console
# go get moul.io/dockerself/dockerself-example
```

```console
# dockerself-example
within docker: false
env: [XDG_SESSION_ID=93 SHELL=/bin/bash TERM=screen SSH_CLIENT=82.254.90.14 50450 22 SSH_TTY=/dev/pts/0 LC_ALL=en_US.UTF-8 USER=moul ...]
starting interactive `/bin/sh` session...
$ ls -la
total 20
drwxrwxr-x 2 moul moul 4096 Feb 14 16:48 .
drwxrwxr-x 7 moul moul 4096 Feb 14 16:48 ..
-rw-rw-r-- 1 moul moul   95 Feb 14 16:48 go.mod
-rw-rw-r-- 1 moul moul 2578 Feb 14 16:48 go.sum
-rw-rw-r-- 1 moul moul  348 Feb 14 16:48 main.go
$
```

```console
# dockerself-example  --dockerize
within docker: true
env: [PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin HOSTNAME=72a73eaf8cbb TERM=xterm HOME=/root]
starting interactive `/bin/sh` session...
$ ls -la
total 10492
drwxr-xr-x  33 root root     4096 Feb 14 15:50 .
drwxr-xr-x  33 root root     4096 Feb 14 15:50 ..
-rwxr-xr-x   1 root root        0 Feb 14 15:50 .dockerenv
drwxr-xr-x   2 root root     4096 Feb  4 21:05 bin
drwxr-xr-x   2 root root     4096 Apr 24  2018 boot
drwxr-xr-x   5 root root      360 Feb 14 15:50 dev
-rwxr-xr-x   1 root root 10667575 Jan  1  1970 dockerself
drwxr-xr-x  31 root root     4096 Feb 14 15:50 etc
drwxr-xr-x   2 root root     4096 Apr 24  2018 home
drwxr-xr-x   8 root root     4096 May 23  2017 lib
drwxr-xr-x   2 root root     4096 Feb  4 21:03 lib64
drwxr-xr-x   2 root root     4096 Feb  4 21:02 media
drwxr-xr-x   2 root root     4096 Feb  4 21:02 mnt
drwxr-xr-x   2 root root     4096 Feb  4 21:02 opt
dr-xr-xr-x 266 root root        0 Feb 14 15:50 proc
drwx------   2 root root     4096 Feb  4 21:04 root
drwxr-xr-x   5 root root     4096 Feb  6 03:37 run
drwxr-xr-x   2 root root     4096 Feb  6 03:37 sbin
drwxr-xr-x   2 root root     4096 Feb  4 21:02 srv
dr-xr-xr-x  13 root root        0 Oct 13 20:12 sys
drwxrwxrwt   2 root root     4096 Feb  4 21:05 tmp
drwxr-xr-x  11 root root     4096 Feb  4 21:02 usr
drwxr-xr-x  13 root root     4096 Feb  4 21:04 var
$
container exited.
#
```
