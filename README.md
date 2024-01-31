### Pasta

Simple clipboard manager

Demo

![Pasta Demo](https://github.com/sudipidus/pasta/blob/main/pasta.gif)


### Instructions
1. Install the app 

```
git clone git@github.com:sudipidus/pasta.git
cd pasta
go build -o pasta
```

2. run pasta

```
./pasta
```

3. Add applescript to add as app or simple run. (This sends SIGUSER1 to the program upon listening which clipboard history is printed. You can trigger this script using a keyboard shortcut or by adding it to applications)

```
sh sender.sh
```


### What did I learn?
1. SIGUSR1, SIGUSR2: These two syscalls are reserved for users/applications in UNIX. You can write a program which listens to these and does certain tasks. (akin to SIGINT, SIGKILL but it won't actually kill the program)

2. You can convert simple scripts into mac Applications by adding them as script and saving them as application