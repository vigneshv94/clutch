module.exports = {
  "@clutch-sh/ec2": {
    terminateInstance: {
      trending: true,
      componentProps: {
        resolverType: "clutch.aws.ec2.v1.Instance",
      },
    },
    resizeAutoscalingGroup: {
      componentProps: {
        resolverType: "clutch.aws.ec2.v1.AutoscalingGroup",
        notes: [
          {
            severity: "info",
            text:
              "The autoscaling group may take several minutes to bring additional instances online.",
          },
        ],
      },
    },
  },
  "@clutch-sh/envoy": {
    remoteTriage: {
      trending: true,
      componentProps: {
        options: {
          Clusters: "clusters",
          Listeners: "listeners",
          Runtime: "runtime",
          Stats: "stats",
          "Server Info": "serverInfo",
        },
      },
    },
  },
  "@clutch-sh/k8s": {
    deletePod: {
      componentProps: {
        resolverType: "clutch.k8s.v1.Pod",
      },
    },
    resizeHPA: {
      trending: true,
      componentProps: {
        resolverType: "clutch.k8s.v1.HPA",
      },
    },
  },
  "@clutch-sh/experimentation": {
    listExperiments: {
      description: "Manage server and mobile fault injection experiments.",
      trending: true,
      componentProps: {
        columns: [
          { id: "target", header: "Target" },
          { id: "start_time", header: "Start Time", sortable: true },
          { id: "end_time", header: "End Time" },
          { id: "status", header: "Status" },
        ],
        links: [
          {
            displayName: "Start Server Experiment",
            path: "/server-experimentation/start",
          },
          {
            displayName: "Start Mobile Experiment",
            path: "/mobile-experimentation/start-experiment",
          },
        ],
      },
    },
    viewExperimentRun: {},
  },
  "@clutch-sh/server-experimentation": {
    startExperiment: {},
  },
};
