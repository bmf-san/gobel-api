# gobel-api
[![GitHub license](https://img.shields.io/github/license/bmf-san/gobel-api)](https://github.com/bmf-san/gobel-api/blob/master/LICENSE)
[![CircleCI](https://circleci.com/gh/bmf-san/gobel-api/tree/master.svg?style=svg)](https://circleci.com/gh/bmf-san/gobel-api/tree/master)

Gobel is a headless cms built with golang.

This is an api repository.

# gobel
- [gobel-api](https://github.com/bmf-san/gobel-api)
- [gobel-admin-client-example](https://github.com/bmf-san/gobel-admin-client-example)
- [gobel-client-example](https://github.com/bmf-san/gobel-client-example)
- [gobel-example](https://github.com/bmf-san/gobel-example)
- [gobel-ops-example](https://github.com/bmf-san/gobel-ops-example)

# Dockerhub
[bmfsan/gobel-api](https://hub.docker.com/r/bmfsan/gobel-api)

# Documentation
- [Coding Rule](https://github.com/bmf-san/gobel-api/blob/master/doc/CodingRule.md)
- [Database](https://github.com/bmf-san/gobel-api/blob/master/doc/database/README.md)
- [API](https://github.com/bmf-san/gobel-api/blob/master/doc/API.md)
- [Specification](https://github.com/bmf-san/gobel-api/blob/master/doc/Specification.md)
- [Talend API Tester](https://github.com/bmf-san/gobel-api/blob/master/doc/talend_api_tester.json)

# Get started
## Edit an env file
`cp .env_example .env`

##  Edit a host file
```
127.0.0.1 gobel-api.local
```

## Creaate a network
`docker network create --driver bridge gobel_link`

## Build containers
`make docker-compose-build`

## Run containers
```
make docker-compose-up
```

or

```
make docker-compose-up-d
```

Then go to `gobel-api.local:80`

# Architecture
gobel-api is based on Clean Architecture.

| Layer                | Directory      |
|----------------------|----------------|
| Frameworks & Drivers | infrastructure |
| Interface            | interfaces     |
| Usecase              | usecase        |
| Entities             | domain         |

cf. [blog.cleancoder.com - The Clean Code Blog by Robert C. Martin (Uncle Bob)](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)

# Contributing
We welcome your issue or pull request from everyone.
Please make sure to read the [CONTRIBUTING.md](https://github.com/bmf-san/gobel-api/.github/CONTRIBUTING.md).

# License
This project is licensed under the terms of the MIT license.

# Author
bmf - Software engineer.

- [github - bmf-san/bmf-san](https://github.com/bmf-san/bmf-san)
- [twitter - @bmf-san](https://twitter.com/bmf_san)
- [blog - bmf-tech](http://bmf-tech.com/)
