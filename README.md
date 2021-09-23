# MarkdownToHtml
A [Go](https://golang.org) function that allow you to transform a Markdown text into HTML
> Tiny utility that I needed to create myself because I coudn't find another ready to use library.


## Supported tags:
| Markdown           | HTML                                   |
| ------------------ | -------------------------------------- |
| **                 | &lt;strong&gt;                         |
| __                 | &lt;em&gt;                             |
| *                  | &lt;b&gt;                              |
| _                  | &lt;i&gt;                              |
| #...               | &lt;h_&gt;                             |
| \[anything\](link) | &lt;a href="link"&gt;anything&lt;a&gt; |
| \[text\](link)     | &lt;img str="link" alt="text"&gt;      |
| \-\-\-             | &lt;hr&gt;                             |

> <h_> tags are supported from 1 to 7 but you must to go to the next line to close the tag

