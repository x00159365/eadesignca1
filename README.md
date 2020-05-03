# eadesignca1
Repo to house code for EA Design CA
Combination of code from Labs 1,2 and 4 from the EA Design module labs.

- Sync folder contains the source from all the news lab, which is the sync version of system

- async folder contains the source from lab3, which has been modified to provide async version of 'all the news' system.

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

9. Created cloud function
 - used python code py and requirements file provided in lecture
 - function to execure needed to match the function in the py file.

10. plan is to collect the data using lab4 (calling CF example, can set it up like the sportsfetcher and tear it down again also. or like a door.)

 


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
4. 

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