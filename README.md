# web-performance-elixir-node-go
### Elixir (Cowboy) vs. Native Node.js vs. Native Golang
## Setup
### Elixir:
```
mix deps.get
iex -S mix
```
### Node.js
```
node main.js
```
### Golang
```
go build main.go
./main
```
## Results
> Core2Duo e8500 4GB
```
wrk -t100 -c500 -d60s http://127.0.0.1:3000
```
### Elixir (Cowboy)
```
Running 1m test @ http://127.0.0.1:3000
  100 threads and 500 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   208.00ms  379.62ms   1.90s    83.83%
    Req/Sec   174.54    123.37     1.92k    66.07%
  967971 requests in 1.00m, 136.80MB read
  Socket errors: connect 0, read 16, write 0, timeout 213
Requests/sec:  16106.05
Transfer/sec:      2.28MB
```

### Node.js
```
Running 1m test @ http://127.0.0.1:3000
  100 threads and 500 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    21.43ms    2.10ms 266.12ms   98.18%
    Req/Sec   234.26     49.52    10.61k    98.91%
  1395441 requests in 1.00m, 138.40MB read
Requests/sec:  23219.00
Transfer/sec:      2.30MB
```

### Golang
```
Running 1m test @ http://127.0.0.1:3000
  100 threads and 500 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   217.54ms  365.31ms   2.00s    84.79%
    Req/Sec   586.42      1.95k   19.78k    93.80%
  2219004 requests in 1.00m, 256.06MB read
  Socket errors: connect 0, read 0, write 0, timeout 1424
Requests/sec:  36931.45
Transfer/sec:      4.26MB
```

```
wrk -t100 -c100 -d60s http://127.0.0.1:3000
```
### Elixir (Cowboy)
```
Running 1m test @ http://127.0.0.1:3000
  100 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    15.63ms   21.94ms 157.82ms   89.86%
    Req/Sec   167.86     60.34     1.52k    49.92%
  997178 requests in 1.00m, 140.93MB read
Requests/sec:  16604.10
Transfer/sec:      2.35MB
```

### Node.js
```
Running 1m test @ http://127.0.0.1:3000
  100 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     3.94ms  493.88us  24.95ms   98.35%
    Req/Sec   255.40     43.27     7.37k    98.94%
  1526803 requests in 1.00m, 151.43MB read
Requests/sec:  25404.54
Transfer/sec:      2.52MB

```

### Golang
```
Running 1m test @ http://127.0.0.1:3000
  100 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    45.39ms   83.46ms   1.05s    88.09%
    Req/Sec   440.95      1.49k   19.12k    93.69%
  2278721 requests in 1.00m, 262.95MB read
Requests/sec:  37937.35
Transfer/sec:      4.38MB
```
