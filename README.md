# `spring-auto-reconfiguration-cnb`
The Cloud Foundry Spring Auto-reconfiguration Buildpack is a Cloud Native Buildpack V3 that provides Auto-reconfiguration functionality to Spring applications.

## Detection
The detection phase passes if:

* The build plan contains `jvm-application`
  * Contributes `auto-reconfiguration` to the build plan

## Build
If the build plan contains

* `auto-reconfiguration`
  * Checks for the existence of a `spring-core` jar in the application.
  * If found,
    * Contributes the Spring Auto-reconfiguration jar to a layer marked launch.
    * Adds the Spring Auto-reconfiguration jar to the classpath.

## License
This buildpack is released under version 2.0 of the [Apache License][a].

[a]: https://www.apache.org/licenses/LICENSE-2.0
