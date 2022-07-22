# API

```
GET /characters/new
```

```
{
    "Name": string,
    "Level": number,
    "Class": string,
    "Species": string,
    "Vocation": string,
    "HitPoints": number,
    "SavingThrow": number,
    "AttackValue": number,
    "Attributes": {
        "str": {
            "score": number,
            "groups": []string,
        },
        "dex": {
            "score": number,
            "groups": []string,
        },
        "con": {
            "score": number,
            "groups": []string,
        },
        "int": {
            "score": number,
            "groups": []string,
        },
        "wis": {
            "score": number,
            "groups": []string,
        },
        "cha": {
            "score": number,
            "groups": []string,
        }
    },
    "Languages": []string,
    "Coins": number
}
```

# To-Do

- Generate equipment
- Generate personality quirks/traits
- Allow choosing starting XP
