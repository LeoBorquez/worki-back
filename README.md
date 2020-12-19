# Worki

## Init Golang project
go mod init github.com/user/project-name
git init 
git add .
git commit -m "initial commit"
git remote add origin git@github.com:user/project-name.git
git push -u origin master

## Folder Structure

my-app
├── server                   # Go project files
│   ├── db                   # MongoDB communications
│   ├── model                # domain objects
│   ├── web                  # REST APIs, web server
│   ├── server.go            # the starting point of the server
│   └── go.mod               # server dependencies
