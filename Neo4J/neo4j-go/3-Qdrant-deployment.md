# Qdrant Deployment on AWS - Quick Reference

## 🎯 TL;DR

| Use Case | Best Option | Cost | Why |
|----------|-------------|------|-----|
| **MVP/POC** | Qdrant Cloud Free | $0 | Zero setup, 1GB free |
| **Small Prod (< 5M vectors)** | Qdrant Cloud | $25-95/mo | Managed, no ops |
| **Large Prod (> 10M vectors)** | EKS StatefulSet | $615+/mo | HA, scaling |
| **Cost-Conscious** | EC2 + EBS | $180/mo | Self-managed |
| **Dev/Testing** | Local Docker | $0 | Simple, fast |

## 📊 Quick Comparison Matrix

| Feature | EC2 | ECS | EKS | Qdrant Cloud |
|---------|-----|-----|-----|--------------|
| **Setup Complexity** | Medium | Medium | High | Easy |
| **HA/Clustering** | Manual | No | Built-in | Built-in |
| **Scaling** | Manual | Limited | Auto | Auto |
| **Ops Overhead** | High | Medium | High | None |
| **Best For** | Cost-sensitive | Microservices | Large-scale | Startups |
| **Production Ready** | ✅ | ⚠️ (ECS on EC2) | ✅ | ✅ |

## 💰 Cost Comparison (us-east-1)

| Option | Monthly Cost | RAM | Use Case |
|--------|--------------|-----|----------|
| Qdrant Cloud Free | **$0** | 1GB | Development, POC |
| Qdrant Cloud 2GB | **$25** | 2GB | Small production |
| Qdrant Cloud 8GB | **$95** | 8GB | Medium production |
| EC2 r6i.large | **$90** | 16GB | Small self-managed |
| EC2 r6i.xlarge + EBS | **$180** | 32GB | Medium self-managed |
| EC2 r6i.2xlarge + EBS | **$360** | 64GB | Large self-managed |
| EKS + 3x r6i.xlarge | **$615** | 96GB (total) | HA production |

## 📏 Memory Sizing Calculator

**Formula:** `RAM = 1.5x × vectors × dimensions × 4 bytes`

| Vectors | Dimensions | RAM Needed | Instance Type | Cost/mo |
|---------|------------|------------|---------------|---------|
| 100K | 768 | ~460MB | r6i.large (16GB) | $90 |
| 1M | 768 | ~4.6GB | r6i.large (16GB) | $90 |
| 5M | 768 | ~23GB | r6i.xlarge (32GB) | $180 |
| 10M | 768 | ~46GB | r6i.2xlarge (64GB) | $360 |
| 50M | 768 | ~230GB | r6i.8xlarge (256GB) | $1,440 |

---

## 🚀 Deployment Options

### 1️⃣ Qdrant Cloud (Managed) 🏆 RECOMMENDED

#### ✅ Pros
- Zero ops overhead
- Auto backups, HA, monitoring
- Free tier (1GB)
- Global deployment

#### ❌ Cons
- Less control
- Data egress costs
- Monthly fees

#### 📋 Best For
- Startups, MVPs
- Limited DevOps team
- RAG/AI applications
- Quick deployment

#### 🔧 Setup
```bash
# 1. Sign up: cloud.qdrant.io
# 2. Create cluster
# 3. Get API key
curl "https://xyz.cloud.qdrant.io:6333/collections" -H "api-key: KEY"
```

---

### 2️⃣ EKS (Kubernetes) ⭐ For Large-Scale

#### ✅ Pros
- Native clustering
- Auto-scaling
- Official Helm chart
- Multi-region support

#### ❌ Cons
- Requires K8s expertise
- Higher cost ($73/mo control plane)
- Complex setup

#### 📋 Best For
- Production > 10M vectors
- HA requirements
- Already using K8s
- Multi-region

#### 🔧 Quick Start
```bash
# 1. Create cluster
eksctl create cluster --name qdrant --region us-east-1 \
  --node-type r6i.xlarge --nodes 3

# 2. Install Qdrant
helm repo add qdrant https://qdrant.github.io/qdrant-helm
helm install qdrant qdrant/qdrant \
  --set replicaCount=3 \
  --set persistence.size=100Gi \
  --set resources.requests.memory=16Gi \
  --namespace qdrant --create-namespace
```

<details>
<summary>📄 Full Kubernetes Manifest (click to expand)</summary>

```yaml
apiVersion: apps/v1
kind: StatefulSet
metadata:
  name: qdrant
spec:
  serviceName: qdrant
  replicas: 3
  template:
    spec:
      containers:
      - name: qdrant
        image: qdrant/qdrant:latest
        ports:
        - containerPort: 6333  # REST
        - containerPort: 6334  # gRPC
        resources:
          requests:
            memory: "16Gi"
            cpu: "2"
          limits:
            memory: "32Gi"
            cpu: "4"
        volumeMounts:
        - name: storage
          mountPath: /qdrant/storage
  volumeClaimTemplates:
  - metadata:
      name: storage
    spec:
      accessModes: ["ReadWriteOnce"]
      storageClassName: gp3
      resources:
        requests:
          storage: 100Gi
```
</details>

---

### 3️⃣ EC2 + EBS (Self-Managed)

#### ✅ Pros
- Full control
- Predictable cost
- Direct EBS access
- Simple setup

#### ❌ Cons
- Manual HA setup
- No auto-scaling
- You manage OS
- Manual failover

#### 📋 Best For
- Cost-sensitive
- < 10M vectors
- Single-region
- Existing EC2 infra

#### 🔧 Setup
```bash
# EC2 User Data
mkfs -t ext4 /dev/xvdf
mount /dev/xvdf /var/lib/qdrant

docker run -d --name qdrant \
  -p 6333:6333 -p 6334:6334 \
  -v /var/lib/qdrant:/qdrant/storage:z \
  qdrant/qdrant:latest
```

#### 💻 Instance Recommendations

| Vectors | Instance | RAM | vCPU | Storage | Cost/mo |
|---------|----------|-----|------|---------|---------|
| < 1M | r6i.large | 16GB | 2 | 50GB gp3 | $90 |
| 1M-10M | r6i.xlarge | 32GB | 4 | 100GB gp3 | $180 |
| 10M-50M | r6i.2xlarge | 64GB | 8 | 200GB gp3 | $360 |
| > 50M | r6i.4xlarge+ | 128GB+ | 16+ | 500GB+ | $720+ |

---

### 4️⃣ ECS (Container)

#### ✅ Pros
- Container orchestration
- AWS integration
- Microservices-friendly

#### ❌ Cons
- **Fargate:** No EBS, max 30GB storage, 120GB RAM
- No native clustering
- Complex persistent volumes

#### 📋 Best For
- Dev/staging only
- Small vectors (< 1M)
- Already using ECS

#### ⚠️ Important
- Use **ECS on EC2** (not Fargate)
- Mount EBS via Docker volumes
- EFS only for backups

---

### 5️⃣ Local Docker (Development)

#### 🔧 Setup
```bash
docker run -d --name qdrant \
  -p 6333:6333 -p 6334:6334 \
  -v $(pwd)/qdrant_storage:/qdrant/storage:z \
  qdrant/qdrant:latest

# Dashboard: http://localhost:6333/dashboard
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
  ├─ DevOps expertise available?
  │   ├─ No  → Qdrant Cloud
  │   └─ Yes → Continue
  │
  ├─ Vector count?
  │   ├─ < 5M   → Qdrant Cloud or EC2
  │   ├─ 5M-10M → EC2 or EKS
  │   └─ > 10M  → EKS
  │
  ├─ Need HA/clustering?
  │   ├─ Yes → EKS or Qdrant Cloud
  │   └─ No  → EC2
  │
  └─ Already using Kubernetes?
      ├─ Yes → EKS
      └─ No  → Qdrant Cloud or EC2
```

---

## ⚙️ Configuration Quick Reference

### Production Qdrant Config
```yaml
storage:
  on_disk_payload: false  # Keep in RAM
  hnsw_config:
    m: 16                 # Connections (↑ = better recall, more RAM)
    ef_construct: 100     # Build quality

service:
  grpc_port: 6334
  max_request_size_mb: 256
```

### Kubernetes Resources
```yaml
resources:
  requests:
    memory: "16Gi"  # Actual usage
    cpu: "2"
  limits:
    memory: "32Gi"  # Allow burst
    cpu: "4"
```

### Clustering (EKS)
```yaml
env:
- name: QDRANT__CLUSTER__ENABLED
  value: "true"
- name: QDRANT__CLUSTER__P2P__PORT
  value: "6335"
```

---

## 🔒 Security Checklist

| Task | Command/Action |
|------|----------------|
| **Enable API Key** | `docker run -e QDRANT__SERVICE__API_KEY=secret` |
| **Private Subnet** | Place in VPC private subnet |
| **Security Groups** | Allow only from app tier |
| **TLS/SSL** | Enable for production |
| **PrivateLink** | Use for Qdrant Cloud |

---

## 📊 Monitoring Essentials

| Metric | Target | Alert If |
|--------|--------|----------|
| **Query Latency (p95)** | < 100ms | > 500ms |
| **Memory Usage** | < 80% | > 90% |
| **Query Throughput** | App-specific | Drops 50% |
| **Search Accuracy** | > 95% recall | < 90% |

### Prometheus Integration
```yaml
annotations:
  prometheus.io/scrape: "true"
  prometheus.io/port: "6333"
  prometheus.io/path: "/metrics"
```

---

## 🔄 Migration Steps

| Step | Command |
|------|---------|
| **1. Export** | `curl -X POST http://localhost:6333/collections/NAME/snapshots` |
| **2. Upload** | `aws s3 cp snapshot.snapshot s3://bucket/` |
| **3. Restore** | `curl -X PUT http://new:6333/collections/NAME/snapshots/upload --data-binary @snapshot.snapshot` |

---

## ⚡ Performance Tips

| Optimization | Impact | When to Use |
|--------------|--------|-------------|
| **Keep vectors in RAM** | 10-100x faster | Always (production) |
| **Use gRPC** | 2-3x throughput | High QPS |
| **Batch inserts** | 5-10x faster indexing | Bulk loads |
| **Clustering** | Linear scaling | > 10M vectors |
| **HNSW tuning** | Better recall/speed | Fine-tuning |

---

## ❌ What to Avoid

| Don't Do This | Why | Do This Instead |
|---------------|-----|-----------------|
| **ECS Fargate** | No EBS, 30GB limit | ECS on EC2 or EKS |
| **EFS storage** | Too slow for vectors | EBS gp3/io2 |
| **t3/t2 instances** | CPU throttling | r6i/r7i (memory-optimized) |
| **No memory limits** | OOM crashes | Set K8s limits |
| **Single node (prod)** | No HA | 3+ replicas |

---

## 📚 Architecture Patterns

### Pattern 1: Startup/MVP
```
Application (ECS/Lambda)
       ↓
Qdrant Cloud (1-2GB)
```
**Cost:** $0-25/mo | **Effort:** Minimal

### Pattern 2: Mid-Sized Production
```
Application (ECS)
       ↓
EC2 r6i.xlarge (32GB RAM)
  └─ EBS gp3 (100GB)
```
**Cost:** $180/mo | **Effort:** Medium

### Pattern 3: Large-Scale HA
```
Application (EKS/ECS)
       ↓
  ALB/NLB
       ↓
EKS Cluster (3 nodes)
  ├─ Qdrant Pod 1 (r6i.xlarge)
  ├─ Qdrant Pod 2 (r6i.xlarge)
  └─ Qdrant Pod 3 (r6i.xlarge)
```
**Cost:** $615/mo | **Effort:** High

---

## 🎓 When to Cluster?

| Trigger | Action |
|---------|--------|
| **> 10M vectors** | Enable clustering for sharding |
| **High QPS (> 1K)** | Add read replicas |
| **HA required** | 3+ replicas across AZs |
| **> 100GB RAM needed** | Shard across nodes |

---

## 🏁 Final Recommendations

### By Use Case

| Scenario | Choice | Reason |
|----------|--------|--------|
| **"Just starting out"** | Qdrant Cloud Free | $0, zero setup |
| **"Production RAG app"** | Qdrant Cloud 2-8GB | $25-95, managed |
| **"Need HA, large-scale"** | EKS StatefulSet | Clustering, auto-scale |
| **"Cost is critical"** | EC2 r6i.xlarge | $180, self-managed |
| **"Already on K8s"** | EKS StatefulSet | Fits existing stack |
| **"Dev/testing"** | Local Docker | Free, instant |

### Quick Decision
- **< 5M vectors + limited ops** → **Qdrant Cloud**
- **> 10M vectors + have K8s team** → **EKS**
- **In between + cost-conscious** → **EC2**

---

## 📖 Resources

- [Qdrant Docs](https://qdrant.tech/documentation/)
- [Helm Chart](https://github.com/qdrant/qdrant-helm)
- [Qdrant Cloud](https://cloud.qdrant.io/)
- [Performance Guide](https://qdrant.tech/documentation/guides/optimization/)
- [AWS EKS Best Practices](https://aws.github.io/aws-eks-best-practices/)

---

## 🆚 Key Differences from Neo4j

| Aspect | Qdrant | Neo4j |
|--------|--------|-------|
| **Type** | Vector DB | Graph DB |
| **Storage** | RAM-first | Disk-first |
| **Instance** | r6i/r7i (memory) | m5/c5 (balanced) |
| **Scaling** | Horizontal (sharding) | Vertical (larger nodes) |
| **K8s Fit** | Excellent (cloud-native) | Good (but prefers EC2) |
| **Best Deploy** | **EKS or Cloud** | **EC2 or AuraDB** |
