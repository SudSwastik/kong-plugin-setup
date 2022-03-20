# Kong Golang Plugin Setup

> This plugin is an embedded server based plugin

> Create the plugin code 
```
$ touch <plugin-name>.go
```

> Create a multi stage Dockefile to build and install plugin in a custom kong image
```
$ docker build -t <custom-docker-image>:tag .
```

> Create a configuration file to test the plugin
```
$ touch kong.yml
``` 

> Run the Kong image to test the plugin
```
$ docker-compose up
$ curl --location --request GET 'http://localhost:8000/anything'
```

> Check the headers to confirm the plugin working

