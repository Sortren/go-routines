# go-routines


## Description
This repo will help me and others better understand thread/go-routines synchronization and general operations.<br>
Every example is based on a simple app. File README.md in every cmd/<app_name> folder describes its purpose and requirements.

## Codebase schema
```sh 
├── apps              # apps source code
│   └── <app_name>    # structure defined as <app_name> folders
├── cmd               # entrypoint for every defined app
│   └── <app_name>    # 'main.go' for running the app
├── pkg               # modules that potentially could be exported 
│   ├── files         # files i/o tools
│   └── synchro       # synchronization tools
```
