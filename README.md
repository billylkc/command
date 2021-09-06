# command
collection of random command line functions for me to check the news, updates from Data science blogs, twitter, etc..

| Group   | Command           | Long                         | Shortcut | Remark                                                 |
|---------|-------------------|------------------------------|----------|--------------------------------------------------------|
| News    | Google Trend      | command news gtrend          | cmd n tr | Read latest google trend from specific country         |
|         | Hacker News       | command news hackernews      | cmd n hn | Read Hackernews                                        |
|         | KD nuggents       | command news kdnuggets       | cmd n kd | Read KD nuggets post by month                          |
|         | Reddit            | command news reddit          | cmd n r  | Read sub reddit                                        |
|         | Twitter           | command news twitter         | cmd n tr | Live streaming Twitter post by tags (Need API account) |
|         |                   |                              |          |                                                        |
| Utility | Cheatsheet        | command utility cheatsheet   | cmd u ch | curl cht.sh for quick programming examples             |
|         | Strip HTML        | command utility stripHtml    | cmd u sh | (TBA)                                                  |
|         | Strip Line Number | command utility stripLineNum | cmd u sl | (TBA)                                                  |
|         | Git tag           | command utility tag          | cmd u t  | (TBA)                                                  |
|         | Wiki              | command utility wiki         | cmd u w  | Grep top 3 paragraph from wiki of the provided query   |
|         |                   |                              |          |                                                        |



# Installation
I would **NOT RECOMMEND** downloading any executable to your system directly. A better way would be compiling from the source code with your Go environment.

Use the following command at you own risk. It will download the executable to your binary directory

```
cd ~/bin
wget https://github.com/billylkc/command/blob/master/bin/command?raw=true
```

and call it with `command`


---
# Demo
![command](demo.gif)



---
# Alias

You can always set linux/windows alias to call the command with less key strokes.
For example, here are some alias that I set in `.bashrc` to call the news and wiki command.
```
alias news='command news'
alias wiki='command utility wiki'
```
