# health-checker
This is a lite website health checker tool written in go

## Usage -

<b>Open help menu</b>

```bash
go run . help
```

<b>Run health-check on google.com for port 80</b>

```bash
go run . -d google.com [by default port 80 will be checked]
```

<b>Run health-check on google.com for port 443</b>

```bash
go run . -d google.com -p 443
```

<br>

> [!NOTE]  
> Highlights information that users should take into account, even when skimming.

> [!TIP]
> Optional information to help a user be more successful.

> [!IMPORTANT]  
> Crucial information necessary for users to succeed.

> [!WARNING]  
> Critical content demanding immediate user attention due to potential risks.

> [!CAUTION]
> Negative potential consequences of an action.