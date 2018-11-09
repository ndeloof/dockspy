# Docker API Spy

Dockspy allows to inspect traffic between your own container and docker daemon to track API usage and payload

## Usage

Run dockspy container bind-mounting actual docker daemon socker as `target.sock`.
```sh
docker run --rm --name dockspy -v /var/run/docker.sock:/target.sock ndeloof/dockspy
```

Then run your own container, which will rely on docker API, re-using volumes from dockerspy (which includes a spoofed /var/run/docker.sock)
```sh
docker run --volumes-from dockerspy your_container
```

## sample

dockspy :

```sh
➜  docker run --rm --name dockspy -v /var/run/docker.sock:/target.sock ndeloof/dockspy 
>  GET /_ping HTTP/1.1
Host: docker
User-Agent: Docker-Client/18.06.1-ce (linux)


<  HTTP/1.1 200 OK
Api-Version: 1.39
Docker-Experimental: true
Ostype: linux
Server: Docker/18.09.0-ce-beta1 (linux)
Date: Fri, 09 Nov 2018 09:29:07 GMT
Content-Length: 2
Content-Type: text/plain; charset=utf-8

OK
>  GET /v1.38/version HTTP/1.1
Host: docker
User-Agent: Docker-Client/18.06.1-ce (linux)


<  HTTP/1.1 200 OK
Api-Version: 1.39
Content-Type: application/json
Docker-Experimental: true
Ostype: linux
Server: Docker/18.09.0-ce-beta1 (linux)
Date: Fri, 09 Nov 2018 09:29:07 GMT
Content-Length: 597

{"Platform":{"Name":"Docker Engine - Community"},"Components":[{"Name":"Engine","Version":"18.09.0-ce-beta1","Details":{"ApiVersion":"1.39","Arch":"amd64","BuildTime":"2018-09-06T22:49:35.000000000+00:00","Experimental":"true","GitCommit":"78a6bdb","GoVersion":"go1.10.3","KernelVersion":"4.9.125-linu
<  xkit","MinAPIVersion":"1.12","Os":"linux"}}],"Version":"18.09.0-ce-beta1","ApiVersion":"1.39","MinAPIVersion":"1.12","GitCommit":"78a6bdb","GoVersion":"go1.10.3","Os":"linux","Arch":"amd64","KernelVersion":"4.9.125-linuxkit","Experimental":true,"BuildTime":"2018-09-06T22:49:35.000000000+00:00"}
```


analyzed container (nothing special here, just works as it is using a plain docker socket) :

```sh
➜  ~ docker run --volumes-from dockspy -it docker docker version
Client:
 Version:           18.06.1-ce
 API version:       1.38
 Go version:        go1.10.3
 Git commit:        e68fc7a
 Built:             Tue Aug 21 17:20:43 2018
 OS/Arch:           linux/amd64
 Experimental:      false

Server: Docker Engine - Community
 Engine:
  Version:          18.09.0-ce-beta1
  API version:      1.39 (minimum version 1.12)
  Go version:       go1.10.3
  Git commit:       78a6bdb
  Built:            Thu Sep  6 22:49:35 2018
  OS/Arch:          linux/amd64
  Experimental:     true


```
