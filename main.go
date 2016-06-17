package main

import (
    "fmt"
    "strings"
    "bufio"
    "os"
    "bytes"
)

const install string = `RUN apt-get update && \
    apt-get install -y wget=1.17.1-1ubuntu1 && \
    wget https://apt.puppetlabs.com/puppetlabs-release-pc1-"$UBUNTU_CODENAME".deb && \
    dpkg -i puppetlabs-release-pc1-"$UBUNTU_CODENAME".deb && \
    rm puppetlabs-release-pc1-"$UBUNTU_CODENAME".deb && \
    apt-get update && \
    apt-get install --no-install-recommends -y puppet-agent="$PUPPET_AGENT_VERSION"-1"$UBUNTU_CODENAME" && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

RUN /opt/puppetlabs/puppet/bin/gem install r10k:"$R10K_VERSION" --no-ri --no-rdoc`

const run string = `RUN apt-get update && \
    /opt/puppetlabs/puppet/bin/r10k puppetfile install --moduledir /etc/puppetlabs/code/modules && \
    /opt/puppetlabs/bin/puppet apply manifests/init.pp --verbose --show_diff  --summarize && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*`

const copyPuppetfile string = "COPY Puppetfile /"

const copyManifests string = "COPY manifests /manifests"

func main() {
    var buffer bytes.Buffer
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        buffer.WriteString(scanner.Text() + "\n")
    }
    value := buffer.String()

    r := strings.NewReplacer("PUPPET_RUN", run,
      "PUPPET_INSTALL", install,
      "PUPPET_COPY_PUPPETFILE", copyPuppetfile,
      "PUPPET_COPY_MANIFESTS", copyManifests)

    result := r.Replace(value)
    fmt.Println(result)
}
