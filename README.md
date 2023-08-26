# gobel-api
[![GitHub license](https://img.shields.io/github/license/bmf-san/gobel-api)](https://github.com/bmf-san/gobel-api/blob/master/LICENSE)
[![CircleCI](https://circleci.com/gh/bmf-san/gobel-api/tree/master.svg?style=svg)](https://circleci.com/gh/bmf-san/gobel-api/tree/master)
[![codecov](https://codecov.io/gh/bmf-san/gobel-api/branch/master/graph/badge.svg?token=HMEQ7EAUED)](https://codecov.io/gh/bmf-san/gobel-api)

The Gobel is a headless cms built with golang.

# gobel
- [gobel-api](https://github.com/bmf-san/gobel-api)
- [gobel-admin-client-example](https://github.com/bmf-san/gobel-admin-client-example)
- [gobel-client-example](https://github.com/bmf-san/gobel-client-example)
- [gobel-example](https://github.com/bmf-san/gobel-example)
- [gobel-ops-example](https://github.com/bmf-san/gobel-ops-example)
- [migrate-rubel-to-gobel](https://github.com/bmf-san/migrate-rubel-to-gobel)

# Dockerhub
[bmfsan/gobel-api](https://hub.docker.com/r/bmfsan/gobel-api)

# Features
- Support Go 1.21

# Documentation
- [Coding Rule](https://github.com/bmf-san/gobel-api/blob/master/doc/CodingRule.md)
- [Database](https://github.com/bmf-san/gobel-api/blob/master/doc/database/README.md)
- [API](https://github.com/bmf-san/gobel-api/blob/master/doc/API.md)
- [Specification](https://github.com/bmf-san/gobel-api/blob/master/doc/Specification.md)
- [Talend API Tester](https://github.com/bmf-san/gobel-api/blob/master/doc/gobel-api.json)

# Get started
## Edit an env file
`cp .env_example .env`

##  Edit a host file
```
127.0.0.1 gobel-api.local
```

## Create a network
`docker network create --driver bridge gobel_link`

## Build containers
`make docker-compose-build`

## Run containers
```
make docker-compose-up

or

make-docker-compose-up-d
```

Then go to `gobel-api.local`

## Run tests
|     command     |                            description                             |
| --------------- | ------------------------------------------------------------------ |
| make test       | Run unit tests                                                     |
| make test-cover | Run unit tests with cover options. ex. make test-cover OUT="c.out" |

# Architecture
gobel-api is based on Clean Architecture.

|        Layer         |   Directory    |
| -------------------- | -------------- |
| Frameworks & Drivers | infrastructure |
| Interface            | interfaces     |
| Usecase              | usecase        |
| Entities             | domain         |

cf. [bmf-tech.com - Golangでクリーンアーキテクチャに入門する](https://bmf-tech.com/posts/Golang%E3%81%A7%E3%82%AF%E3%83%AA%E3%83%BC%E3%83%B3%E3%82%A2%E3%83%BC%E3%82%AD%E3%83%86%E3%82%AF%E3%83%81%E3%83%A3%E3%81%AB%E5%85%A5%E9%96%80%E3%81%99%E3%82%8B)

# Contribution
Issues and Pull Requests are always welcome.

We would be happy to receive your contributions.

Please review the following documents before making a contribution.

[CODE_OF_CONDUCT](https://github.com/bmf-san/gobel-api/blob/master/.github/CODE_OF_CONDUCT.md)
[CONTRIBUTING](https://github.com/bmf-san/gobel-api/blob/master/.github/CONTRIBUTING.md)

# License
Based on the MIT License.

[LICENSE](https://github.com/bmf-san/gobel-api/blob/master/LICENSE)

## Author
[bmf-san](https://github.com/bmf-san)

- Email
  - bmf.infomation@gmail.com
- Blog
  - [bmf-tech.com](http://bmf-tech.com)
- Twitter
  - [bmf-san](https://twitter.com/bmf-san)