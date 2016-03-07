<h1 align="center">Domainchecker</h1>

<p align="center">
Check domain availability from your terminal and open your favorite supplier with the choosed domain. Just from your terminal.
</p>

<p align="center">
    <a href="http://opensource.org/licenses/MIT">
        <img src="https://img.shields.io/npm/l/express.svg">
    </a>    
    <a href="https://github.com/KeizerDev/domainchecker/releases/tag/v0.0.1">
        <img src="http://img.shields.io/badge/release-v0.0.1-1eb0fc.svg">
    </a>
</p>

----


```
Usage:
  domainchecker <query> [flags]

Flags:
  -l, --list-providers    list supported providers
  -p, --provider string   set doamin supply provider (default "godaddy") (TODO)
  -v, --verbose           display url when opening
      --version           display version
```

## What should it do?
It should check which domain is available and which not. 
When the domain is available it would be ideal to pass it to a domain supplier like [godaddy](https://godaddy.com/) (this is WIP now).


## Install

```
go get -v github.com/KeizerDev/domainchecker
cd $GOPATH/src/github.com/KeizerDev/domainchecker
make
make install
```

## Examples

To get a list of a specific domain just add an `.*` to the and of your name. This will list some domains who are available. 
```
domainchecker myepicname.*
```

Or just specify the name like so.
```
domainchecker myepicname.nl
```

**Todo !THIS STILL DOESN'T WORK!:**   
Add a provider flag to check the particular domain on a supplier site.
```
domainchecker myepicname.nl -p godaddy
```

To get a list of the suppliers just run.
```
domainchecker -l
```


## Supported Providers

* **TODO**

#### Contributors

* [Robert-Jan Keizer (KeizerDev)](https://github.com/KeizerDev/)

#### License

Domainchecker is released under the MIT license.
