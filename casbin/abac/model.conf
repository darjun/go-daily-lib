[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[matchers]
m = r.sub.Hour >= 9 && r.sub.Hour < 18 || r.sub.Name == r.obj.Owner

[policy_effect]
e = some(where (p.eft == allow))