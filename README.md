# Publish Newsletter Curated by a Group of People

Even though the name says **Group of People**, it can be just you. The aim of this project is to publish and archive newsletters to a target email address. 

![](https://github.com/codingpot/newsletter_awesome_articles/blob/main/assets/overview.png)

## How to publish?

It basically collects the number of markdowns under `/current` directory. The name of markdowns should be formatted as `YYYY-MM-DD hh:mm` because the newsletter system could pick up the most latest top-k ones.

```shell
# pseudo CLI

$ newsletter publish current --top K
```

## How to archive?

When the newsletter system publishes a newsletter, it will collect up to top-k markdowns and move them to `/archive/issue{number}` directory. Also possibly, it will give you a way to send archieved newsletter afterwards with a CLI.

```shell
# pseudo CLI

$ newsletter publish archive --issue NUMBER
```

## Markdown format 

```md
date: YYYY-MM-DD hh:mm
author: author name
title: title
thumbnail: image URL 
link: link URL to the original 
summary: preferably up to 2-3 sentences
opinion: preferably up to 5-8 sentences
```

## Newsletter layout

The referenced newsletter layout is below that I have used for other purposes.

![](https://raw.githubusercontent.com/deep-diver/fb-group-post-fetcher/master/static/images/preview.png)
