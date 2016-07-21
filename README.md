A _very much_ proof-of-concept showing how simple it is to build
Dockerfile pre-processors.

[![Build
Status](https://travis-ci.org/garethr/dockerfilepp-puppet.svg)](https://travis-ci.org/garethr/dockerfilepp-puppet)
[![Go Report
Card](https://goreportcard.com/badge/github.com/garethr/dockerfilepp-puppet)](https://goreportcard.com/report/github.com/garethr/dockerfilepp-puppet)
[![GoDoc](https://godoc.org/github.com/garethr/dockerfilepp-puppet?status.svg)](https://godoc.org/github.com/garethr/dockerfilepp-puppet)

In this case `dockerfilepp-puppet` is a trivial go application which takes a
Dockerfile on stdin and simply replaces some pre-defined values. The
idea would be to make Dockerfile declarative again, making multiple
Dockerfiles doing the same thing easier to maintain.

The examples centre around Puppet, but this is for demonstration
purposes only. You could imagine building your own library of DSL
extensions in a similar way, or extending into a general purpose tool.
For this purpose most of the work has been split out into a sepatate
library at
[github.com/garethr/dockerfile](https://github.com/garethr/dockerfilepp).

## Problem

Dockerfile is wonderfully simple when it comes to hello-world examples,
but the line-orientated nature and evolving best practices mean than
it's commonn for some quite crazy imperative bash juggling to make it's
way into what is best suited to a declarative build description. See
the best practice for installing debian packages if you don't believe me.

The same hoops are often jumped through in multiple Dockerfiles, so the
complex implementation details are copied and pasted into many places,
making maintenance more costly.

## Usage

So, given the following Dockerfile. We note that it is:

* Concise and declarative
* Totally not going to work if you pass it to docker because of the
  `PUPPET_*` instructions.

```dockerfile
FROM ubuntu:16.04
MAINTAINER Gareth Rushgrove "gareth@puppet.com"

ENV PUPPET_AGENT_VERSION="1.5.0" \
    R10K_VERSION="2.2.2" \
    UBUNTU_CODENAME="xenial"

PUPPET_INSTALL
PUPPET_COPY_PUPPETFILE
PUPPET_COPY_MANIFESTS
PUPPET_RUN

EXPOSE 80

CMD ["nginx"]
```

So lets run that through `dockerfilepp-puppet`.

First, build a binary. For this you'll need a Go environment, along with
the the mentioned github.com/garethr/dockerfilepp dependency. You can
install that with `go get`, or using the [glide package
manager](https://glide.sh/), simple type `glide up`.

The project uses [go-bindata](https://github.com/jteeuwen/go-bindata) to
make management of processors easier so first install that. This is only
required for building your own binaries, not using the resulting tool.
At some point I'll start releasing precompiled binaries.

Running `make build` should then be enough to generate a `dockerfilepp-puppet`
binary.

```
make build
```

Once you have the binary you can use it like so:

```
cat Dockerfile | ./dockerfilepp-puppet
```

This should output to stdout with a new Dockerfile which is:

* Much more verbose and much more imperative
* Going to work as an input to `docker build`

```dockerfile
FROM ubuntu:16.04
MAINTAINER Gareth Rushgrove "gareth@puppet.com"

ENV PUPPET_AGENT_VERSION="1.5.0" \
    R10K_VERSION="2.2.2" \
    UBUNTU_CODENAME="xenial"

RUN apt-get update && \
    apt-get install -y wget=1.17.1-1ubuntu1 && \
    wget https://apt.puppetlabs.com/puppetlabs-release-pc1-"$UBUNTU_CODENAME".deb && \
    dpkg -i puppetlabs-release-pc1-"$UBUNTU_CODENAME".deb && \
    rm puppetlabs-release-pc1-"$UBUNTU_CODENAME".deb && \
    apt-get update && \
    apt-get install --no-install-recommends -y puppet-agent="$PUPPET_AGENT_VERSION"-1"$UBUNTU_CODENAME" && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

RUN /opt/puppetlabs/puppet/bin/gem install r10k:"$R10K_VERSION" --no-ri --no-rdoc

COPY Puppetfile /
COPY manifests /manifests

RUN apt-get update && \
    /opt/puppetlabs/puppet/bin/r10k puppetfile install --moduledir /etc/puppetlabs/code/modules && \
    /opt/puppetlabs/bin/puppet apply manifests/init.pp --verbose --show_diff  --summarize && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

EXPOSE 80

CMD ["nginx"]
```
