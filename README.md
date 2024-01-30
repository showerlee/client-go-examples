# client-go-examples
This repo is intent to practice k8s development via client-go SDK

## Prerequsite

### Setup local k8s via `kind`

- Install `kind` CMD via go

    ```
    go install sigs.k8s.io/kind@v0.20.0
    ```

- Create and apply k8s cluster via `kind`

    ```
    kind create cluster --image=kindest/node:v1.29.0 --name=dev
    kubectl cluster-info --context kind-dev
    ```

### Install go dependencies

- Install required packages

    ```
    go get k8s.io/apimachinery/pkg/apis/meta/v1
    go get k8s.io/client-go/kubernetes
    go get k8s.io/client-go/tools/clientcmd
    go get k8s.io/client-go/util/homedir
    ```

## in-cluster-configuration

### Compilation

```
cd in-cluster-configuration
GOOS=linux go build -o ./in-cluster .
```

### Image build

```
docker build -t in-cluster:v1 .
kind load docker-image in-cluster:v1 --name=dev
```

### Create ClusterRoleBinding

```
kubectl create clusterrolebinding default-view --clusterrole=view --serviceaccount=default:default
```

### Rollout pod

```
kubectl run -i in-cluster --image=in-cluster:v1
```
