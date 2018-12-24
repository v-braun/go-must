# go-must
> simple go assertion package

By [v-braun - viktor-braun.de](https://viktor-braun.de).

[![](https://img.shields.io/github/license/v-braun/go-must.svg?style=flat-square)](https://github.com/v-braun/go-must/blob/master/LICENSE)
[![Build Status](https://img.shields.io/travis/v-braun/go-must.svg?style=flat-square)](https://travis-ci.org/v-braun/go-must)
[![codecov](https://codecov.io/gh/v-braun/go-must/branch/master/graph/badge.svg)](https://codecov.io/gh/v-braun/go-must)
![PR welcome](https://img.shields.io/badge/PR-welcome-green.svg?style=flat-square)




## Installation
```sh
go get github.com/v-braun/go-must
```



## Usage

``` go

// if val == nil will panic with "argument val is nil"
must.ArgNotNil(val, "val") 

// if val <= 1 will panic with val invalid: (0 > 1) == false
must.ArgBeValid(val > 1, "val invalid: (%d > 1) == false", val)

```




## Authors

![image](https://avatars3.githubusercontent.com/u/4738210?v=3&amp;s=50)  
[v-braun](https://github.com/v-braun/)



## Contributing

Make sure to read these guides before getting started:
- [Contribution Guidelines](https://github.com/v-braun/go-must/blob/master/CONTRIBUTING.md)
- [Code of Conduct](https://github.com/v-braun/go-must/blob/master/CODE_OF_CONDUCT.md)

## License
**go-must** is available under the MIT License. See [LICENSE](https://github.com/v-braun/go-must/blob/master/LICENSE) for details.
