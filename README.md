# deploor

`deploor` is the poor's man git push deployment solution.


## How to use

1. Clone this repository

2. Init your new repository using this as the template
```
# git init --bare --template=deploor <environment>/<project-name>.git
git init --bare --template=deploor production/my-cool-project.git
```

3. Push to it (the repository must contain a Fabric script with the following call syntax
```
# fab <environment> deploy:<git_ref: branch or commit or tag>
fab production deploy:1.0
```

