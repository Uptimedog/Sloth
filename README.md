<p align="center">
    <img alt="Sloth Logo" src="https://raw.githubusercontent.com/clivern/sloth/master/assets/img/gopher.png?v=0.0.1" width="150" />
    <h3 align="center">Sloth</h3>
    <p align="center">Liteweight HA Monitoring & Alerting System.</p>
    <p align="center">
        <a href="https://godoc.org/github.com/clivern/sloth"><img src="https://godoc.org/github.com/clivern/sloth?status.svg"></a>
        <a href="https://travis-ci.org/clivern/Sloth"><img src="https://travis-ci.org/clivern/Sloth.svg?branch=master"></a>
        <a href="https://github.com/clivern/Sloth/releases"><img src="https://img.shields.io/badge/Version-0.0.1-red.svg"></a>
        <a href="https://goreportcard.com/report/github.com/clivern/Sloth"><img src="https://goreportcard.com/badge/github.com/clivern/Sloth?v=0.0.1"></a>
        <a href="https://github.com/clivern/Sloth/blob/master/LICENSE"><img src="https://img.shields.io/badge/LICENSE-MIT-orange.svg"></a>
    </p>
</p>

<p align="center">
    <img alt="Sloth Logo" src="https://raw.githubusercontent.com/clivern/Sloth/master/assets/img/diagram.png?v=0.1.0" width="65%" />
</p>


**Sloth** is a liteweight highly available monitoring and alerting system. It comes with three roles:
 * `Agent Role`: Runs health checks on hosts.
 * `Worker Role`: Updates agents state, health checks and alerting. The elected leader watches if agents up or down.
 * `API Role`: A REST service to interact with the cluster & collect metrics.

**Sloth** uses RabbitMQ as default message broker and etcd as a data storage and for leader election (electing the leader worker). It is still an experimental project so use at your own risk.


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

© 2019, clivern. Released under [MIT License](https://opensource.org/licenses/mit-license.php).

**Sloth** is authored and maintained by [@clivern](http://github.com/clivern).
