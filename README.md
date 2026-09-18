# push

Push the current Git repository to all distinct configured push remotes.

## Installation

```sh
go install github.com/h5law/push@latest
````

## Usage

Run `push` from inside any Git repository:

```sh
push
```

The command discovers all configured remotes and their push URLs, removes duplicates, and pushes to each unique destination.

For example:

```text
origin  → git@github.com:h5law/project.git
gitlab  → git@gitlab.com:h5law/project.git
mirror  → git@github.com:h5law/project.git
```

Running `push` will push to:

```text
git@github.com:h5law/project.git
git@gitlab.com:h5law/project.git
```

The duplicate GitHub URL is only pushed once.
