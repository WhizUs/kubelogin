# Getting Started with Keycloak

## Prerequisite

- You have an administrator role of the Keycloak realm.
- You have an administrator role of the Kubernetes cluster.
- You can configure the Kubernetes API server.
- `kubectl` and `kubelogin` are installed.

## 1. Setup Keycloak

Open the Keycloak and create an OIDC client as follows:

- Client ID: `kubernetes`
- Valid Redirect URLs:
    - `http://localhost:8000`
    - `http://localhost:18000` (used if the port 8000 is already in use)
- Issuer URL: `https://keycloak.example.com/auth/realms/YOUR_REALM`

You can associate client roles by adding the following mapper:

- Name: `groups`
- Mapper Type: `User Client Role`
- Client ID: `kubernetes`
- Client Role prefix: `kubernetes:`
- Token Claim Name: `groups`
- Add to ID token: on

For example, if you have the `admin` role of the client, you will get a JWT with the claim `{"groups": ["kubernetes:admin"]}`.

Now test authentication with the Keycloak.

```sh
kubectl oidc-login get-token -v1 \
  --oidc-issuer-url=https://keycloak.example.com/auth/realms/YOUR_REALM \
  --oidc-client-id=kubernetes \
  --oidc-client-secret=YOUR_CLIENT_SECRET
```

You should get claims like:

```
I0827 12:29:03.086476   23722 get_token.go:59] the ID token has the claim: groups=[kubernetes:admin]
I0827 12:29:03.086531   23722 get_token.go:59] the ID token has the claim: aud=kubernetes
I0827 12:29:03.086553   23722 get_token.go:59] the ID token has the claim: iss=https://keycloak.example.com/auth/realms/YOUR_REALM
I0827 12:29:03.086561   23722 get_token.go:59] the ID token has the claim: sub=f08655e2-901f-48e5-8c64-bb9f7784d5df
```

## 2. Setup Kubernetes API server

Configure your Kubernetes API server accepts [OpenID Connect Tokens](https://kubernetes.io/docs/reference/access-authn-authz/authentication/#openid-connect-tokens).

```
--oidc-issuer-url=https://keycloak.example.com/auth/realms/YOUR_REALM
--oidc-client-id=kubernetes
--oidc-groups-claim=groups
```

If you are using [kops](https://github.com/kubernetes/kops), run `kops edit cluster` and add the following spec:

```yaml
spec:
  kubeAPIServer:
    oidcIssuerURL: https://keycloak.example.com/auth/realms/YOUR_REALM
    oidcClientID: kubernetes
    oidcGroupsClaim: groups
```

## 3. Setup Kubernetes cluster

Here assign the `cluster-admin` role to the `kubernetes:admin` group.

```yaml
kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: keycloak-admin-group
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: cluster-admin
subjects:
- kind: Group
  name: kubernetes:admin
```

You can create a custom role and assign it as well.

## 4. Setup kubeconfig

Configure the kubeconfig like:

```yaml
apiVersion: v1
clusters:
- cluster:
    server: https://api.example.com
  name: example.k8s.local
contexts:
- context:
    cluster: example.k8s.local
    user: keycloak
  name: keycloak@example.k8s.local
current-context: keycloak@example.k8s.local
kind: Config
preferences: {}
users:
- name: keycloak
  user:
    exec:
      apiVersion: client.authentication.k8s.io/v1beta1
      command: kubelogin
      args:
      - get-token
      - --oidc-issuer-url=https://keycloak.example.com/auth/realms/YOUR_REALM
      - --oidc-client-id=kubernetes
      - --oidc-client-secret=YOUR_CLIENT_SECRET
```

You can share the kubeconfig to your team members for on-boarding.

## 5. Run kubectl

Make sure you can access the Kubernetes cluster.

```
% kubectl get nodes
Open http://localhost:8000 for authentication
You got a valid token until 2019-05-16 22:03:13 +0900 JST
Updated ~/.kubeconfig
NAME                                    STATUS    ROLES     AGE       VERSION
ip-1-2-3-4.us-west-2.compute.internal   Ready     node      21d       v1.9.6
ip-1-2-3-5.us-west-2.compute.internal   Ready     node      20d       v1.9.6
```
