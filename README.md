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

![Alt text](./images/Plugin_Header_Test.png?raw=true "Plugin Test")


> Configuration Change to Test Custom Plugin in Kong Ingress Controller

- Download the KIC installation yml
- Update the kong image with the custom image containing plugin code.
- Add the environment name:value to start the plugins  

```
$ image: <custom kong image>
$ 
- name: KONG_PLUGINS
  value: bundled,go-setheader
- name: KONG_PLUGINSERVER_NAMES
  value: go-setheader,go-hello
- name: KONG_PLUGINSERVER_GO_SETHEADER_QUERY_CMD
  value: /usr/local/bin/go-setheader -dump
```

```
$ kubectl apply -f ./resources/kic-custom-plugin-installation.yml
```

> Create the k8s resources to test the plugin 

- create the backend resources
```
$ kubectl apply -f ./resources/echo.yml
```
- create the custom plugin
```
$ kubectl apply -f ./resources/go-setheader-plugin.yml
```
- create the ingress and add plugin to it
```
$ kubectl apply -f ./resources/echo-ingress.yml
```

> Test the ingress path
```
$ curl -i localhost:80/foo
```

> A successful header with the name "go-config-message" means plugin is applied succesfully