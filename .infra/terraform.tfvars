#--------------------------------------------------------------
# General
#--------------------------------------------------------------
app_name          = "brdev2024002"
location          = "brazilsouth"
domain_name_label = ""

#--------------------------------------------------------------
# AKS configuration params
#--------------------------------------------------------------
kubernetes_version  = "1.30.3"
vm_size_node_pool   = "Standard_D4ds_v5"
node_pool_min_count = "1"
node_pool_max_count = "10"

#--------------------------------------------------------------
# Helm Chart versions
#--------------------------------------------------------------
helm_pod_identity_version = ""
helm_csi_secrets_version  = ""
helm_agic_version         = ""
helm_keda_version         = ""