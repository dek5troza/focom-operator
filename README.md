# focom-operator
// TODO(user): Add simple overview of use/purpose

## Description
// TODO(user): An in-depth paragraph about your project and overview of use

## Getting Started

### Prerequisites
- go version v1.22.0+
- docker version 17.03+.
- kubectl version v1.11.3+.
- Access to a Kubernetes v1.11.3+ cluster.
- oapi-codegen installed with: go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
- setup-env test: go install sigs.k8s.io/controller-runtime/tools/setup-envtest@latest
- export KUBEBUILDER_ASSETS=$(pwd)/testbin/k8s/1.26.0-darwin-arm64 for the above

### To Deploy on the cluster
**Build and push your image to the location specified by `IMG`:**

```sh
make docker-build docker-push IMG=<some-registry>/focom-operator:tag
```

**NOTE:** This image ought to be published in the personal registry you specified.
And it is required to have access to pull the image from the working environment.
Make sure you have the proper permission to the registry if the above commands don’t work.

**Install the CRDs into the cluster:**

```sh
make install
```

**Deploy the Manager to the cluster with the image specified by `IMG`:**

```sh
make deploy IMG=<some-registry>/focom-operator:tag
```

> **NOTE**: If you encounter RBAC errors, you may need to grant yourself cluster-admin
privileges or be logged in as admin.

**Create instances of your solution**
You can apply the samples (examples) from the config/sample:

```sh
kubectl apply -k config/samples/
```

>**NOTE**: Ensure that the samples has default values to test it out.

### To Uninstall
**Delete the instances (CRs) from the cluster:**

```sh
kubectl delete -k config/samples/
```

**Delete the APIs(CRDs) from the cluster:**

```sh
make uninstall
```

**UnDeploy the controller from the cluster:**

```sh
make undeploy
```

## Current reconcilation method

1. The Reconcile method as is now, reads like a high-level “table of contents”: 
```
fetch → handle deletion → ensure finalizer → validate → build remote → ensure remote.
```
2.All the details of how finalizers or remote creation works are delegated to smaller helper methods that do one job.
3. Example Reconcile Flow

    Example flow:
    ```
    Fetch CR → if not found, done.
    If being deleted → handleDeletion(...) → done.
    If no finalizer → add finalizer → requeue.
    Validate → if invalid, set status → done.
    Build remote → if fails, set status → done.
    Ensure remote resource (create or poll) → set status → requeue if needed
   ```

## Notes
    - Code needs testcases and better refactoring of certain parts.
    - Code needs to ensure that templateParameters are immutable, this could be done with validating webhook.
    - Code needs to ensure template validation is performed, not just check for template existence.
    - Possibly wrong group for the TemplateInfo?


## Project Distribution

Following are the steps to build the installer and distribute this project to users.

1. Build the installer for the image built and published in the registry:

```sh
make build-installer IMG=<some-registry>/focom-operator:tag
```

NOTE: The makefile target mentioned above generates an 'install.yaml'
file in the dist directory. This file contains all the resources built
with Kustomize, which are necessary to install this project without
its dependencies.

2. Using the installer

Users can just run kubectl apply -f <URL for YAML BUNDLE> to install the project, i.e.:

```sh
kubectl apply -f https://raw.githubusercontent.com/<org>/focom-operator/<tag or branch>/dist/install.yaml
```

## Contributing
// TODO(user): Add detailed information on how you would like others to contribute to this project

**NOTE:** Run `make help` for more information on all potential `make` targets

More information can be found via the [Kubebuilder Documentation](https://book.kubebuilder.io/introduction.html)

## License

Copyright 2025.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

