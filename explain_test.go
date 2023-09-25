package main

import (
	"testing"
)

func TestExplain(t *testing.T) {
	//Explain("job").ToYamlFile()                     // 生成job资源文件
	//Explain("cronjob").ToYamlFile()                 // 生成cronjob资源文件
	//Explain("deployment").ToYamlFile()              // 生成deployment资源文件
	Explain("statefulset").ToYamlFile()             // 生成statefulset资源文件
	Explain("service").ToYamlFile()                 // 生成service资源文件
	Explain("ingress").ToYamlFile()                 // 生成ingress资源文件
	Explain("configmap").ToYamlFile()               // 生成configmap资源文件
	Explain("secret").ToYamlFile()                  // 生成secret资源文件
	Explain("serviceaccount").ToYamlFile()          // 生成serviceaccount资源文件
	Explain("role").ToYamlFile()                    // 生成role资源文件
	Explain("rolebinding").ToYamlFile()             // 生成rolebinding资源文件
	Explain("clusterrole").ToYamlFile()             // 生成clusterrole资源文件
	Explain("clusterrolebinding").ToYamlFile()      // 生成clusterrolebinding资源文件
	Explain("persistentvolume").ToYamlFile()        // 生成persistentvolume资源文件
	Explain("persistentvolumeclaim").ToYamlFile()   // 生成persistentvolumeclaim资源文件
	Explain("storageclass").ToYamlFile()            // 生成storageclass资源文件
	Explain("namespace").ToYamlFile()               // 生成namespace资源文件
	Explain("limitrange").ToYamlFile()              // 生成limitrange资源文件
	Explain("resourcequota").ToYamlFile()           // 生成resourcequota资源文件
	Explain("horizontalpodautoscaler").ToYamlFile() // 生成horizontalpodautoscaler资源文件
	Explain("poddisruptionbudget").ToYamlFile()     // 生成poddisruptionbudget资源文件
	Explain("podsecuritypolicy").ToYamlFile()       // 生成podsecuritypolicy资源文件
	Explain("networkpolicy").ToYamlFile()           // 生成networkpolicy资源文件
	Explain("priorityclass").ToYamlFile()           // 生成priorityclass资源文件
	Explain("pod").ToYamlFile()                     // 生成pod资源文件
	Explain("daemonset").ToYamlFile()               // 生成daemonset资源文件
	Explain("replicaset").ToYamlFile()              // 生成replicaset资源文件
	Explain("replicationcontroller").ToYamlFile()   // 生成replicationcontroller资源文件
	Explain("job").ToYamlFile()                     // 生成job资源文件
	Explain("cronjob").ToYamlFile()                 // 生成cronjob资源文件
	Explain("deployment").ToYamlFile()              // 生成deployment资源文件
	Explain("statefulset").ToYamlFile()             // 生成statefulset资源文件
	Explain("service").ToYamlFile()                 // 生成service资源文件
	Explain("ingress").ToYamlFile()                 // 生成ingress资源文件
	Explain("configmap").ToYamlFile()               // 生成configmap资源文件
	Explain("secret").ToYamlFile()                  // 生成secret资源文件
	Explain("serviceaccount").ToYamlFile()          // 生成serviceaccount资源文件
	Explain("role").ToYamlFile()                    // 生成role资源文件
	Explain("rolebinding").ToYamlFile()             // 生成rolebinding资源文件
	Explain("clusterrole").ToYamlFile()             // 生成clusterrole资源文件
	Explain("clusterrolebinding").ToYamlFile()      // 生成clusterrolebinding资源文件
	Explain("persistentvolume").ToYamlFile()        // 生成persistentvolume资源文件
	Explain("persistentvolumeclaim").ToYamlFile()   // 生成persistentvolumeclaim资源文件
	Explain("storageclass").ToYamlFile()            // 生成storageclass资源文件
	Explain("namespace").ToYamlFile()               // 生成namespace资源文件
	Explain("limitrange").ToYamlFile()              // 生成limitrange资源文件
	Explain("resourcequota").ToYamlFile()           // 生成resourcequota资源文件
	Explain("horizontalpodautoscaler").ToYamlFile() // 生成horizontalpodautoscaler资源文件
	Explain("poddisruptionbudget").ToYamlFile()     // 生成poddisruptionbudget资源文件
	Explain("podsecuritypolicy").ToYamlFile()       // 生成podsecuritypolicy资源文件
	Explain("networkpolicy").ToYamlFile()           // 生成networkpolicy资源文件
	Explain("priorityclass").ToYamlFile()           // 生成priorityclass资源文件
	Explain("pod").ToYamlFile()                     // 生成pod资源文件
	Explain("daemonset").ToYamlFile()               // 生成daemonset资源文件
	Explain("replicaset").ToYamlFile()              // 生成replicaset资源文件
	Explain("replicationcontroller").ToYamlFile()   // 生成replicationcontroller资源文件
}
