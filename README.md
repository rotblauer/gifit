

# `$ gifit`: commit messages worth a ~~thousand pictures~~ gif

gets used by doing:
- `$ gifit cats sorting things`
- parses $@ to query encoded string per [giphy api specs](https://github.com/Giphy/GiphyAPI)
- queries giphy public api and parses json response to find suitable gif id
- formats gif id into a raw gif url 
    + final .gif downloadable/markdown-insertable url looks like `http://i.giphy.com/jFGozCmdXxm0w.gif`, where jFGozCmdXxm0w is the gif's ID.
- inserts that url into md image syntax per `![cats sorting things](http://i.giphy.com/jFGozCmdXxm0w.gif)` ie ![cats sorting things](http://i.giphy.com/iuHaJ0D7macZq.gif)
- does `$ git add -A && git commit -m "$that_md_img_string"`
