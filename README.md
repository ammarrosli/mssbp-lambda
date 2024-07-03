# MSSBP Microsites Serverless API

## Frontend



## Serverless API

Use *POST* method to the given URL to create a lead

### Base URL

- **Staging**: https://l4x36a6kz1.execute-api.ap-southeast-1.amazonaws.com/dev/mssbp/contacts
- **Live**: 

### Endpoint

1. Mssbp: `mssbp/contacts`

## Maintenance

1. To update staging function, change branch to `staging`.
2. After commit in `staging` branch, merge with `main` branch

## Standard Deployment Steps

1. Run `rm -rf build && make build && make zip`
2. Run `sls deploy`