# Architecture

# how to make changes

## changing graphql

see thunder ...

"github.com/samsarahq/thunder/graphql"

maybe rename folder zupfmanager to assetstore

assetconfig represents the interface to the config jsonn part of a zupfnoter abc file

cli is built with cobra github.com/spf13/cobra

go install github.com/spf13/cobra/cobra to install the cobra cli

## webui

grahql code gen erzeugt aus dem graphql introspection json die Typescript tyen

```
cd webui
yarn graphql
```

nach änderungen am Schema wieder ausfürhen (generiert auch das instrospection json siehe package.json)

Queries, welche von der Webui verwendet werden

queries werden vordefiniert und ggf. parametriesiert

Aus diesen Queries werde react - Komponenten - diese liegen unter generated

AssetComponent entsteht aus der Query namens Asset in queries/asset.graphql



