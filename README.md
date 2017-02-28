 This is a utility helping to search words from a docx file.

Our usecase situation:
  In our company we have outside sales reps who are drawing up contracts with an internal team tasked with making sure those contracts contain specific language. Instead of taking in search words as a parameter on program run, we needed to always look for the same set of words in a proposal to make sure there is coverage.
  

Command:
>GoFinder.exe "C:\users\joe.smith\desktop\check this document.docx"

The result is printed as
-------------------------------
true    expiration

false   confidentiality

true	background

true	pricing

true 	executive summary

false	assumptions

true	scope of services

true	travel

true	schedule


These indicate by using a true/false if the words exist in the docx file.

docx is a compressed zip file, following Office Open XML specifcations.
It contains various xml files. It follows ECMA-376 specfication.
More details about ECMA are available here
http://www.ecma-international.org/publications/standards/Ecma-376.htm

All important data for a docx file is present in word/document.xml.
This utility looks through the document.xml for given words.

Since this utility changes all text to lowercase in the file, case sensitivity isn't an issue.
