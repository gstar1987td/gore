rule: demo 1.0

condition:
    fact.a > 1
    fact.b <= 0
;

run:
    if fact.a + fact.b >= 0
    then
        var.hit = true
    end
;