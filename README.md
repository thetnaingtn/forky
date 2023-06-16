# forky
**forky** is a CLI tool that you can use to synchronize your forks with upstream repositories. You can select more than one repository at a time and synchronize them.\
To use **forky** you'll need to create a personal access token.
## How does `forky` work basically?
**forky** try to match the `default` branch(`main`, `master`, or `trunk` whatever it is) of your forked repository with the upstream repository's `same-named` branch to find how many commits behind by your forked repository are. Then it will show a list of forks which are one or more commits behind upstream repositories.
## Installation
### Mac OS
```sh
brew install thetnaingtn/tap/forky
```
then
```sh
forky --token `your github token`
```
## Usage
You can use following keys to interact with **forky**
* `a` to select all listed forks
* `n` to un-select all listed forks
* `space` to toggle the fork
* `r` to refresh the list
* `m` merge the selected fork with upstream repository(might change to `s`(which stand for synchronize) in future!)
* `q` or `ctrl+c` to quit
