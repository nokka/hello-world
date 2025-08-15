AZURE_REPO=nokkatest-b5exb6f6gge3d7bk.azurecr.io/hello-world
AZURE_COMMIT_TAG=$(AZURE_REPO):$(GITHUB_SHA)
VALID_TAG=$(shell echo $(TAG_NAME) | sed 's/[^a-z0-9_\.-]/-/g')

docker/local:
	docker build -f Dockerfile -t hello-world:local . 

# builds docker image with with the registry commit tag.
docker/build:
	docker build -f Dockerfile -t $(AZURE_COMMIT_TAG) .

# tags the image with a valid tag.
docker/tag:
	docker tag $(AZURE_COMMIT_TAG) $(AZURE_REPO):$(VALID_TAG)

# push the image to Azure container registry.
docker/push:
	docker push -a $(AZURE_REPO)
