# Copyright 2017 The Kubernetes Authors. All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

TAG = v6
REGISTRY = docker.io/lcwyo
IMAGE = $(REGISTRY)/gb-frontend-nginx

all: clean build push

push: 
	docker push $(IMAGE):$(TAG)
build:
	docker build -t $(IMAGE):$(TAG) .

clean: 
	docker rmi -f $(IMAGE):$(TAG) || true
sub-clean-%:
	docker rmi -f $(IMAGE)-$*:$(TAG) || true
