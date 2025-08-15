#Install az-cli

```bash
az login
```

```bash
# store the full group name path in a variable.
groupId=$(az group show \
  --name <resource-group-name> \
  --query id --output tsv)
```

```bash
# create a service principal for the group (outputs client id and secret).
az ad sp create-for-rbac \
  --scope $groupId \
  --role Contributor \
  --sdk-auth
```

## Give access for the service principal to the container registry

```bash
registryId=$(az acr show \
  --name nokkatest \
  --resource-group nokka-test \
  --query id --output tsv)
```

```bash
# use az role assignment create to assign the AcrPush role, which gives push and pull access to the registry. Substitute the client ID of your service principal:
az role assignment create \
  --assignee <client-id> \
  --scope $registryId \
  --role AcrPush
```
