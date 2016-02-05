<h1 align="center">Domainchecker</h1>

<p align="center">
Domain checking from the terminal at your favorite supplier. Just opens your default browser when choosed.
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
  -b, --binary string     binary to launch search uri
  -l, --list-providers    list supported providers
  -p, --provider string   set search provider (default "google")
  -v, --verbose           display url when opening
      --version           display version
```

## What should it do?
You should get a list of names



## Install

```
go get -v github.com/KeizerDev/domainchecker
cd $GOPATH/src/github.com/KeizerDev/domainchecker
make
make install
```

## Examples

Try the `.*` as domain extension to list the available domain extension for a specific name.
```
domainchecker myepicname.*
```

Use the provider flag to checkout an available domain on suppliers website.
```
domainchecker myepicname.nl -p godaddy
```

**Todo:**

You can also change the default provider in your domainchecker config file like this: 
```
default: namecheap.com
```

## Supported Providers

* transip

#### Contributors

* [Robert-Jan Keizer (KeizerDev)](https://github.com/KeizerDev/)

#### License

Domainchecker is released under the MIT license.