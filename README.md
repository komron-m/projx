## Getting started

```shell
# clone repo
mkdir -p $GOPATH/src/github.com/komron-m/
cd $GOPATH/src/github.com/komron-m/
git clone git@github.com:komron-m/projx.git
cd projx

# copy and change env variables if needed
cp .env.example .env

# install tools and app dependencies
make install

# build fresh docker images and run services
make fresh

# apply migrations
make migrate
```

## Contributing

Pull requests are welcome. For any changes, please open an issue first to discuss what you would like to change.
