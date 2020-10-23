Elemental is Golang skeleton app that using :
-
- Viper for configuration
- Echo for http library

Dependency management natively supported by Go Mod

Install dependencies using `go mod vendor` for folder separation  

DockerFile is also included. Base image using go 1.15

You can change config in config.json. Also don't forget to change expose port in docker build as well. 
