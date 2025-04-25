# Plugin Discovery

## Dependencies
- [Kong API gateway](https://github.com/Kong/kong) 

## How to build
```
make build
```

## How to run
```
make run
```

## Env variables

| Var             | Required|Default      |
| ----------------| ------- |-------------|
| KONG_HOST       | yes     |  -          |
| KONG_SCHEME     | http    |  http       |
| KONG_PLUGIN_TAG | no      |  pcm-plugin |

## Usage

get the list of plugins
```
curl http<s>://<host>:8080/plugins
```

## Contributing
State if you are open to contributions and what your requirements are for accepting them.

For people who want to make changes to your project, it's helpful to have some documentation on how to get started. Perhaps there is a script that they should run or some environment variables that they need to set. Make these steps explicit. These instructions could also be useful to your future self.

You can also document commands to lint the code or run tests. These steps help to ensure high code quality and reduce the likelihood that the changes inadvertently break something. Having instructions for running tests is especially helpful if it requires external setup, such as starting a Selenium server for testing in a browser.

## Authors and acknowledgment
Valerii Kalashnikov

## License
For open source projects, say how it is licensed.
