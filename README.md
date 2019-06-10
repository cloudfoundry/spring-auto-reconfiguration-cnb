# `spring-auto-reconfiguration-cnb`
The Spring Auto-reconfiguration Buildpack is a Cloud Native Buildpack V3 that provides Auto-reconfiguration functionality to Spring applications.

## Detection
The detection phase passes if:

* A `spring-core` jar exists in the application
  * Contributes `auto-reconfiguration` to the build plan

## Build
If the build plan contains

* `auto-reconfiguration`
  * Contributes the Spring Auto-reconfiguration jar to a layer marked launch.
  * Adds the Spring Auto-reconfiguration jar to the classpath.

## License
This buildpack is released under version 2.0 of the [Apache License][a].

[a]: http://www.apache.org/licenses/LICENSE-2.0
