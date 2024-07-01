```
░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░
░        ░░░      ░░░   ░░░  ░░░      ░░░        ░░░.-.=\-░░
▒▒▒▒▒▒  ▒▒▒  ▒▒▒▒  ▒▒    ▒▒  ▒▒  ▒▒▒▒  ▒▒▒▒▒  ▒▒▒▒(_)=='(_)▒
▓▓▓▓  ▓▓▓▓▓  ▓▓▓▓  ▓▓  ▓  ▓  ▓▓  ▓▓▓▓  ▓▓▓▓▓  ▓▓▓▓▓▓▓▓(▓▓▓▓▓
██  ███████  ████  ██  ██    ██        █████  █████████)████
█        ███      ███  ███   ██  ████  ██        ███████\_██
██████████████████████████████████████████████████████████)█
Welcome to Zonai ================================ ٩(╹ꇴ ╹๑)۶
```

Zonai is a tool that helps you dockerize your applications, from saving a project to loading it into internal server
The main idea is to serve system that without internet and heavily relies on internal server, so that they cannot use CI/CD.

This tool will help you to dockerize with `Dockerfile` automatically, or deploy the built projects into the server, changing `docker-compose.yml` for you.

_If you are using CI/CD, you probably don't nee this tool_

## Preview

You can either use prompt mode or command line mode. Here is the preview for both

#### Prompt Mode
![Prompt](https://github.com/Walter0697/Zonai/blob/fan/github/prompt.gif?raw=true)

#### Command Line Mode
![Cmd](https://github.com/Walter0697/Zonai/fan/github/cmd.gif?raw=true)

## Installation
To check which version of OS you need, use the following command
```bash
dpkg --print-architecture
```

After checking the OS version, you can build the binary version of the app using the following example command
```
GOOS=windows GOARCH=386 go build -o zonai_windows main.go
GOOS=linux GOARCH=amd64 go build -o zonai_amd64 main.go
```

For the below example, we will use the linux version of the binary
Move the binary to the directory you want, for the following example, we will move it to `/usr/local/bin`
```bash
    mv zonai_amd64 /usr/local/bin/zonai
```

Then, add the following line into `.bashrc` or `.zshrc` or `.profile` or `.bash_profile`
```bash
    alias zonai="/usr/local/bin/zonai"
```

Then, run the following command to make the changes take effect
```bash
    source ~/.bashrc
```

## Development
To run the app in development mode, you can use the following command
```bash
    go run main.go
```

## Usage
To use this tool, you will need to setup the configuration and project list.
You can choose either prompt mode for this tool by using `zonai` or using command line arguments. We will use command line arguments as the following example.

#### For Building Project
You should set up configuration for local development.
Using `zonai setup` you can setup different configuration.
For building project, you need to set `output-image-path`, `environment-path`. For M-series Mac User, you might need to change docker command for `docker-build-command`

Here is an example for setting one of the path
```bash
    zonai setup --output-image-path /path/to/your/project
```

After that, you will need to add a project using the following comand
```bash
    zonai add project your-project-name child-name /path/to/your/project
```

Then, you can build the project, usually we will build most version  using the following command
```bash
    zonai build --project-name -ac
```

#### For Deploying Project
You should set up configuration for local development.
Using `zonai setup` you can setup different configuration.
For deploying project, you can choose to set up `input-image-path` optionally if you decide to use prompt mode. However, with command line tools, you don't need to

Here is an example for setting one of the path
```bash
    zonai setup --input-image-path /path/to/your/project
```

After that, you will need to add a project using the following comand
```bash
    zonai add deployment your-project-name child-name /path/to/your/project
```

Then, you can deploy the project, usually we will deploy most version  using the following command
```bash
    zonai deploy /path/to/your/gzzipped_file
```

After that, to ensure everythings work, we won't deploy the project for you, you will still need to run `docker-compose up -d` since we want you to check to make sure things work perfectly.

## Completion

If you want to add completion for your shell, use the following commmand

`zonai completion zsh > /tmp/completion`
`source /tmp/completion`

For available shell, you can use the following command

`zonai completion`

## Author
Walter Cheng