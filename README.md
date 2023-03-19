# [justmark0.me](https://justmark0.me) backend repo

This is repo for my personal site.


### How to use it.
If you want the same site. You should fork it then configure secrets of repo, add github runner and start pipeline ☺️ !

## Local development
1. You can start development container via `docker-compose up` and need to run all migrations.
2. You will need to import env variables. You can specify them in `.env` and import them via `export $(grep -v '^#' .env | xargs)`
3. Then you can run app (you need to build `run go run src/main.go`)

### Pre commit hook
    pip3 install pre-commit && pre-commit install


### Contributions
Contributions are welcome! You can suggest them in github, also if needed you can contact me, my active contacts are in my github profile 😊
