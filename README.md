# gitea-createRepo
a tiny go binary to create a repo on a gitea server

## usage

* set GITEA_TOKEN to a gitea API Token
* run `createRepo <repoName>` 
* `<repoName>`  should be in the form of "repository" without an organization, this will create a repo in the
account of the user who owns the API Token.

createRepo exits with an error level of 0 if the repository is created, 1 or 2 if it isn't.
