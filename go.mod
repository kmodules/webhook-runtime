module kmodules.xyz/webhook-runtime

go 1.12

require (
	github.com/evanphx/json-patch v4.5.0+incompatible
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/json-iterator/go v1.1.6
	gomodules.xyz/jsonpatch/v2 v2.0.0
	k8s.io/api v0.0.0-20190503110853-61630f889b3c
	k8s.io/apimachinery v0.0.0-20190508063446-a3da69d3723c
	k8s.io/apiserver v0.0.0-20190516230822-f89599b3f645
	k8s.io/client-go v11.0.0+incompatible
	k8s.io/kubernetes v1.14.0
	k8s.io/utils v0.0.0-20190506122338-8fab8cb257d5 // indirect
	kmodules.xyz/client-go v0.0.0-20190808141354-bbb9e14f60ab
	kmodules.xyz/openshift v0.0.0-20190808144841-c8f9a927f1d1
)

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.2+incompatible
	k8s.io/api => k8s.io/api v0.0.0-20190313235455-40a48860b5ab
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.0.0-20190315093550-53c4693659ed
	k8s.io/apimachinery => github.com/kmodules/apimachinery v0.0.0-20190508045248-a52a97a7a2bf
	k8s.io/apiserver => github.com/kmodules/apiserver v0.0.0-20190508082252-8397d761d4b5
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.0.0-20190314001948-2899ed30580f
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.0.0-20190314002645-c892ea32361a
	k8s.io/klog => k8s.io/klog v0.3.0
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.0.0-20190314000639-da8327669ac5
	k8s.io/kube-openapi => k8s.io/kube-openapi v0.0.0-20190228160746-b3a7cee44a30
	k8s.io/metrics => k8s.io/metrics v0.0.0-20190314001731-1bd6a4002213
	k8s.io/utils => k8s.io/utils v0.0.0-20190221042446-c2654d5206da
	sigs.k8s.io/structured-merge-diff => sigs.k8s.io/structured-merge-diff v0.0.0-20190302045857-e85c7b244fd2
)
