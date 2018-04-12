# Introduction

Create golang scrath docker image.

# Commands

1. build go binary: `GOOS=linux go build  -o hello .`

2. build docker image: `docker build -t hello .`

3. run container: `docker run hello`