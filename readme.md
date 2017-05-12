[![Build Status](https://travis-ci.org/stevenbraham/gitil.svg?branch=master)](https://travis-ci.org/stevenbraham/gitil)
# About

Gitil (a contraction of Git and utils) is a set of tools that I have developed to simplify some common tasks I often have todo in Git such as merging a branch into master our creating a tag and pushing the tag. Originialy some of these tools where python scripts, but I have decided to bundle them into a singular CLI program. Everyone is free to use and extend my program, however there **aren't** are a lot of error handeling checks build in, therefore I don't recommend using this for mission critcial enterprise projects as it can really mess up your git repo's.

# Why Go?

Like I said I wanted to bundle my python scripts in a cross platform CLI tool. C and C++ where to difficult and too verbose for me. I was learning Go as a side project and decided that this would be an excellent test case for trying to improve my Go skills.

# How to use

Grab a binary from the [latest release](https://github.com/stevenbraham/gitil/releases) and add it to your $path. You can also clone this repo and run the makefile. After instalation run `gitil help`.

# Supported functions

* Single command merging a branch to master
* Single command mering master into a branch
* Single command merging master in all brances
* Git cloning
* Single commmand rest hard
* Create a gitignore from [gitignore.io](https://www.gitignore.io/) from the command line
* Single commnad fetch all
* Single command add all, commmit and push
* Single command create tag and push

Tip run `gitil help` to see all commands.

# Todo list
* Automaticly clonening, creating a gitignore, creating a v0.1.0 tag and pushing it all
* Way better error handling
* A plugin system
* Plugins for my IDE's and Editors
