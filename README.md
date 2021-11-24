# ETCD configuration center

**A lightweight dashboard for ETCD, determined to build a dynamic configuration service in a centralized web panel, allows you to configure across environments**

A monomer application, frontend use vue + arco design, backend use etcd go client interface

## Get Started

``` bash
# run application
go mod tidy
go run main.go

# run just frontend
yarn serve

# build to update frontend
yarn build

```

## todo

* [x] base v3 function
  * [x] get/put/del/directory
  * [ ] auth
* [ ] frontend
* [ ] finish v3 function
  * [ ] tls
  * [ ] history
  * [ ] namespace
  * [ ] rollback
  * [ ] listen
  * [ ] canary publish
* [ ] test
* [ ] v2 support
