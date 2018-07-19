#Stash

This is for the demo.

Stash is a minimalistic web based file store.  Anyone is welcome to contribute to or use this software.

![alt text][demo]

# Running

A valid installation of golang is required to run Stash.

Running scripts may require running:

`chmod 775 .`

 - ./local-run will run stash on OSX.
 - ./build followed by ./docker-run will run stash if Docker is installed.
 - ./build followed by ./openshift-run will run stash when logged into an OpenShift project.

Stash will probably work on any Kubernetes installation.

# First time OpenShift Installation

While ./openshift-run is useful for quick iterations, first time setup should be done with:

```
./build
oc new-build --binary=true --name=stash -l app=stash
oc start-build stash --from-dir=. --follow=true
oc new-app -i stash -l app=stash
oc expose svc/stash
```

# Persistence

Attaching a persistent volume claim to the deployment config in OpenShift or Kubernetes at the directory "/resources/persistence", files will survive a pod restart.
  The persistent volume claim must allow for read and write access.

# Deleting files

A DELETE api call will remove files once they have been added to the stash. If a file named "foo.txt" is in the stash, a DELETE call to <host>/resources/persistent/foo.txt will remove it.

### Protecting files from deletion

By default, any DELETE api call can delete files in the stash.  Setting an environment variable "TOKEN" will require DELETE requests to have a header with key "token" associated to them.

# Why?

Sometimes you just need a simple file store.

[demo]: ./images/demo.png "Stash"

