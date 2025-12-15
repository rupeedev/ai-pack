# Neo4j Deployment on AWS - Quick Reference

## 🎯 TL;DR

| Use Case | Best Option | Cost | Why |
|----------|-------------|------|-----|
| **MVP/POC** | Neo4j AuraDB Free | $0 | Zero setup, managed |
| **Small Prod (< 100GB)** | Neo4j AuraDB | $65+/mo | Managed, optimized |
| **Large Prod (> 100GB)** | EC2 + EBS | $150-500/mo | Performance, control |
| **Cost-Conscious** | EC2 + EBS gp3 | $150/mo | Self-managed |
| **Dev/Testing** | Local Docker | $0 | Simple, fast |

## 📊 Quick Comparison Matrix

| Feature | EC2 | ECS | EKS | AuraDB |
|---------|-----|-----|-----|--------|
| **Setup Complexity** | Medium | Medium | High | Easy |
| **Performance** | Excellent | Good | Good | Excellent |
| **HA/Clustering** | Manual | No | Helm charts | Built-in |
| **Scaling** | Vertical | Limited | Complex | Auto |
| **Ops Overhead** | High | Medium | High | None |
| **Best For** | Production | Dev/staging | K8s shops | All use cases |
| **Production Ready** | ✅ | ⚠️ (ECS on EC2) | ✅ | ✅ |

## 💰 Cost Comparison (us-east-1)

| Option | Monthly Cost | vCPU | RAM | Storage | Use Case |
|--------|--------------|------|-----|---------|----------|
| **AuraDB Free** | **$0** | 0.5 | 1GB | 1GB | POC, learning |
| **AuraDB Professional** | **$65** | 1 | 2GB | 8GB | Small production |
| **AuraDB Business** | **$400+** | 4+ | 16GB+ | 100GB+ | Enterprise |
| **EC2 m5.large** | **$75** | 2 | 8GB | + EBS | Small self-managed |
| **EC2 m5.xlarge + EBS** | **$150** | 4 | 16GB | 100GB gp3 | Medium production |
| **EC2 m5.2xlarge + EBS** | **$300** | 8 | 32GB | 200GB gp3 | Large production |
| **EC2 r5.xlarge + EBS** | **$250** | 4 | 32GB | 100GB gp3 | Memory-intensive |
| **EKS + m5.xlarge nodes** | **$220+** | 4+ | 16GB+ | 100GB+ | K8s production |

## 📏 Sizing Recommendations

### By Data Size

| Graph Size | Nodes | Relationships | RAM | Instance | Storage | Cost/mo |
|------------|-------|---------------|-----|----------|---------|---------|
| **Small** | < 1M | < 5M | 8-16GB | m5.large/xlarge | 50GB gp3 | $75-150 |
| **Medium** | 1M-10M | 5M-50M | 16-32GB | m5.xlarge/2xlarge | 100-200GB gp3 | $150-300 |
| **Large** | 10M-100M | 50M-500M | 32-64GB | m5.2xlarge/4xlarge | 200-500GB gp3 | $300-600 |
| **Enterprise** | > 100M | > 500M | 64GB+ | r5.2xlarge+ or AuraDB | 500GB+ io2 | $500+ |

### By Query Pattern

| Pattern | Instance Type | Why |
|---------|---------------|-----|
| **Read-heavy** | m5/c5 series | Balanced compute |
| **Write-heavy** | m5 series | Good I/O + compute |
| **Memory-intensive traversals** | r5/r6i series | More RAM |
| **Mixed workload** | m5 series | Best all-around |

---

## 🚀 Deployment Options

### 1️⃣ Neo4j AuraDB (Managed) 🏆 RECOMMENDED

#### ✅ Pros
- Fully managed by Neo4j
- Auto backups, HA, monitoring
- **Free tier** (1GB)
- Best performance optimization
- Multi-region support
- Automatic upgrades

#### ❌ Cons
- Higher cost at scale
- Less infrastructure control
- Vendor dependency
- Data egress fees

#### 📋 Best For
- **Startups, MVPs**
- Limited DevOps resources
- Focus on application development
- Need guaranteed performance
- **Any production workload**

#### 🔧 Setup
```bash
# 1. Sign up: console.neo4j.io
# 2. Create database
# 3. Get connection URI and credentials
# 4. Connect via Bolt protocol

NEO4J_URI="neo4j+s://xxxxx.databases.neo4j.io"
NEO4J_USER="neo4j"
NEO4J_PASSWORD="your-password"
```

#### 💡 Pricing Tiers
- **Free**: 1GB RAM, 0.5 vCPU, 1GB storage - **$0**
- **Professional**: 2-8GB RAM - **$65-200/mo**
- **Business**: 16GB+ RAM, dedicated - **$400+/mo**
- **Enterprise**: Custom SLA, support - **Custom pricing**

---

### 2️⃣ EC2 + EBS (Self-Managed) ⭐ For Control

#### ✅ Pros
- **Best performance** (direct I/O)
- Full control over tuning
- Predictable costs
- Simple architecture
- Neo4j's recommended approach

#### ❌ Cons
- You manage OS, backups
- Manual HA setup
- No auto-scaling
- Manual failover

#### 📋 Best For
- **Production > 100GB**
- Performance-critical apps
- Cost-conscious teams
- Existing EC2 infrastructure
- Large datasets (> 1TB)

#### 🔧 Setup

```bash
#!/bin/bash
# EC2 User Data

# Format and mount EBS volume
mkfs -t ext4 /dev/xvdf
mkdir -p /var/lib/neo4j
mount /dev/xvdf /var/lib/neo4j

# Run Neo4j
docker run -d \
  --name neo4j \
  -p 7474:7474 \
  -p 7687:7687 \
  -v /var/lib/neo4j:/data \
  -e NEO4J_AUTH=neo4j/password \
  -e NEO4J_dbms_memory_heap_initial__size=2g \
  -e NEO4J_dbms_memory_heap_max__size=4g \
  -e NEO4J_dbms_memory_pagecache_size=4g \
  neo4j:latest
```

#### 💻 Instance Recommendations

**Compute-Optimized (Read-Heavy)**
| Data Size | Instance | RAM | vCPU | Storage | Cost/mo |
|-----------|----------|-----|------|---------|---------|
| < 50GB | **m5.large** | 8GB | 2 | 50GB gp3 | $75 |
| 50-100GB | **m5.xlarge** | 16GB | 4 | 100GB gp3 | $150 |
| 100-500GB | **m5.2xlarge** | 32GB | 8 | 200GB gp3 | $300 |
| > 500GB | **m5.4xlarge** | 64GB | 16 | 500GB io2 | $600 |

**Memory-Optimized (Complex Traversals)**
| Data Size | Instance | RAM | vCPU | Storage | Cost/mo |
|-----------|----------|-----|------|---------|---------|
| < 100GB | **r5.xlarge** | 32GB | 4 | 100GB gp3 | $250 |
| 100-500GB | **r5.2xlarge** | 64GB | 8 | 200GB gp3 | $500 |
| > 500GB | **r5.4xlarge** | 128GB | 16 | 500GB io2 | $1,000 |

#### 📦 Storage Configuration

| Data Size | EBS Type | Size | IOPS | Throughput | Cost/mo |
|-----------|----------|------|------|------------|---------|
| < 100GB | gp3 | 100GB | 3000 | 125 MB/s | $8 |
| 100-500GB | gp3 | 200GB | 5000 | 250 MB/s | $20 |
| 500GB-1TB | gp3 | 500GB | 10000 | 500 MB/s | $55 |
| > 1TB (production) | io2 | 1TB+ | 20000+ | 1000 MB/s | $125+ |

---

### 3️⃣ EKS (Kubernetes)

#### ✅ Pros
- StatefulSet support
- EBS CSI driver
- Official Helm charts
- HA with causal clustering
- Auto-scaling

#### ❌ Cons
- **High complexity**
- Requires K8s expertise
- EKS control plane cost ($73/mo)
- Overkill for single DB
- More moving parts

#### 📋 Best For
- Already using Kubernetes
- Need clustering/HA
- Multi-region deployments
- Teams with K8s expertise
- Large-scale (> 1TB)

#### 🔧 Quick Start

```bash
# 1. Create EKS cluster
eksctl create cluster --name neo4j-prod --region us-east-1 \
  --node-type m5.xlarge --nodes 3

# 2. Install Neo4j Helm chart
helm repo add neo4j https://helm.neo4j.com/neo4j
helm install my-neo4j neo4j/neo4j \
  --set neo4j.edition=community \
  --set volumes.data.mode=volume \
  --set volumes.data.volume.size=100Gi \
  --namespace neo4j --create-namespace
```

<details>
<summary>📄 Full Kubernetes Manifest (click to expand)</summary>

```yaml
apiVersion: apps/v1
kind: StatefulSet
metadata:
  name: neo4j
  namespace: neo4j
spec:
  serviceName: neo4j
  replicas: 1
  selector:
    matchLabels:
      app: neo4j
  template:
    metadata:
      labels:
        app: neo4j
    spec:
      containers:
      - name: neo4j
        image: neo4j:latest
        ports:
        - containerPort: 7474  # HTTP
          name: http
        - containerPort: 7687  # Bolt
          name: bolt
        env:
        - name: NEO4J_AUTH
          value: neo4j/password
        - name: NEO4J_dbms_memory_heap_initial__size
          value: "2g"
        - name: NEO4J_dbms_memory_heap_max__size
          value: "4g"
        - name: NEO4J_dbms_memory_pagecache_size
          value: "4g"
        resources:
          requests:
            memory: "8Gi"
            cpu: "2"
          limits:
            memory: "16Gi"
            cpu: "4"
        volumeMounts:
        - name: neo4j-data
          mountPath: /data
  volumeClaimTemplates:
  - metadata:
      name: neo4j-data
    spec:
      accessModes: ["ReadWriteOnce"]
      storageClassName: gp3
      resources:
        requests:
          storage: 100Gi
---
apiVersion: v1
kind: Service
metadata:
  name: neo4j
  namespace: neo4j
spec:
  type: LoadBalancer
  selector:
    app: neo4j
  ports:
  - name: http
    port: 7474
    targetPort: 7474
  - name: bolt
    port: 7687
    targetPort: 7687
```
</details>

---

### 4️⃣ ECS (Container)

#### ✅ Pros
- Container orchestration
- AWS integration
- Task definitions

#### ❌ Cons
- **Fargate:** No EBS support
- Complex volume management
- Not ideal for stateful apps
- No native clustering

#### 📋 Best For
- Development/staging only
- Microservices already on ECS
- **Not recommended for production**

#### ⚠️ Important
- Use **ECS on EC2** (not Fargate)
- Mount EBS via Docker volumes
- Consider EC2 instead for production

---

### 5️⃣ Local Docker (Development)

#### 🔧 Setup
```bash
# Simple local deployment
docker run -d \
  --name neo4j \
  -p 7474:7474 \
  -p 7687:7687 \
  -v $HOME/neo4j/data:/data \
  -e NEO4J_AUTH=neo4j/password \
  neo4j:latest

# Browser: http://localhost:7474
# Bolt: bolt://localhost:7687
```

---

## 🎯 Decision Tree

```
Start
  │
  ├─ Need production deployment?
  │   ├─ No  → Local Docker
  │   └─ Yes → Continue
  │
  ├─ DevOps resources available?
  │   ├─ No  → Neo4j AuraDB
  │   └─ Yes → Continue
  │
  ├─ Data size?
  │   ├─ < 100GB   → AuraDB or EC2 m5.xlarge
  │   ├─ 100GB-1TB → EC2 m5.2xlarge+ or AuraDB Business
  │   └─ > 1TB     → EC2 r5.4xlarge+ or AuraDB Enterprise
  │
  ├─ Need HA/clustering?
  │   ├─ Yes → AuraDB or EKS
  │   └─ No  → EC2
  │
  ├─ Budget priority?
  │   ├─ Cost-conscious → EC2 self-managed
  │   ├─ Time-conscious → AuraDB
  │   └─ Balanced       → Start with AuraDB
  │
  └─ Already using Kubernetes?
      ├─ Yes → EKS StatefulSet
      └─ No  → AuraDB or EC2
```

---

## ⚙️ Configuration Quick Reference

### Memory Settings (Critical!)

**Formula:** `heap + page_cache ≤ 90% of total RAM`

| Total RAM | Heap (min) | Heap (max) | Page Cache | Config |
|-----------|------------|------------|------------|--------|
| 8GB | 1g | 2g | 4g | Small |
| 16GB | 2g | 4g | 8g | Medium |
| 32GB | 4g | 8g | 16g | Large |
| 64GB | 8g | 16g | 32g | XL |

### Production Configuration

```bash
# Environment variables for Docker/ECS
NEO4J_dbms_memory_heap_initial__size=4g
NEO4J_dbms_memory_heap_max__size=8g
NEO4J_dbms_memory_pagecache_size=16g

# Enable query logging
NEO4J_dbms_logs_query_enabled=INFO
NEO4J_dbms_logs_query_threshold=1s

# Connection limits
NEO4J_dbms_connector_bolt_thread__pool__max__size=400

# Transaction timeout
NEO4J_dbms_transaction_timeout=60s
```

### neo4j.conf (EC2)

```properties
# Memory
dbms.memory.heap.initial_size=4g
dbms.memory.heap.max_size=8g
dbms.memory.pagecache.size=16g

# Network
dbms.default_listen_address=0.0.0.0
dbms.connector.bolt.listen_address=:7687
dbms.connector.http.listen_address=:7474

# Performance
dbms.checkpoint.interval.time=15m
dbms.checkpoint.iops.limit=1000
```

---

## 🔒 Security Checklist

| Task | Action | Priority |
|------|--------|----------|
| **Change default password** | Set strong NEO4J_AUTH | Critical |
| **Enable HTTPS** | Use SSL certificates for 7474 | High |
| **Private subnet** | Deploy in VPC private subnet | High |
| **Security groups** | Allow only app tier on 7687 | High |
| **Disable HTTP** | Use only Bolt (7687) in prod | Medium |
| **Auth plugin** | LDAP/SSO for enterprise | Medium |
| **Encryption at rest** | Enable EBS encryption | High |
| **Network encryption** | Use neo4j+s:// protocol | High |

### Firewall Rules (Security Group)

| Port | Protocol | Source | Purpose |
|------|----------|--------|---------|
| 7687 | TCP | App tier SG | Bolt (required) |
| 7474 | TCP | VPN/bastion | Browser (optional) |
| 6362 | TCP | Backup server | Backup port (optional) |

---

## 📊 Monitoring Essentials

### Key Metrics

| Metric | Target | Alert If | Tool |
|--------|--------|----------|------|
| **Heap usage** | < 80% | > 90% | JMX/CloudWatch |
| **Page cache hit ratio** | > 90% | < 80% | Neo4j metrics |
| **Query latency (p95)** | < 100ms | > 1s | Query log |
| **Disk IOPS** | < 80% of limit | > 90% | CloudWatch |
| **Store size growth** | Track trend | Unexpected spike | Disk usage |
| **Transaction rate** | Track baseline | Drops 50% | JMX |

### CloudWatch Metrics (EC2)

```bash
# Publish custom metrics
aws cloudwatch put-metric-data \
  --namespace Neo4j \
  --metric-name HeapUsage \
  --value 75 \
  --unit Percent
```

### Prometheus Integration

Neo4j exposes metrics at: `http://localhost:2004/metrics`

```yaml
# prometheus.yml
scrape_configs:
- job_name: neo4j
  static_configs:
  - targets: ['neo4j:2004']
```

---

## ⚡ Performance Tuning

### Storage Performance

| Optimization | Impact | When |
|--------------|--------|------|
| **Use gp3 over gp2** | 20% better IOPS | Always |
| **Provision IOPS (io2)** | 5-10x faster | > 500GB data |
| **Increase page cache** | 2-5x read speed | Read-heavy |
| **Warm up cache** | Faster first queries | After restart |
| **Enable checkpointing** | Stable performance | Production |

### Query Optimization

| Tip | Impact | Example |
|-----|--------|---------|
| **Use indexes** | 10-100x faster | `CREATE INDEX ON :User(email)` |
| **Avoid cartesian products** | Massive | Use WHERE properly |
| **Limit result sets** | Reduce memory | Add `LIMIT 100` |
| **Use PROFILE** | Identify bottlenecks | `PROFILE MATCH ...` |
| **Parameterized queries** | Better caching | Use `$param` syntax |

### Heap Sizing Rules

| Workload | Heap (% of RAM) | Page Cache (% of RAM) |
|----------|-----------------|----------------------|
| **Read-heavy** | 25% | 60% |
| **Write-heavy** | 40% | 45% |
| **Mixed** | 30-35% | 50-55% |
| **Graph analytics** | 20% | 65% |

---

## 🔄 Backup & Recovery

### Backup Strategies

| Method | Frequency | Retention | Cost |
|--------|-----------|-----------|------|
| **EBS snapshots** | Daily | 30 days | $0.05/GB/mo |
| **neo4j-admin dump** | Weekly | 90 days | S3 standard |
| **Continuous backup** | Real-time | 7 days | AuraDB included |
| **Point-in-time** | Transaction log | 24 hours | AuraDB Business+ |

### Backup Commands

```bash
# 1. Stop Neo4j (or use online backup for Enterprise)
docker stop neo4j

# 2. Create backup
neo4j-admin dump \
  --database=neo4j \
  --to=/backup/neo4j-$(date +%Y%m%d).dump

# 3. Upload to S3
aws s3 cp /backup/neo4j-*.dump s3://bucket/backups/

# 4. Restore
neo4j-admin load \
  --from=/backup/neo4j-20250101.dump \
  --database=neo4j \
  --force
```

---

## ❌ What to Avoid

| Don't Do This | Why | Do This Instead |
|---------------|-----|-----------------|
| **ECS Fargate** | No EBS support | EC2 or AuraDB |
| **EFS for data** | Too slow for graph queries | EBS gp3/io2 |
| **t3/t2 instances** | CPU throttling kills performance | m5/r5 series |
| **Shared storage** | Contention issues | Dedicated EBS |
| **No page cache** | 10-100x slower queries | Size properly |
| **Default passwords** | Security risk | Change immediately |
| **No backups** | Data loss risk | Automate backups |
| **Single AZ** | No HA | Multi-AZ or AuraDB |

---

## 📚 Architecture Patterns

### Pattern 1: Startup/MVP
```
Application (ECS/Lambda)
       ↓
Neo4j AuraDB Free/Pro
```
**Cost:** $0-65/mo | **Effort:** Minimal | **Time to Deploy:** 5 min

### Pattern 2: Production (Self-Managed)
```
Application (ECS)
       ↓
  ALB (7687)
       ↓
EC2 m5.xlarge (16GB RAM)
  └─ EBS gp3 (100GB)
  └─ Daily snapshots → S3
```
**Cost:** $150/mo | **Effort:** Medium | **Time to Deploy:** 1 hour

### Pattern 3: High-Availability
```
Application (Multi-AZ)
       ↓
    Route53
       ↓
Neo4j AuraDB Business (Clustered)
  ├─ Primary (us-east-1a)
  ├─ Replica (us-east-1b)
  └─ Replica (us-east-1c)
```
**Cost:** $400+/mo | **Effort:** Minimal | **HA:** 99.95%

### Pattern 4: Enterprise Scale
```
Multi-Region App
       ↓
   Global Accelerator
       ↓
EKS Clusters (Multi-Region)
  ├─ Neo4j Causal Cluster (3 cores)
  ├─ Read Replicas (6 nodes)
  └─ EBS io2 (1TB+)
```
**Cost:** $2,000+/mo | **Effort:** High | **Scale:** Billions of nodes

---

## 🎓 High Availability Setup

### AuraDB (Automatic)
- Built-in clustering
- Automatic failover (< 1 min)
- Multi-AZ deployment
- **No configuration needed**

### EC2 (Manual)
```bash
# Requires Neo4j Enterprise Edition
# 3-node causal cluster (1 leader + 2 followers)

# Core server 1 (Leader candidate)
NEO4J_dbms_mode=CORE
NEO4J_causal__clustering_minimum__core__cluster__size__at__formation=3
NEO4J_causal__clustering_initial__discovery__members=core1:5000,core2:5000,core3:5000

# Read replica
NEO4J_dbms_mode=READ_REPLICA
NEO4J_causal__clustering_initial__discovery__members=core1:5000,core2:5000,core3:5000
```

---

## 🏁 Final Recommendations

### By Scenario

| Scenario | Choice | Reason |
|----------|--------|--------|
| **"Just starting out"** | AuraDB Free | $0, instant |
| **"Need production now"** | AuraDB Professional | $65, managed |
| **"Large dataset (> 1TB)"** | EC2 r5.4xlarge+ | Performance |
| **"Cost is critical"** | EC2 m5.xlarge | $150 self-managed |
| **"Already on K8s"** | EKS StatefulSet | Fits stack |
| **"Need 99.95% uptime"** | AuraDB Business | SLA included |
| **"Dev/testing"** | Local Docker | Free, instant |
| **"Enterprise (billions of nodes)"** | AuraDB Enterprise | Proven scale |

### Quick Decision

- **< 100GB + any team size** → **Neo4j AuraDB**
- **> 100GB + DevOps team** → **EC2 m5/r5**
- **> 1TB + need HA** → **AuraDB Business or EKS**
- **Already on K8s** → **EKS StatefulSet**

---

## 📖 Resources

- [Neo4j Docs](https://neo4j.com/docs/)
- [AuraDB](https://neo4j.com/cloud/aura/)
- [Neo4j Operations Manual](https://neo4j.com/docs/operations-manual/)
- [Performance Tuning](https://neo4j.com/docs/operations-manual/current/performance/)
- [Helm Charts](https://helm.neo4j.com/neo4j)
- [AWS Best Practices](https://neo4j.com/developer/neo4j-cloud-aws/)

---

## 🆚 Neo4j vs Traditional Databases

| Aspect | Neo4j | PostgreSQL | MongoDB |
|--------|-------|------------|---------|
| **Data Model** | Graph | Relational | Document |
| **Best For** | Relationships | Transactions | Flexibility |
| **Query Language** | Cypher | SQL | MQL |
| **Joins** | Native (traversals) | Expensive | Manual |
| **Scaling** | Vertical first | Horizontal (sharding) | Horizontal |
| **Storage** | Disk + page cache | Disk + shared buffers | RAM + disk |
| **AWS Recommendation** | **EC2 or AuraDB** | RDS or EC2 | DocumentDB or EC2 |
| **Instance Type** | m5/r5 | m5/r5 | r5/r6i |

---

## 🔧 Troubleshooting

| Issue | Cause | Solution |
|-------|-------|----------|
| **Slow queries** | No indexes | `CREATE INDEX ON :Label(property)` |
| **Out of memory** | Heap too small | Increase heap + page cache |
| **Connection refused** | Firewall/SG | Open port 7687 |
| **Transaction timeout** | Long-running query | Increase `dbms.transaction.timeout` |
| **Disk full** | No log rotation | Enable `dbms.logs.query.rotation` |
| **High CPU** | Inefficient queries | Use `PROFILE` to optimize |
| **Page cache thrashing** | Cache too small | Increase page cache size |

---

## Bottom Line

**For production Neo4j:**
- **Default choice**: **Neo4j AuraDB** (managed, optimized, HA)
- **Cost-sensitive**: **EC2 + EBS** (self-managed)
- **Large-scale (> 1TB)**: **EC2 r5.4xlarge+** or **AuraDB Enterprise**
- **Already on K8s**: **EKS StatefulSet**

Neo4j performs best on:
- **Direct disk I/O** (EC2 > containers)
- **Fast random reads** (EBS gp3/io2)
- **Proper memory tuning** (heap + page cache)

**Start with AuraDB, migrate to EC2 only if:**
- Cost savings justify ops overhead
- Need > 1TB at lower cost
- Have dedicated DevOps team
