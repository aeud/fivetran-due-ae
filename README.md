# Diduenjoy Fivetran Cloud Function

## Endpoint

```
https://fivetran-due-rps3r5yvgq-ew.a.run.app
```

Do not forget to ask LVMH to grant access to your Fivetran Agent Service account.

_For LVMH team_
```
gcloud run services add-iam-policy-binding CLOUD_RUN_SERVICE_NAME \
    --region 'europe-west1' \
    --member 'serviceAccount:TO_CHANGE_HERE@fivetran-production.iam.gserviceaccount.com' \
    --role 'roles/run.invoker' \
    --project host-project
```

## Secret

```
{"api_key": "KEY_HERE", "entities": ["answer_sets", "surveys"]}
```

### List of available entities

- `clients`
- `answer_sets`
- `surveys`
- ~~`feedbacks`~~
