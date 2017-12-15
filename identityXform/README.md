# IdentityXform
This utility will do an identity transformation, transforming
the input XML document to an equivalent output XML document.

Limitations:
- Neither Processing Instructions or Directives are supported.
These can be in the input, but will not be output.
- Comments may change position, but will be correctly associated
to the "parent" element.

There is an option to output in indented format, making this a
"pretty print" option. Strictly speaking, this would not
be an identity transformation.

To show help:
```
$ go run identityXform.go -help
Help Message

Usage: identityXform [options]
  -help
    	Show usage message
  -i string
    	Input XML filename
  -indent
    	Use indented format (pretty print); default is false
  -o string
    	Output XML filename; default STDOUT
$
```

## Using indented format
The input XML is a single line:
```
$ cat test1.xml 
<Person type="alien" thumbs="yes">here is some character data<FullName>Grace R. Emlin</FullName><!-- this is not a comment... just kidding --><Company>Example Inc.<!-- this company excels --></Company></Person>
$ 
```

Pretty print it:
```
$ go run identityXform.go -i test1.xml -indent
<Person type="alien" thumbs="yes">here is some character data
    <!-- this is not a comment... just kidding -->
    <FullName>Grace R. Emlin</FullName>
    <Company>Example Inc.
        <!-- this company excels --></Company>
</Person>
$
```

## Identity transformation
Notice that the comment changes position, but otherwise the 
transformation is equivalent.
```
$ go run identityXform.go -i test1.xml 
<Person type="alien" thumbs="yes">here is some character data<!-- this is not a comment... just kidding --><FullName>Grace R. Emlin</FullName><Company>Example Inc.<!-- this company excels --></Company></Person>
$ 
```

