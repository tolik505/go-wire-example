# Go Wire Examples

This repository offers several examples of how to use [wire](https://github.com/google/wire).
The examples start from the folder 1, which is the simplest example.
Each next folder contains more complex example based on the previous one.

- To play with it you have to install Wire by running:

```shell
go install github.com/google/wire/cmd/wire@latest
```

- To regenerate dependencies initialization you can run `wire` in terminal or use go generate feature.

- To have a proper behavior in IDE don't forget to add the build tag `-tags=wireinject`.

Have fun!
