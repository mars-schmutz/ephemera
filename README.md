# Ephemera

This is a cli tool I developed to help grade assignments for the Software Development
program I TA'd for. It works with .zip, .tar.gz, and git repositories.
It creates a temporary directory and starts a new shell instance.
You can install dependencies, run the project, do whatever it is needed to grade the
project. Once you're done, you can type `exit` to exit the new shell instance
and it will also delete the temporary directory and project files.

## Installation

Make sure you install Go, then clone the repository and run `go build`.

## Use

![ephemera use](https://github.com/mars-schmutz/ephemera/blob/8fd0bb1164fffb608f74e35d26b485bb44a88658/imgs/tool.png "Help example")
This is an example of what the tool looks like when no flags are supplied. So far you need to either give it a repository link or
a path to an archive.

![using ephemera on a repository](https://github.com/mars-schmutz/ephemera/blob/8fd0bb1164fffb608f74e35d26b485bb44a88658/imgs/repo.png "Example with a repository")
Example usage with a repository link.

![exiting temporary shell](https://github.com/mars-schmutz/ephemera/blob/8fd0bb1164fffb608f74e35d26b485bb44a88658/imgs/exit.png "Exiting the temporary shell")
Exiting the tool.
