# Makefile for multi-host testing using ginkgo

# configuration variables
BRIDGE_ADDRESS = 10.0.0.1
NODE1 = 10.0.0.11
NODE2 = 10.0.0.12
POD1 = 10.0.0.21
GINKGO = $(GOPATH)/bin/ginkgo
PLACEMAT_DATADIR = /var/scratch/placemat

export BRIDGE_ADDRESS NODE1 NODE2 POD1 GINKGO

# non-configuration variables
SSH_PRIVKEY := ./mtest_key
OUTPUT := ./output
UBUNTU_IMAGE := ubuntu-18.04-server-cloudimg-amd64.img
PLACEMAT = $(OUTPUT)/placemat
PMCTL = $(OUTPUT)/pmctl
CLUSTER_YAML = $(OUTPUT)/cluster.yml
EXAMPLE_CLUSTER_YAML = $(OUTPUT)/cluster.example.yml

export SSH_PRIVKEY PLACEMAT PMCTL CLUSTER_YAML EXAMPLE_CLUSTER_YAML

GENERATED_FILES = $(OUTPUT)/placemat \
	$(OUTPUT)/pmctl \
	$(OUTPUT)/cluster.yml \
	$(OUTPUT)/user-data_node1.yml \
	$(OUTPUT)/user-data_node2.yml \
	$(OUTPUT)/network1.yml \
	$(OUTPUT)/network2.yml \
	$(OUTPUT)/cluster.example.yml \
	$(OUTPUT)/network-config.example.yml \
	$(OUTPUT)/user-data.example.yml \
	$(UBUNTU_IMAGE)

all: test

$(UBUNTU_IMAGE):
	curl -sSLf -o $@ https://cloud-images.ubuntu.com/releases/18.04/release/$@

$(OUTPUT)/cluster.yml: cluster.yml
	mkdir -p $(OUTPUT)
	sed -e 's/@BRIDGE_ADDRESS@/$(BRIDGE_ADDRESS)/' \
		-e 's/@POD1@/$(POD1)/' \
		-e 's/@UBUNTU_IMAGE@/$(UBUNTU_IMAGE)/' $< > $@

$(OUTPUT)/user-data_node1.yml: user-data.yml
	mkdir -p $(OUTPUT)
	sed -e 's/@HOST_NAME@/node1/' $< > $@

$(OUTPUT)/user-data_node2.yml: user-data.yml
	mkdir -p $(OUTPUT)
	sed -e 's/@HOST_NAME@/node2/' $< > $@

$(OUTPUT)/network1.yml: network.yml
	mkdir -p $(OUTPUT)
	sed -e 's/@BRIDGE_ADDRESS@/$(BRIDGE_ADDRESS)/' \
		-e 's/@HOST@/$(NODE1)/' $< > $@

$(OUTPUT)/network2.yml: network.yml
	mkdir -p $(OUTPUT)
	sed -e 's/@BRIDGE_ADDRESS@/$(BRIDGE_ADDRESS)/' \
		-e 's/@HOST@/$(NODE2)/' $< > $@

$(OUTPUT)/cluster.example.yml: ../examples/cluster.example.yml
	mkdir -p $(OUTPUT)
	sed -e 's/^url: .*/file: ..\/$(UBUNTU_IMAGE)/' $< > $@

$(OUTPUT)/network-config.example.yml: ../examples/network-config.example.yml
	mkdir -p $(OUTPUT)
	cp $< $@

$(OUTPUT)/user-data.example.yml: ../examples/user-data.example.yml
	mkdir -p $(OUTPUT)
	cp $< $@

$(OUTPUT)/placemat:
	mkdir -p $(OUTPUT)
	cd ..; go get ./... && GOBIN=$(abspath $(OUTPUT)) go install ./pkg/$(notdir $@)

$(OUTPUT)/pmctl:
	mkdir -p $(OUTPUT)
	cd ..; go get ./... && GOBIN=$(abspath $(OUTPUT)) go install ./pkg/$(notdir $@)

test: $(GENERATED_FILES)
	sudo rm -rf $(PLACEMAT_DATADIR)/volumes/node1
	sudo rm -rf $(PLACEMAT_DATADIR)/volumes/node2
	$(GINKGO) -v

clean:
	rm -rf $(OUTPUT)

setup:
	GO111MODULE=off \
		    go get -u github.com/onsi/ginkgo/ginkgo

.PHONY:	all clean setup test
