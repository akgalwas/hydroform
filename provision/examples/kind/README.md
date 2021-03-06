# Provision a Kind cluster

## Overview

This example shows you how to use Hydroform to provision a Kind cluster.
## Prerequisites
To provision a Kind cluster you need Docker with memory settings configured to support the cluster you want to run.
## Installation

### Run the example

1. To provision a new cluster on Azure, go to the `provision` directory and run:

    ```bash
    go run ./examples/kind/main.go -p {project-name} -n {dockerhub/image:tag} --persist
    ```

2. If you have Kind installed, you can run `kind get clusters` to see if your cluster is running.

3. Export the **KUBECONFIG** environment variable pointing to the `kubeconfig` file generated by running the example. This will allow you to access the cluster.

    ```bash
    export KUBECONFIG=$(pwd)/kubeconfig.yaml
    ```
