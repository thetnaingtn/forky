# forky
**forky** is a CLI tool that you can use to synchronize your forks with upstream repositories. You can select more than one repository at a time and synchronize them. To use **forky** you'll need to create a Github's personal access token.
## Demo
### Dark
![forky-dark](./forky-dark.gif)
### Light
![forky-light](./forky-light.gif)
## How does forky work basically?
**forky** detect the `default` branch(`main`, `master`, or `trunk` whatever it is) of your forked repository and will try to compare it with the upstream repository's `same-named` branch to find how many commits behind by your forked repositories are. Then it will show available forks which left behind the upstream repositories to synchronize.
## Installation
### Mac OS
```sh
brew install thetnaingtn/tap/forky
```
### npm
Are you a Node.js developer or Node.js is installed on your system? Perfect! **forky** also available as **npm** package
```sh
npm install -g forkyy
```
> Yes, you are right! The package name has an extra `y`. That is because the name `forky` is already taken by another package. 
## Usage
```sh
forky --token `your github token`
```
## Keymaps
You can use the following keys to interact with **forky**
| Key              | Description                                 |
|:-----------------|:--------------------------------------------|
| <kbd>a</kbd>     | Select all forks                            |
| <kbd>n</kbd>     | Select none of the forks                    |
| <kbd>space</kbd> | Toggle(select/unselect) the fork            |
| <kbd>r</kbd>     | Refresh                                     |
| <kbd>m</kbd>     | Merge the selected fork with upstream branch|
| <kbd>q</kbd>     | Quit                                        |
