# readFileEnv
A tiny, awful golang package for reading single-line environment configuration files.

## Using
* Import
```import github.com/hmschreck/readFileEnv```

* Create a directory reader
```configs := readFileEnv.MakeEnvFileReader("/path/to/config/dir")```

* Reading a configuration file
```database_ip := configs("database_ip")```
