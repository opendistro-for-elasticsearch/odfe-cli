# Release guidelines


odfe-cli versions are expressed as **x.y.z**, where **x** is the major version, **y** is the minor version, and **z** is the patch version, following [Semantic Versioning](https://semver.org/) terminology. 

**The major, minor & patch version is related to odfe-cli compatibility, instead of odfe’s compatibility.**

 Note: We only backport applicable fixes, including security fixes to released branches, depending on severity and feasibility, everything else needs to go in next available minor release train.

AI: Action Item for change by plugin owners

**Q: What goes into new patch version release train?**

* Known bugs that prevent users from executing tasks. 
    * AI: Fix the bug, write/update unit test and integration test (if desired).
    * eg: Failed to delete non-running detectors by delete command, failed to handle error.
* Security vulnerabilities. 
    * AI: Fix code and merge.
    * eg: leaking user credentials, insecure interaction with components.
* cli’s dependency is updated with new patch version which fixes bugs and security issues that affects your feature. ( Treat your dependent library as your library)
    * AI: Updating your dependency in go.mod file to latest patch version. 
    * eg: cobra released a new version which fixes some of security vulnerabilities like [#1259](https://github.com/spf13/cobra/pull/1259)

The plugin owner, who merged commit to master, **related to above use case**, should notify admin to initiate new patch release. There is no SLA for patch releases since this is triggered by individual plugin owners.



**Q: What goes into new minor version release train?**

* Any new commands ( could be new plugin, new sub command from any plugin ). A new command could be new API that is release long ago but being included in new odfe-cli release or new API that will be  released in next odfe release.
    * AI: Merge PR, update documentation to reflect new changes.
    * eg: on-board [k-nn](https://github.com/opendistro-for-elasticsearch/odfe-cli/pull/20/) plugin , add auto-complete feature for cli commands.
* New parameter/ flags for any command.
    * AI: Merge PR, update documentation to reflect new changes.
    * eg: odfe-cli 1.0.0 only displays detector configuration using get command for given name, if user would also like to see [detector job](https://opendistro.github.io/for-elasticsearch-docs/docs/ad/api/#get-detector), they can add new flag (job) to get command to enable this feature.
* Any incompatible changes that was introduced with respect to API changes.
    * AI: Plugin owners should work on this change as early as possible once the change is introduced in their master branch, thus it gives ample time to work on incorporating this change in cli. We recommend plugin owners to plan this change as part of API migration step itself.
    * eg: if API is added in odfe 1.13.0 and updated in a backward incompatible way in later releases, this will be addressed in CLI as minor version

admin will trigger new minor release train whenever new plugin is on-boarded or every 45 days if any commits related to above use case is added. After every minor release, admin will create a new table where contributors can include candidates for next release.



**Q**: **What goes into new major version release train?**

* Any changes related to framework or addition of new component with respect to odfe-cli.
    * eg: Use new rest library (networking layer), framework re-design to ease new plugin on-boarding, user experience changes like update command formats, rename command names, etc. 

admin will trigger new major release train every 120 days if any commits related to above use case is added. After every major release, admin will create a new table related to next release where contributors can include candidates for next release.


