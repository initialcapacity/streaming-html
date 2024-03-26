# Google Cloud deployment

1.  Set the GCP project ID as an environment variable.

    ```shell
    export PROJECT_ID={google project id}
    ```

1.  Create a service account for the pipeline.

    ```shell
    gcloud auth login
    gcloud config set project ${PROJECT_ID}
    gcloud auth application-default login
    gcloud services enable \
     iamcredentials.googleapis.com \
     run.googleapis.com \
     cloudbuild.googleapis.com \
     artifactregistry.googleapis.com \
     --project "${PROJECT_ID}"
    gcloud iam service-accounts create github-service-account --project "${PROJECT_ID}"
    ```

1.  Create a workload identity pool.

    ```shell
    gcloud iam workload-identity-pools create github-pool \
      --project="${PROJECT_ID}" \
      --location="global" \
      --display-name=github-pool
    gcloud iam workload-identity-pools describe github-pool \
      --project="${PROJECT_ID}" \
      --location="global" \
      --format="value(name)"
    ```

1.  Set the workload identity pool ID from the output of the last command.

    ```shell
    export WORKLOAD_IDENTITY_POOL_ID={from previous command output}
    ```

1.  Create a workload identity pool provider.

    ```shell
    gcloud iam workload-identity-pools providers create-oidc github-provider \
      --project="${PROJECT_ID}" \
      --location="global" \
      --workload-identity-pool=github-pool \
      --display-name=github-provider \
      --attribute-mapping="google.subject=assertion.sub,attribute.actor=assertion.actor,attribute.repository=assertion.repository" \
      --issuer-uri="https://token.actions.githubusercontent.com"
    gcloud iam service-accounts add-iam-policy-binding "github-service-account@${PROJECT_ID}.iam.gserviceaccount.com" \
      --project="${PROJECT_ID}" \
      --role="roles/iam.workloadIdentityUser" \
      --member="principalSet://iam.googleapis.com/${WORKLOAD_IDENTITY_POOL_ID}/attribute.repository/initialcapacity/streaming-html"
    gcloud iam workload-identity-pools providers describe github-provider \
      --project="${PROJECT_ID}" \
      --location="global" \
      --workload-identity-pool=github-pool \
      --format="value(name)"
    ```

1.  Give api permissions to the service account.

    ```shell
    gcloud projects add-iam-policy-binding $PROJECT_ID --member="serviceAccount:github-service-account@${PROJECT_ID}.iam.gserviceaccount.com" \
        --role="roles/artifactregistry.admin"
    gcloud projects add-iam-policy-binding $PROJECT_ID --member="serviceAccount:github-service-account@${PROJECT_ID}.iam.gserviceaccount.com" \
        --role="roles/run.admin"
    gcloud projects add-iam-policy-binding $PROJECT_ID --member="serviceAccount:github-service-account@${PROJECT_ID}.iam.gserviceaccount.com" \
        --role="roles/viewer"
    gcloud projects add-iam-policy-binding $PROJECT_ID --member="serviceAccount:github-service-account@${PROJECT_ID}.iam.gserviceaccount.com" \
        --role="roles/iam.serviceAccountUser"
    gcloud projects add-iam-policy-binding $PROJECT_ID --member="serviceAccount:github-service-account@${PROJECT_ID}.iam.gserviceaccount.com" \
        --role="roles/cloudbuild.builds.viewer"
    gcloud projects add-iam-policy-binding $PROJECT_ID --member="serviceAccount:github-service-account@${PROJECT_ID}.iam.gserviceaccount.com" \
        --role="roles/cloudbuild.builds.builder"
    gcloud projects get-iam-policy $PROJECT_ID --flatten="bindings[].members" \
        --format='table(bindings.role)' \
        --filter="bindings.members:github-service-account@${PROJECT_ID}.iam.gserviceaccount.com"
    ```