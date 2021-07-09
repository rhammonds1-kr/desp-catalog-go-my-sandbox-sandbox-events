This repository contains generated golang code for:
* Event Family: https://desp.kroger.com/event-family/d3471fbb-8d8a-478b-bffa-6a59426ade1f
* Version: v0.1.2

To use this in your go client, add the following requirements to your go.mod file.

```
module github.com/krogertechnology/my-go-client-application

go 1.13

require (
	github.com/actgardner/gogen-avro/v9 v9.0.0 // indirect
	github.com/rhammonds1-kr/desp-catalog-go-my-sandbox-sandbox-events v0.1.2
	...
```