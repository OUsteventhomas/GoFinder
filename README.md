This is a utility helping to search words from a docx file.
Our usecase situation:
  In our company we have outside sales reps who are drawing up contracts with an internal team tasked with making sure those contracts contain specific language. Instead of taking in search words as a parameter on program run, we needed to always look for the same set of words in a proposal to make sure there is coverage.
  

The first argument is the Utility name, GoFinder
The next is docx filename

The result is printed as
-------------------------------
true  C++ 

false Java

indicating, C++ word exists n sample.docx and Java does not.

docx is a compressed zip file, following Office Open XML specifcations.
It contains various xml files. It follows ECMA-376 specfication.
More details about ECMA are available here
http://www.ecma-international.org/publications/standards/Ecma-376.htm

All important data for a docx file is present in word/document.xml.
This utility looks through the document.xml for given words.

The utility does a case-senistive search
It uses mainly ioutil and zip package from golang
