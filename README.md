# eadesignca1
Repo to house code for EA Design CA
Combination of code from Labs 1,2 and 4 from the EA Design module labs.

- Sync folder contains the source from all the news lab, which is the sync version of system

- async folder contains the source from lab3, which has been modified to provide async version of 'all the news' system.

1. change UI so that headers change to channel names i.e. news, weather, sports
2. change sec_con.go to use msg content instead of counting.
3. to do this added another arg to door yaml manifests, which contained the message from each news source i.e. each 'door' yaml manifest was specific to news, weather or sports

Used GCP k8s logs viewer to check what was being logged.
Log viewer Query (which params were specific to my GCP project):
resource.type="k8s_container"
resource.labels.project_id="eades-273019"
resource.labels.location="us-central1-c"
resource.labels.cluster_name="cluster-1"
resource.labels.namespace_name="default"
resource.labels.pod_name="seccon-deployment-8694f47bf9-mr2wv"

For checking go syntax: https://play.golang.org/

Add timestamp to message content to show when message was written to redis

