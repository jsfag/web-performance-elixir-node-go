const http = require('http')

http.createServer((req, res) => {
  res.end('world')
}).listen(3000)