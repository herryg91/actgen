# actgen
This package will generate user activity datas toward items into file. This package will be used for me to create data for data processing pipeline (batch processing).

Batched Datas (this package) -> Spark (Processing) -> Cassandra (Store)

flag:
--fname=string [default: default] | Location + filename where the data to be saved
--totalDatas=int [default: 10000] | Total rows to be generated to file
--user=int [default: 50] | Total user id your platform have
--item=int [default: 1000] | Total item id your platform have
--action=string [default: atw,watch,comment] | List of Action user did

Result Example
./actgen --totalDatas=10000 --user=50 --item=1000

```
12,76,atw
12,371,watch
31,119,atw
7,535,watch
19,855,atw
36,288,comment
42,151,comment
11,112,watch
13,251,comment
28,190,atw
41,986,atw
25,721,atw
31,533,atw
20,571,atw
43,715,atw
41,297,atw
5,83,watch
6,379,watch
45,335,watch
31,121,comment
```
