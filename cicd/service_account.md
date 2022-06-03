One possiblity that we recommend is to use [Service Account Tokens](https://kubernetes.io/docs/reference/access-authn-authz/authentication/#service-account-tokens)

##create kubeconfigs for cicd

### create a service account
`kubectl create serviceaccount cicd-sa`

### get the token for the service account
`kubectl get secret cicd-sa-* -o jsonpath="{.data.token}" | base64 -d `

change the **** with the correct value from the secret

create a kubeconfig with the secret from the token

```yaml
apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUM1ekNDQWMrZ0F3SUJBZ0lVY1R2ZGlIRStKR1UyNnBhL3dVa084WTB4eW1Nd0RRWUpLb1pJaHZjTkFRRUwKQlFBd0RURUxNQWtHQTFVRUF3d0NZMkV3SGhjTk1qRXdNVEEyTVRNd056VTBXaGNOTXpFd01UQTBNVE13TnpVMApXakFOTVFzd0NRWURWUVFEREFKallUQ0NBU0l3RFFZSktvWklodmNOQVFFQkJRQURnZ0VQQURDQ0FRb0NnZ0VCCkFMNzA4R2RpdzVWdndBbWN4UHNLbHVvRGltYUw0MmhCMlRiY2QweVhwRVUvZTBRTURJZVFBY2ZvV2k3ODNDTU0KRVRrdWc4TXdMNTAvamM0Z1c2M25BcFJUWkxvZDBYczBjK0t6RWltUTFBdGYyNElKQTNZdWh5K0ZUVzBLbVAraApYUm9Fd2VreWZNODJxRUpxSjJQZUl4djlmOUxzVndZRTVxSGlFcDYxZFZhaFE2Nzd3ZUh6SzFnSUJXaVpxbWpKClJTcHhxRmExOFdRRTFJQ1c4NGpjTGUwVURhQVlHV3M5OUJpdG1Xbkx3UFpBNHAyakMyRk5PMU0wUTNxYnJzcUgKdFp1a1JUNC9QZG9ueDNxbnp4Z2JzdDJsVitzSHJ4N3JVYnI0ZzlnZzRWMkhnY2ZLWHVMb0Z4K2ZLa0Zrc1FJNQpzR1NzREk2ZTFaSGRTMFJZbG15eHhWY0NBd0VBQWFNL01EMHdDd1lEVlIwUEJBUURBZ0VHTUE4R0ExVWRFd0VCCi93UUZNQU1CQWY4d0hRWURWUjBPQkJZRUZGUGxMVDZISlk1RG9qRVdEN1FlRjBsV2UyM2ZNQTBHQ1NxR1NJYjMKRFFFQkN3VUFBNElCQVFBLzQzZ0pyTmRndzJpSDl3T2h4eTdxUWtnQTFaMDc5TFJ4enlKRS9HSHYxbm9uWmtNZgpkZUlseTRBcDZhbVNxaVBWVDI5dFdwSnZOaXZrMjc3TTVyUTk0ZE1scnA5dFFZTlJkbWIwOWloRjZDOWdjTmxqCkJlRW92OUFWYmV0TEFLTkE3amx0NCsxOTZQRkIxSUxReUdLWllDeTNBRmt0bjgyWkNkWUNWOTMrYkNQcDNnVEQKa1RsWDlpVG9TRVFoRkpmNGhIT0doVElsUnJ0bkxZN1Z6TllLTU5NU0NhVlZVQW5qUndyR0VGYmpvZzFhWVJmUgo1WmhhUXJhQjQ2SnYzSXlQVEhFa2NWV0thRW5obXBXbjFRSTZWTFdLYms2MXNRZlpySlFTRGZpNzJzUFhkUFdaCkJCcm5XQS84NDZKdlRVYytaaW5HTDhNQjU5aVBuT1p5QWtlMwotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg==
    server: https://apiserver.agile-dingo.cns.eu1.cloudmobility.io
  name: agile-dingo
contexts:
- context:
    cluster: agile-dingo
    namespace: <NAMESPACE-ID-HERE>
    user: cicd-sa
  name: agile-dingo-<NAMESPACE-ID-HERE>
current-context: agile-dingo-<NAMESPACE-ID-HERE>
kind: Config
preferences: {}
users:
- name: cicd-sa
  user:
    token: <TOKEN-HERE>
```

### edit the rolebinding
` kubectl edit rolebindings <namespace-id>-kns-edit`

```diff
+ - kind: ServiceAccount
+  name: cicd-sa
+  namespace: <namespace-id>
```
