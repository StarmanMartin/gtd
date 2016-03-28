# gtd

gtd (go test deployment) is a quick local deployment tool

## Why gtd

*gtd* deploys your go packages to your local go path. If you have your own go packages outside your *GoPath* you can:

* add the package directory to your *GoPath* environment variable.
* install your package to your local system.

*gtd* helps you to easily install your packages to your local system if you do not want to add all package paths to your *GoPath*.

## Install gtd

```bash
go get github.com/starmanmaritn/gtd
```

To run gtd it is necessary to add the */bin* folder in your *GoPath* to your *Path* environment variable

## How to use gtd

Just enter *gtd* to your terminal or cmd. On your first run you need to configure gtd.
In order to do this you need to enter the following information:

* `GoPath` The *gopath* you want to deploy to
* `cwd (current working directory)` the directory all your projects are in
* `package prefix` this prefix is the package prefix like *github.com/starmanmartin* if you leave this field empty gtd copies the packages but not installs it
* `packagelist` a list with all the packages you want to deploy

## Commandline tags

### Reset

```bash
gtd -r
```

if you enter *gtd -r* you reset the settings


### Add package

```bash
gtd -a
```

if you enter *gtd -a* you can add a new package