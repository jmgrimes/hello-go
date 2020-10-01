# Hello Go

Hello Go is a simple Gorilla/MUX application that says hello to a REST 
API caller through the /hello-world resource.  It serves as an example 
of how to build out an application with the Gorilla/MUX library, create 
a Docker container for its deployment, and then automate the gitflow 
release cycle with Github, Github Actions, and Github Packages.

## Workflows

The workflows for the CI/CD pipeline are contained within the 
.github/workflows folder.  There is currently a single workflow which 
handles building the application whenever a push occurs to the develop or 
master branches, as well as to any feature, release, hotfix, and support 
branches.  Finally, it triggers the build whenever a release is published.

## Build Process

Of special note, this project uses a multi-stage Docker build, pulling a 
Golang SDK container, building the binary, then starting from a fresh 
Alpine image, copies only the resulting binary to the runtime image.  This 
results in a build image of nearly 350MB being reduced to a mere 13MB, 
and allows an easy upgrade path for the build and runtime containers.

## Configuration

When running the application, it expects to find a yaml file at the 
/app/conf/application.yaml location.  The file must contain the following 
contents.

```yaml
template: Hello, %s! # The format for the hello message, must include a %s for the name substitution.
defaultName: Stranger # The name to use when a name is not specified in the "name" query parameter.
version: "1.0.0" # The version of the container to use in the response.
```