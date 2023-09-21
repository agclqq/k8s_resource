# k8s_resource
闲来无事，把k8s资源文件捞一下，看看完整的资源文件是怎么样的

还是顺手推荐一下好用的库，https://github.com/kubernetes/apimachinery ，对应的声明 https://github.com/kubernetes/apimachinery/blob/master/pkg/apis/meta/v1/types.go
## 1. 前提
- 本程序依赖kubectl，所以运行环境中要有kubectl

## 2. 使用方法
```
    go get github.com/agclqq/k8s_resource
```
```
    Explain("job").ToYaml() // 生成job资源文件
    Explain("cronjob").ToYaml() // 生成cronjob资源文件
    Explain("deployment").ToYaml() // 生成deployment资源文件
    Explain("statefulset").ToYaml() // 生成statefulset资源文件
    Explain("service").ToYaml() // 生成service资源文件
    Explain("ingress").ToYaml() // 生成ingress资源文件
    Explain("configmap").ToYaml() // 生成configmap资源文件
    Explain("secret").ToYaml() // 生成secret资源文件
    Explain("serviceaccount").ToYaml() // 生成serviceaccount资源文件
    Explain("role").ToYaml() // 生成role资源文件
    Explain("rolebinding").ToYaml() // 生成rolebinding资源文件
    Explain("clusterrole").ToYaml() // 生成clusterrole资源文件
    Explain("clusterrolebinding").ToYaml() // 生成clusterrolebinding资源文件
    Explain("persistentvolume").ToYaml() // 生成persistentvolume资源文件
    Explain("persistentvolumeclaim").ToYaml() // 生成persistentvolumeclaim资源文件
    Explain("storageclass").ToYaml() // 生成storageclass资源文件
    Explain("namespace").ToYaml() // 生成namespace资源文件
    Explain("limitrange").ToYaml() // 生成limitrange资源文件
    Explain("resourcequota").ToYaml() // 生成resourcequota资源文件
    Explain("horizontalpodautoscaler").ToYaml() // 生成horizontalpodautoscaler资源文件
    Explain("poddisruptionbudget").ToYaml() // 生成poddisruptionbudget资源文件
    Explain("podsecuritypolicy").ToYaml() // 生成podsecuritypolicy资源文件
    Explain("networkpolicy").ToYaml() // 生成networkpolicy资源文件
    Explain("priorityclass").ToYaml() // 生成priorityclass资源文件
    Explain("pod").ToYaml() // 生成pod资源文件
    Explain("daemonset").ToYaml() // 生成daemonset资源文件
    Explain("replicaset").ToYaml() // 生成replicaset资源文件
    Explain("replicationcontroller").ToYaml() // 生成replicationcontroller资源文件
    Explain("job").ToYaml() // 生成job资源文件
    Explain("cronjob").ToYaml() // 生成cronjob资源文件
    Explain("deployment").ToYaml() // 生成deployment资源文件
    Explain("statefulset").ToYaml() // 生成statefulset资源文件
    Explain("service").ToYaml() // 生成service资源文件
    Explain("ingress").ToYaml() // 生成ingress资源文件
    Explain("configmap").ToYaml() // 生成configmap资源文件
    Explain("secret").ToYaml() // 生成secret资源文件
    Explain("serviceaccount").ToYaml() // 生成serviceaccount资源文件
    Explain("role").ToYaml() // 生成role资源文件
    Explain("rolebinding").ToYaml() // 生成rolebinding资源文件
    Explain("clusterrole").ToYaml() // 生成clusterrole资源文件
    Explain("clusterrolebinding").ToYaml() // 生成clusterrolebinding资源文件
    Explain("persistentvolume").ToYaml() // 生成persistentvolume资源文件
    Explain("persistentvolumeclaim").ToYaml() // 生成persistentvolumeclaim资源文件
    Explain("storageclass").ToYaml() // 生成storageclass资源文件
    Explain("namespace").ToYaml() // 生成namespace资源文件
    Explain("limitrange").ToYaml() // 生成limitrange资源文件
    Explain("resourcequota").ToYaml() // 生成resourcequota资源文件
    Explain("horizontalpodautoscaler").ToYaml() // 生成horizontalpodautoscaler资源文件
    Explain("poddisruptionbudget").ToYaml() // 生成poddisruptionbudget资源文件
    Explain("podsecuritypolicy").ToYaml() // 生成podsecuritypolicy资源文件
    Explain("networkpolicy").ToYaml() // 生成networkpolicy资源文件
    Explain("priorityclass").ToYaml() // 生成priorityclass资源文件
    Explain("pod").ToYaml() // 生成pod资源文件
    Explain("daemonset").ToYaml() // 生成daemonset资源文件
    Explain("replicaset").ToYaml() // 生成replicaset资源文件
    Explain("replicationcontroller").ToYaml() // 生成replicationcontroller资源文件
```
