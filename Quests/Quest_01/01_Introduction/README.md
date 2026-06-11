## Instruction

1- get-ready
Create in your Gitea account the repository named [b]*piscine-go*[/b].

This repository will be the folder where all the exercises must be uploaded.

Once created, clone that repository on your desktop.

To do so, open a Unix shell (e.g. Git Bash on Windows), you are going to type commands in it.

First, tell Git to remember your password (like a web browser would):

git config --global credential.helper store.

If your username was choumi this is the command that will need to be used:

git clone https://acad.learn2earn.ng/git/choumi/piscine-go.git
This command needs to be adapted with your own username.

2- set
Once the repository is created, use your code editor to write your first shell script called hello.sh

When executed, this script must print Hello choumi!, where choumi is your username.

Usage
If the username is choumi:

$ bash hello.sh
Hello choumi!
$
3- go-say-hello
After that the hello.sh is executing correctly, it needs to be uploaded to the repository with the following commands:

git add hello.sh

git commit -m "My very first commit"

git push

Once these steps are applied, the file can now be submitted for grading on the platform by clicking on the "RUN INTRODUCTION TEST" button.

This action will run the tests on your submitted hello.sh file.