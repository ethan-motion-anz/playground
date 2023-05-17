
## Input files

### goals.json

This file is fetched manually: 
- Open dev tools
- Network tab
- Toggle Fetch/XHR
- Right-click the `https://api.xplore.services.x.gcp.anz/graphql?operationName=GetHierarchy&variables...` request, and `Copy response`
- Save the response as `goals.json` under the `/data` folder in this project.

There is an Xplore API for fetching all Goals.
https://backstage.service.anz/catalog/default/api/xplore-api/definition#/goals/getGoals

We could get access to this API in future to fetch this file programmatically. 


### slos.yaml

This file is fetched manually:

Run `sloctl get slos -l anz-x-tech-asset=fabric -o yaml | >> slos.yaml`.

The `gopkg.in/yaml.v3` packag seems to require the YAML file have a top-level(?), so add `slos:` to the top of the yaml file and indent everything else.
