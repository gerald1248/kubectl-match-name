kubectl-match-name
==================

Access pods quickly without using the clipboard or listing all pods first. Here are some typical use cases:

```
$ kubectl logs `kubectl match-name proxy` -f
I0207 12:51:25.322909       1 server.go:444] Version: v1.10.0
...
$ kubectl match-name proxy
kube-proxy-cvtm5
$ kubectl match-name -k svc .
default-http-backend
$ kubectl match-name minikube
etcd-minikube kube-addon-manager-minikube kube-apiserver-minikube kube-controller-manager-minikube kube-scheduler-minikube
```

Let's take a closer look at the first example. What have we gained? This line replaces two common approaches. In many cases we would use the clipboard or type in the hash component of the name:

```
$ kubectl get po | grep proxy
kube-proxy-cf2df                            1/1     Running   0          5h
kube-proxy-8dj6v                            1/1     Running   0          5h
kube-proxy-s6wvq                            1/1     Running   0          5h
$ kubectl logs kube-proxy-cf2df -f
```

What about a one-liner? It's straightforward but hardly concise:

```
$ kubectl logs `kubectl get po --no-headers | grep proxy | cut -d' ' -f1 | head -n1` -f
```

Why not just a shell script? Mainly because it probably will not work as intended on Windows. It is hard to overestimate the number of Kubernetes users using `kubectl` via Git Bash or PowerShell.

Note that although the binary itself is called `kubectl-match_name`, it is invoked by the command `kubectl match-name`.

## Usage
```
$ kubectl match-name -h
Usage: kubectl-match_name [-kubeconfig=PATH] [-a] [-k KIND] [-n NAMESPACE] REGEX
  -A    match in all namespaces
  -a    return all matching objects
  -c    count matching objects (implies -a)
  -k string
        match objects of kind (default "pod")
  -kubeconfig string
        absolute path to the kubeconfig file
  -n string
        namespace
```

The search expression is interpreted by the Golang `regexp` package (sadly not PCRE).

## Run
To try the plugin on your computer, download one of the release binaries above (Linux, Mac, Windows).

## Build
```
$ make
$ sudo make install
```

