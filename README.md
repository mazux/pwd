# Password manager

A simple passaword manager in golang.

## Domain model

Each profile has:
- Username
- Secret (hashed value)
- list of creentials (i.e logins):
  - Domain
  - Username
  - Password (encrypted password with a salt generated from both domain and username values)

## Application (use cases)

- Signup: to create a new profile with empty list of logins
- Adding new credentials
- Removing exisiting credentials
- Fetch a specific credentials by domain name and username
- Search for multiple credentials by domain name

## Presentation

For simplicity, I'm only considering CLI for now :')
But implementing a UI would be fairly simple because of the isolation of all layers and responsibilities.

## Infrustructure

For the data storage implementation, I'm considering file storage (as json) and mysql one as well.
Via environment variables you can choose either or both implementations ;)
