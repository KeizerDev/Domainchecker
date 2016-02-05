[![License][License-Image]][License-Url] [![ReportCard][ReportCard-Image]][ReportCard-Url] [![Build][Build-Status-Image]][Build-Status-Url] [![Release][Release-Image]][Release-Url]
# Domainchecker
Domain checking from the terminal at your favorite supplier. Just opens your default browser when choosed.

```
Usage:
  s <query> [flags]

Flags:
  -b, --binary string     binary to launch search uri
  -l, --list-providers    list supported providers
  -p, --provider string   set search provider (default "google")
  -v, --verbose           display url when opening
      --version           display version
```

## Install

```
go get -v github.com/KeizerDev/domainchecker
cd $GOPATH/src/github.com/KeizerDev/domainchecker
make
make install
```

## Examples

Try example.com on godaddy.com.
```
domainchecker example.com
```

## Provider Expansion

Just use your preffered supplier using the `-p` tag.
```
domainchecker example.com -p namecheap
```

**Todo:**

You can also change the default provider in your domainchecker config file like this: 
```
default: namecheap.com
```

## Supported Providers

* godaddy
* transip

#### Contributors

* [Robert-Jan Keizer (KeizerDev)](https://github.com/KeizerDev/)

#### License

s is released under the MIT license.

[License-Url]: http://opensource.org/licenses/MIT
[License-Image]: https://img.shields.io/npm/l/express.svg
[Release-Url]: https://github.com/KeizerDev/domainchecker/releases/tag/v0.0.1
[Release-image]: http://img.shields.io/badge/release-v0.2.1-1eb0fc.svg
