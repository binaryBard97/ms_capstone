this debug folder is analyzing why helm upgrade didnt work.

helm upgrade cmd: `helm upgrade redis bitnami/redis --set image.repository=spdocker81/my-redis-with-criu --set image.tag=latest `

evidence of helm upgrade didnt work

1. criu -v output results into CRIU not being present in one of redis replica pod
2. pods keep failing when doing `kubectl get pods -n redis` after the upgrade cmd
    redis %         kubectl get pods -n  redis                                                                                                                   

    NAME               READY   STATUS             RESTARTS         AGE
    redis-master-0     0/1     ImagePullBackOff   0                31m
    redis-replicas-0   0/1     Running            13 (5m12s ago)   10h
    redis-replicas-1   0/1     CrashLoopBackOff   12 (5m2s ago)    10h
    redis-replicas-2   0/1     ImagePullBackOff   0   
3. and lastly, CRIU checkpoint cmd didnt work.
    `kubectl exec -n redis -it redis-replicas-0 -- criu dump --shell-job --tcp-established -o checkpoint.tar      `