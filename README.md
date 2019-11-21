<p align="center">
    <img alt="Sloth Logo" src="https://raw.githubusercontent.com/Clivern/Sloth/master/assets/img/logo.png" width="130" />
    <h3 align="center">Sloth</h3>
    <p align="center">Highly Available Liteweight Monitoring & Alerting System.</p>
    <p align="center">
        <a href="https://godoc.org/github.com/clivern/sloth"><img src="https://godoc.org/github.com/clivern/sloth?status.svg"></a>
        <a href="https://travis-ci.org/Clivern/Sloth"><img src="https://travis-ci.org/Clivern/Sloth.svg?branch=master"></a>
        <a href="https://github.com/Clivern/Sloth/releases"><img src="https://img.shields.io/badge/Version-0.0.1-red.svg"></a>
        <a href="https://goreportcard.com/report/github.com/Clivern/Sloth"><img src="https://goreportcard.com/badge/github.com/Clivern/Sloth?v=0.0.1"></a>
        <a href="https://github.com/Clivern/Sloth/blob/master/LICENSE"><img src="https://img.shields.io/badge/LICENSE-MIT-orange.svg"></a>
    </p>
</p>

<p align="center">
    <img alt="Sloth Logo" src="https://raw.githubusercontent.com/Clivern/Sloth/feature/architecture/assets/img/sloth.png" height="450" />
</p>


Sloth is a highly available liteweight monitoring and alerting system. It comes with three roles, agent role that runs health checks on hosts, orchestrator role that watch the agents and split the work across workers and worker role that is responsible for updating agents state, health checks and alerting. 

Sloth uses RabbitMQ as default message broker and etcd as a data storage and for leader election (electing the orchestrator). It is still an experimental project so use at your own risk.


## Documentation

### Installation:

```golang
#
```


## Versioning

For transparency into our release cycle and in striving to maintain backward compatibility, Sloth is maintained under the [Semantic Versioning guidelines](https://semver.org/) and release process is predictable and business-friendly.

See the [Releases section of our GitHub project](https://github.com/clivern/sloth/releases) for changelogs for each release version of Sloth. It contains summaries of the most noteworthy changes made in each release.


## Bug tracker

If you have any suggestions, bug reports, or annoyances please report them to our issue tracker at https://github.com/clivern/sloth/issues


## Security Issues

If you discover a security vulnerability within Sloth, please send an email to [hello@clivern.com](mailto:hello@clivern.com)


## Contributing

We are an open source, community-driven project so please feel free to join us. see the [contributing guidelines](CONTRIBUTING.md) for more details.


## License

© 2019, Clivern. Released under [MIT License](https://opensource.org/licenses/mit-license.php).

**Sloth** is authored and maintained by [@Clivern](http://github.com/clivern).
