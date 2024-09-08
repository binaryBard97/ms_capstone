# enhance-schedule-k8s
**MS CS Capstone at RIT - Understanding and Optimizing Scheduling in Kubernetes**

Singularity, a distributed scheduling service developed by Microsoft, is designed to efficiently manage deep learning workloads across a global fleet of AI accelerators such as GPUs and FPGAs. It achieves this by making AI workloads **preemptible, migratable,** and **dynamically resizable**. Singularity’s ability to checkpoint, preempt, and migrate workloads transparently allows jobs to be moved across nodes, clusters, or even regions without losing progress.

Kubernetes, while using **preemption** and **migration** to improve resource utilization and efficiency, does not do so to the same extent as Singularity. Specifically, Kubernetes lacks native support for **checkpointing** and **resuming** jobs, which means that when a pod is preempted, it often has to restart from scratch unless the application itself is designed to handle restarts. Kubernetes can preempt jobs, but it does not "migrate" them in the sense of seamlessly moving a running job from one node to another with its state intact.

Although Kubernetes provides tools for scaling and scheduling to enhance efficiency, it lacks Singularity’s advanced features like **checkpointing, seamless migration,** and **real-time dynamic scaling** without restarting jobs. Singularity’s more fine-grained approach to resource management allows for better use of global resources through its stateful migration and elasticity.

The goal of this capstone project is to bring key features from Singularity, such as **checkpointing** and **dynamic workload migration**, to Kubernetes in order to improve scheduling efficiency.
