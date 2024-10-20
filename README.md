# hookup
No strings attached!
Your friendly git hooks buddy!

Git hooks manager CLI tool

## Introduction
Git hooks automate and implement processes in your workflow, increasing code quality and consistency.

However, many developers avoid git hooks due to a lack of awareness and the perceived complexity of setup, discouraging them from using this feature.

**hookup** simplifies the management of git hooks, allowing you to focus on the logic and usage of your hooks instead.

## Installation
Instructions on how to install the tool.
#### External Dependencies
- `git`
- `curl` (if installing via script)

## Usage Quickstart
This section provides basic information to understand the core usage.

> [!NOTE]
> Run `hu --help` for full usage documentation.

#### hu init
To initialize, run `hu init` in your current directory.

On the first run, it will create the **.hookup** folder and set the local **core.hooksPath** variable. If the **.hookup folder** already exists, only the path variable will be updated.

The **hooksPath** variable is set as a relative path, ensuring that moving the repository to another location won’t affect the sourcing of the .hookup folder.

#### hu add & hu remove


#### hu update


## Configuration
Information on any configuration files or settings that users may need to know about.

## Future TODOs
- [ ] branch specific hooks
- [ ] store custom git hooks as templates for future use
- [ ] functionality to save custom setups (ie gitdir and workdir are not in same location)
- [ ] github actions maybe

## Contributing
Guidelines for contributing to the project, including how to submit issues or pull requests.
