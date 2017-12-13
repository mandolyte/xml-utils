# UnmarshalAny
This utility will parse any XML file and create a CSV
formatted file. This can be useful to quickly understand
the contents of an XML file.

To show help:
```
$ go run unmarshalAny.go -help
Help Message

Usage: parseAny [options]
  -help
    	Show usage message
  -i string
    	Input XML filename
  -o string
    	Output CSV filename; default STDOUT
$ 
```

Example:
```
$ go run unmarshalAny.go -i test1.xml -o test1.csv
$ cat test1.csv
Depth,Type,Value
0,XMLName Local,Person
0,XMLName Space,
0,Comment,this is a comment
1,XMLName Local,FullName
1,XMLName Space,
1,CharData,Grace R. Emlin
1,XMLName Local,Company
1,XMLName Space,
1,CharData,Example Inc.
1,XMLName Local,Email
1,XMLName Space,
1,Attr Name,where
1,Attr value,home
2,XMLName Local,Addr
2,XMLName Space,
2,CharData,gre@example.com
1,XMLName Local,Email
1,XMLName Space,
1,Attr Name,where
1,Attr value,work
2,XMLName Local,Addr
2,XMLName Space,
2,CharData,gre@work.com
1,XMLName Local,Group
1,XMLName Space,
2,XMLName Local,Value
2,XMLName Space,
2,CharData,Friends
2,XMLName Local,Value
2,XMLName Space,
2,CharData,Squash
1,XMLName Local,City
1,XMLName Space,
1,CharData,Hanga Roa
1,XMLName Local,State
1,XMLName Space,
1,CharData,Easter Island
$ 
```

A performance test:
```
$ time go run unmarshalAny.go   -i $HOME/data/leading_causes_of_death_us.xml   -o $HOME/data/leading_causes_of_death_us.csv

real	0m5.715s
user	0m1.591s
sys	0m4.022s

$ cd $HOME/data
$ ls -al lea*
-rw-r--r-- 1 cecil cecil 9843826 Dec 13 17:01 leading_causes_of_death_us.csv
-rw-r--r-- 1 cecil cecil 5027058 Dec 11 07:35 leading_causes_of_death_us.xml
$ 
```