# `spring-auto-reconfiguration-cnb`
The Cloud Foundry Spring Auto-reconfiguration Buildpack provides [Auto-reconfiguration][a] functionality to Spring applications.

[a]: https://github.com/cloudfoundry/java-buildpack-auto-reconfiguration#what-is-auto-reconfiguration

## Behavior
The buildpack will participate if all of the following conditions are met

* The application is a Java application
* A `spring-core` JAR exists in the application
* `BP_AUTO_RECONFIGURATION` is set to `true` (or unset)

The buildpack will do the following:
* Contribute the Spring Auto-reconfiguration JAR and add it to the classpath

## Configuration
| Environment Variable | Description
| -------------------- | -----------
| `$BP_AUTO_RECONFIGURATION_ENABLED` | Boolean value indicating whether this buildpack should participate.  Defaults to `true`.

## Detail
* **Provides**
  * `auto-reconfiguration`
* **Requires**
  * `auto-reconfiguration`
  * `jvm-application` 

## License
This buildpack is released under version 2.0 of the [Apache License][a].

[a]: https://www.apache.org/licenses/LICENSE-2.0
