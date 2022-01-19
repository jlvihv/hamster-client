# hamster-client
User client Implementation of hamster share


| technology    | version               |
| ------------- | ----------------------|
| go            | go1.17.2              |
| sqlite        | v1.2.3                |
| gorm          | v1.22.2               |
| wails         | v1.16.8               |
| vue.js        | 3.0.0                 |
| ant-design-vue| 2.2.8                 |

### wails Install

#### precondition
- Ubuntu
```
sudo apt install libgtk-3-dev libwebkit2gtk-4.0-dev
```
- MacOS
```
xcode-select --install
```

#### install command
```
go install github.com/wailsapp/wails/cmd/wails@latest
```

### project compilation
```
wails build
cd build
./hamster-client
```
### development and debugging
```
wails serve
cd frontend
npm run serve
```