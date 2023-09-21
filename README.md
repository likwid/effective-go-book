# Effective Go Book

## Setup

This repository uses Nix to manage the dependencies of the small applications built from the book source code.

Using Nix allows us to unify the local developer's environment with CI/CD using the same package manager (Nix).

### flake.nix

This file represents the development environment we want for our application. The most interesting parts of flake.nix are found between lines 16 and 29, where we define a developer environment.  We use the mkShell function from Nix to instantiate a development environment using the packages (golang toolchain) defined on lines 20, 23 and 26.

### .github/workflows/nix.yml

This file uses DeterminateSystems Github Actions to install Nix and enable the Nix Magic Cache. This also demonstrates how the unification of local development environments and CI/CD can work.

On your local machine, or inside of Github Actions, you can use `nix develop --command go version` and it will use the `flake.nix` definitions to ensure that the two environments are the same.

