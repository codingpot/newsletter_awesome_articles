[![ci](https://github.com/codingpot/newsletter_awesome_articles/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/codingpot/newsletter_awesome_articles/actions/workflows/ci.yml)  [![send_newsletter](https://github.com/codingpot/newsletter_awesome_articles/actions/workflows/newsletter.yml/badge.svg)](https://github.com/codingpot/newsletter_awesome_articles/actions/workflows/newsletter.yml)

# Publish Newsletter Curated by a Group of People

Even though the name says **Group of People**, it can be just you. The aim of this project is to publish and archive newsletters to a target email address. 

![](https://github.com/codingpot/newsletter_awesome_articles/blob/main/assets/overview.png)

## How to publish?

It basically collects every yaml files under `/current` directory. The name of yaml file should be formatted as `YYYY-MM-DD Title` so the files can be ordered correctly by themselves.

```shell
# under src/ directory
$ go run main.go publish current 
```

## YAML format 

```yaml
date: YYYY-MM-DD hh:mm
author: author name
title: title
thumbnail: image URL 
link: link URL to the original 
summary: preferably up to 2-3 sentences
opinion: preferably up to 5-8 sentences
tags: ["tag1", "tag2"]
```

## Newsletter layout

The referenced newsletter layout is below that I have used for other purposes.

![](https://i.ibb.co/NLN2Lhq/Screen-Shot-2022-01-23-at-10-57-12-PM.png)

## Todo

- [X] Parsing Markdown
- [X] Filling Template
- [X] Sending Email
- [X] Move Current Markdowns to Archive
- [X] Write CI/CD script (GitHub Action)
