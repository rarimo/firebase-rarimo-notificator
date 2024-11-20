# firebase-rarimo-notificator

## Description
 Firebase-Rarimo-Notificator is a robust server-side application designed to seamlessly integrate with the RariMe and United Space mobile applications, delivering real-time notifications to iOS and Android devices through Firebase Cloud Messaging (FCM). This project serves as the backbone for effective communication, ensuring users receive timely and relevant updates directly on their mobile devices.

## Install

  ```
  git clone gitlab.com/tokend/firebase-rarimo-notificator
  cd firebase-rarimo-notificator
  go build main.go
  export KV_VIPER_FILE=./config.yaml
  ./main run service
  ```

## Documentation

We do use openapi:json standard for API. We use swagger for documenting our API.

To open online documentation, go to [swagger editor](http://localhost:8080/swagger-editor/) here is how you can start it
```
  cd docs
  npm install
  npm start
```
To build documentation use `npm run build` command,
that will create open-api documentation in `web_deploy` folder.

To generate resources for Go models run `./generate.sh` script in root folder.
use `./generate.sh --help` to see all available options.

Note: if you are using Gitlab for building project `docs/spec/paths` folder must not be
empty, otherwise only `Build and Publish` job will be passed.  

## Running from docker 
  
Make sure that docker installed.

use `docker run ` with `-p 8080:80` to expose port 80 to 8080

  ```
  docker build -t gitlab.com/tokend/firebase-rarimo-notificator .
  docker run -e KV_VIPER_FILE=/config.yaml gitlab.com/tokend/firebase-rarimo-notificator
  ```

## Running from Source

* Set up environment value with config file path `KV_VIPER_FILE=./config.yaml`
* Provide valid config file
* Launch the service with `run service` command



### Third-party services


## Contact

Responsible 
The primary contact for this project is  [//]: # (TODO: place link to your telegram and email)
