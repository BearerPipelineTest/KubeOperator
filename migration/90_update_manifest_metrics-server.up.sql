update ko_cluster_manifest set other_vars = '[{\"name\":\"coredns\",\"version\":\"1.6.7\"},{\"name\":\"dns-cache\",\"version\":\"1.17.0\"},{\"name\":\"traefik\",\"version\":\"v2.2.1\"},{\"name\":\"ingress-nginx\",\"version\":\"0.33.0\"},{\"name\":\"metrics-server\",\"version\":\"v0.5.0\"},{\"name\":\"helm-v2\",\"version\":\"v2.16.9\"},{\"name\":\"helm-v3\",\"version\":\"v3.2.4\"}]' where name in ("v1.18.4-ko1", "v1.18.6-ko1", "v1.18.8-ko1", "v1.18.10-ko1");

update ko_cluster_manifest set other_vars = '[{\"name\":\"coredns\",\"version\":\"1.6.7\"},{\"name\":\"dns-cache\",\"version\":\"1.17.0\"},{\"name\":\"traefik\",\"version\":\"v2.2.1\"},{\"name\":\"ingress-nginx\",\"version\":\"0.33.0\"},{\"name\":\"metrics-server\",\"version\":\"v0.5.0\"},{\"name\":\"helm-v2\",\"version\":\"v2.17.0\"},{\"name\":\"helm-v3\",\"version\":\"v3.4.1\"}]' where name in ("v1.18.12-ko1", "v1.18.14-ko1", "v1.18.15-ko1", "v1.18.18-ko1");

update ko_cluster_manifest set other_vars = '[{\"name\":\"coredns\",\"version\":\"1.7.0\"},{\"name\":\"dns-cache\",\"version\":\"1.17.0\"},{\"name\":\"traefik\",\"version\":\"v2.4.8\"},{\"name\":\"ingress-nginx\",\"version\":\"0.33.0\"},{\"name\":\"metrics-server\",\"version\":\"v0.5.0\"},{\"name\":\"helm-v2\",\"version\":\"v2.17.0\"},{\"name\":\"helm-v3\",\"version\":\"v3.6.0\"}]' where name in ("v1.18.20-ko1", "v1.20.6-ko1", "v1.20.8-ko1");

update ko_cluster_manifest set other_vars = '[{\"name\":\"coredns\",\"version\":\"1.7.0\"},{\"name\":\"dns-cache\",\"version\":\"1.17.0\"},{\"name\":\"traefik\",\"version\":\"v2.2.1\"},{\"name\":\"ingress-nginx\",\"version\":\"0.33.0\"},{\"name\":\"metrics-server\",\"version\":\"v0.5.0\"},{\"name\":\"helm-v2\",\"version\":\"v2.17.0\"},{\"name\":\"helm-v3\",\"version\":\"v3.4.1\"}]' where name in ("v1.20.4-ko1");
