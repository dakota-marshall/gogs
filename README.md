# gogs
![example workflow](https://github.com/dakota-marshall/gogs/actions/workflows/go.yml/badge.svg)
[![codecov](https://codecov.io/gh/dakota-marshall/gogs/graph/badge.svg?token=2FDL5F3OES)](https://codecov.io/gh/dakota-marshall/gogs)
[![Go Reference](https://pkg.go.dev/badge/github.com/dakota-marshall/gogs.svg)](https://pkg.go.dev/github.com/dakota-marshall/gogs)
[![Go Report Card](https://goreportcard.com/badge/github.com/dakota-marshall/gogs)](https://goreportcard.com/report/github.com/dakota-marshall/gogs)

An API wrapper library for the Go server [https://online-go.com/](https://online-go.com/) written in... Go.

This is a WIP reimplementation of my [Python-OGS](https://gitlab.com/dakota.marshall/ogs-python/-/blob/main/src/ogsapi/ogsrestapi.py?ref_type=heads) library.

I am using Bruno to document the OGS API, as its official documentation is lacking.


## Current To-Do List:

- [x] Implement the core server class to handle the actual API calls
- [/] Implement the basic REST API calls 
- [ ] Figure out how to re-implement the Socket.io API in Go

## REST API Endpoints Checklist

- [ ] `/announcements`
	- [ ] `/history`
- [ ] `/demos`
- [ ] `/reviews`
	- [ ] `/{id}`
	- [ ] `/{id}/png`
	- [ ] `/{id}/sgf`
- [ ] `/games`
	- [x] `/{id}` ✅ 2024-10-15
	- [x] `/{id}/png` ✅ 2024-10-15
	- [ ] `/{id}/sgf`
	- [ ] `/{id}/reviews`
- [ ] `/groups`
	- [ ] `/{id}`
	- [ ] `/{id}/ladders`
	- [ ] `/{id}/members`
	- [ ] `/{id}/news`
- [ ] `/ladders`
	- [ ] `/{id}`
- [ ] `/leaderboards`
- [ ] `/library/{id}`
- [ ] `/me`
	- [ ] `/account_settings`
	- [ ] `/blocks`
	- [ ] `/challenges`
	- [ ] `/friends`
	- [ ] `/games`
	- [ ] `/groups`
	- [ ] `/ladders`
	- [ ] `/settings`
- [ ] `/ui/overview`
- [ ] `/players`
	- [ ] `/{id}`
	- [ ] `/{id}/full`
	- [ ] `/{id}/games`
	- [ ] `/{id}/groups`
	- [ ] `/{id}/ladders`
	- [ ] `/{id}/tournaments`
- [ ] `/puzzles`
	- [ ] `/full`
	- [ ] `/{id}`
	- [ ] `/collections`
- [ ] `/tournaments`
	- [ ] `/{id}`
	- [ ] `/{id}/players`
	- [ ] `/{id}/rounds`
- [ ] `/tournament_records`
	- [ ] `/{id}`
	- [ ] `/{id}/players`
	- [ ] `/{id}/rounds`
