# eadesignca1
Repo to house code for EA Design CA
Combination of code from Labs 1,2 and 4 from the EA Design module labs. The 2 main folders for code are:
- 'Sync' folder contains the source from all the news lab, which is the sync version of system
- 'async' folder contains the source from lab3, which has been modified to provide async version of 'all the news' system.

- 'tests' folder houses the postman test collection json files that were used to test for response times.

- 'GoSample' was just for test GO code. 
- 'curl' folder was just to preserve curl output formatting options


Part 1:
-------
1. change UI so that headers change to channel names i.e. news, weather, sports
2. change sec_con.go to use msg content instead of counting.
3. to do this added another arg to door yaml manifests, which contained the message from each news source i.e. each 'door' yaml manifest was specific to news, weather or sports
4. updated publishing of each door to use this content
5. updated seccon.go to add logging of message content so that this can be viewed for debugging in logs  
6. Used same image for all 'doors' or news sources.
7. yaml manifests need to be updated to change 'wait' and 'interval' times for seccon/newsreader. 
8. remove random number generator to make 'doors' more predictable.
9. Add timestamp to message content to show when message was published

sync urls: http://34.66.193.129:31916/allthenews?style=plain
         : http://34.66.193.129:31916/allthenews?style=colourful
         : http://34.66.193.129:31916/allthenews?style=blackandwhite
async url: http://104.197.52.181:31080/

Part 2:
-------
1. Looked at pyhton script provided
2. Looked at curl options. got curl working locally from windows. Save file to determine format of metrics. See appendix Curl format. Then run curl with write out option.
3. Looked in market place for Prometheus and grafana offerings. Found an option that installs both into cluster. Tried, but it failed twice with 'Does not have minimum availability' on the kube-state-metrics pod. In details, it says innsufficient CPU available. Tips in GKE ui suggest auto-scaling of node pool. could have also gone with cluster, but this seemed a more controlled option. cluster restarted, deleted broken pod, k8s restarted it and looks better.
4. grafana admin pwd didn't work. tried for a long time to reset pwd. gave up. 
5. from marketplace: trying telegraf-influxdb-grafana - failed to deploy from marketplace. think it looks permission related. No time to workout. error: backofflimitexceeded
6. going to deploy prom and graf using Helm chart from local windows machine.
`helm install mike-grafana grafana`
got secret password from gcp as it provided a linux command
7. grabbed cmd to port forward from Services and Ingress UI in GCP
8. ran port forward in powershell locally. logged in.
`kubectl --namespace default port-forward $POD_NAME 3000`
url http://127.0.0.1:8080/

9. Created cloud function
 - used python code py and requirements file provided in lecture
 - function to execure needed to match the function in the py file.
 - tested and saved file here: https://console.cloud.google.com/storage/browser/_details/plotstorage/aname.png?project=eades-273019
---
10. plan is to collect the data using lab4 (calling CF example, can set it up like the sportsfetcher and tear it down again also. or like a door.)
---
 
11. Looking at postman again. Can send results to influxdb using newman
12. Used Helm to install influxdb to k8s, plan is to use Grafana to plot this.
- to connect using port forwarding: kubectl port-forward --namespace default $(kubectl get pods --namespace default -l app=mike-influxdb -o jsonpath='{ .items[0].metadata.name }') 8086:8086
- port forwarding failed. trying to expose from k8s ui using 'expose' option in pod details. TCP:80 and load balancer
13. installed newman locally  `npm install -g newman`
14. installed newman-reporter-influxdb `npm install -g newman-reporter-influxdb`
15. this will run the collection of tests I saved from postman, but I need to authenticate with influxdb. lots of issues trying to get this going.
16. checked influxdb, authentication set by default. updated Helm chart values yaml and delete/reinstall using Helm
- this gave errors during install, but eventually returned ok. could get user and pwd from GCP cloud shell using commands shown in appendix 2. 
- Exposed ip again using k8s UI in GCP. copied port forwarding address. setup locally.
- Downloaded influxdb cli, tried to connect. issues. 
- tried connecting to influx cli inside the pod, using ssh command below in appendix 2
- the service ingress was targeting the wrong port. working after that!
- once connected with the influxdb cli, created the database. 
Some influx cli commands:
`create datanase mydb`
`show databases`
`SELECT * from atnasync`

- Ran newman tests, and was able to run a select from the CLI and see the data in the db
- Setup dashboard in grafana to plot the data from the db
18. added a delay to newman run, so that graph might look better
----
19. async version stopped working. made updates and redeployed (pushed new images and replaced, back up!)
----
20. duplicated postman tests.
- didn't work, postmand has a guid for every collection. so need to create another collection manually as opposed to working with the json.
cmd to run sync tests:
- `newman run ATN_test_scripts.postman_collection.json -r influxdb --reporter-influxdb-server 34.66.228.1 --reporter-influxdb-port 80 --reporter-influxdb-name mydb --reporter-influxdb-measurement atn --reporter-influxdb-username admin --reporter-influxdb-password 1RJCMjQl5N --delay-request 500`

cmd to run async tests:
`newman run ATN_test_scripts.postman_collection_async.json -r influxdb --reporter-influxdb-server 34.66.228.1 --reporter-influxdb-port 80 --reporter-influxdb-name mydb --reporter-influxdb-measurement atn --reporter-influxdb-username admin --reporter-influxdb-password 1RJCMjQl5N --delay-request 500`

---
21. got sick of port forwarding to grafana, created an external IP using ingress in GKE UI.
---
22. looked at lots of options to measure the pod recovery, strange that this is not provided.
- tried kube-state-metrics install as per datadog blog.
23. tried teh k8s dashboard, which looked promising, but ran out of time.
24. in the end, ran out of time. created a crude bash script
pd_deleter, which had timestamps and used
kubectl delete pod
kubectl get pods
Then looked at timings.
Used graphing function to plot the bar chart graph by using the test function. dirty. graph stored in storage bucket and saved to repo under tests. 


NOTE: for kill/restart could try updating the surge option in cluster details to see if it makes this better.

Dev tools
---------
GCP console shell  and cloud editor in browser, used to make code changes
gcp console used
- to run bash scripts
- run docker commands (build, tag, push)
- pull from github repo
- run kubectl commands (get pods, apply, replace)
gcp editor was quiet useful for making quick edits to the code. but kept having to provide creds to github, which was annoying. So push code from local machine as much as possible and did a git pull from gcp.
There was an option to boost the perf of gcp shell and editor. So you get 24hrs of higher performance VM behind the scenes. TBH not that noticeable a difference in perf.

Locally, didn't have docker installed so used VSCode to make edits and push to github. But this was a better editing experience, so the preferred dev environment. Setting up CI/CD pipelines would have been good, but time consuming so chose not to do this.

For checking GO syntax: https://play.golang.org/

- bash
Created bash script to combine commands and help.
'chmod 700 buildpushdeploy' was used to make is executable 

- gcp ui tools
Used GCP k8s logs viewer to check what was being logged.
Log viewer Query (which params were specific to my GCP project):
resource.type="k8s_container"
resource.labels.project_id="eades-273019"
resource.labels.location="us-central1-c"
resource.labels.cluster_name="cluster-1"
resource.labels.namespace_name="default"
resource.labels.pod_name="seccon-deployment-8694f47bf9-mr2wv"

Used GCP container registry to check what images were published
Used GCP kubernetes engine ui useful for looking at details of deployment

References:
-----------
1. https://cloud.google.com/kubernetes-engine/docs
2. https://cloud.google.com/kubernetes-engine/docs/how-to/cluster-autoscaler
3. https://github.com/GoogleCloudPlatform/click-to-deploy/blob/master/k8s/prometheus/README.md
4. https://docs.influxdata.com/influxdb/v1.2/introduction/installation/ - https://docs.influxdata.com/influxdb/v1.2/guides/querying_data/
5. https://www.npmjs.com/package/newman + https://www.npmjs.com/package/newman#external-reporters
6. https://github.com/vs4vijay/newman-reporter-influxdb
7. https://github.com/GoogleCloudPlatform/click-to-deploy/blob/master/k8s/prometheus/README.md
8. https://grafana.com/
9. https://www.datadoghq.com/blog/monitoring-kubernetes-performance-metrics/
10. https://github.com/johanhaleby/kubetail


APPENDIX:
---------
1. Curl format
    time_namelookup:  %{time_namelookup}\n
    time_connect:  %{time_connect}\n
    time_appconnect:  %{time_appconnect}\n
    time_pretransfer:  %{time_pretransfer}\n
    time_redirect:  %{time_redirect}\n
    time_starttransfer:  %{time_starttransfer}\n
    ----------\n
    time_total:  %{time_total}\n

2. Helm influxdb notes:
NOTES from influxdb Helm install:

InfluxDB can be accessed via port 8086 on the following DNS name from within your cluster:

- http://mike3-influxdb.default:8086

You can easily connect to the remote instance with your local influx cli. To forward the API port to localhost:8086 run the following:

- kubectl port-forward --namespace default $(kubectl get pods --namespace default -l app=mike3-influxdb -o jsonpath='{ .items[0].metadata.name }') 8086:8086

You can also connect to the influx cli from inside the container. To open a shell session in the InfluxDB pod run the following:

- kubectl exec -i -t --namespace default $(kubectl get pods --namespace default -l app=mike3-influxdb -o jsonpath='{.items[0].metadata.name}') /bin/sh

To tail the logs for the InfluxDB pod run the following:

- kubectl logs -f --namespace default $(kubectl get pods --namespace default -l app=mike3-influxdb -o jsonpath='{ .items[0].metadata.name }')

To retrieve the default user name:

- echo $(kubectl get secret mike3-influxdb-auth -o "jsonpath={.data['influxdb-user']}" --namespace default | base64 --decode)

To retrieve the default user password:

- echo $(kubectl get secret mike3-influxdb-auth -o "jsonpath={.data['influxdb-password']}" --namespace default | base64 --decode)

