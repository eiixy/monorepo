k8s.config.%:
	$(eval ENV:= $*)
	@monorepo-deploy k8s gen config --env=$(ENV)

k8s.config.qa: k8s.config.qa
k8s.config.live: k8s.config.live
k8s.config.local: k8s.config.local
