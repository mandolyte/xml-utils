# ParseAny
This utility will parse any XML file and create a CSV
formatted file. This can be useful to quickly understand
the contents of an XML file.

To show help:
```
$ go run parseAny.go -help
Help Message

Usage: parseAny [options]
  -help
    	Show usage message
  -i string
    	Input XML filename
  -maxattr int
    	Maximum number of attributes for an element; default 5 (default 5)
  -o string
    	Output CSV filename; default STDOUT
$ 
```

Example:
```
$ go run parseAny.go -i test1.xml -o test1.csv -maxattr 1
$ cat test1.csv 
Depth,Type,Name,Text,Attribute 1,Value 1
0,ProcInst,target,instructions
0,Directive,this-is-a-directive
0,Start,Person,
1,Start,FullName,
2,CharData,FullName,Grace R. Emlin
1,End,FullName
1,Start,Company,
2,CharData,Company,Example Inc.
1,End,Company
1,Start,Email,,where,home
2,Start,Addr,
3,CharData,Addr,gre@example.com
2,End,Addr
1,End,Email
1,Start,Email,,where,work
2,Start,Addr,
3,CharData,Addr,gre@work.com
2,End,Addr
1,End,Email
1,Start,Group,
2,Start,Value,
3,CharData,Value,Friends
2,End,Value
2,Start,Value,
3,CharData,Value,Squash
2,End,Value
1,End,Group
1,Start,City,
2,CharData,City,Hanga Roa
1,End,City
1,Start,State,
2,CharData,State,Easter Island
1,End,State
1,Comment,this is a comment
0,End,Person
$ 
```