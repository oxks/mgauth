# Buffalo go framework startup template. It has Goth and Auth modules integrated. 

## Whom it can help

It's intended to be a startup application for projects based on Buffalo framework that need to have Goth (cloud login) and Auth (local login) to be present from the very beginning. 

## What it does not do 

Goth authentification does not store any data in the database. You can store it and operate the information as per your needs. 

## Setup and starting the app

Assuming that you have a working Buffalo installation in your system. 

Clone the project. 

Set global variables for private data (GOOGLE_KEY, GITHUB_KEY etc), refer to file actions/auth.go

run: 
	$ yarn add webpack --dev
	$ buffalo dev

and open in the browser: 
http://127.0.0.1:3000


Good luck!

[Powered by Buffalo](http://gobuffalo.io)
