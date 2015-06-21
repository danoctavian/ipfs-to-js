require('node-go-require')

console.log("loading go code...")

var mainGo = require("../main.go")


console.log("wtf man")
console.log(mainGo.ops.add(2, 5))
mainGo.ops.print("holy shit")

mainGo.ops.readFile("/home/dan/hello")