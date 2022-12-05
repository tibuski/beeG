
# beeG : BEE Client written in Go

![](assets/beeg.jpg)

## Installation
You can either `git clone` the reposioty or directly download binaries for Windows or Linux from Gitea + example .ini file

* **Git** :  
  * Clone repo : `git clone http://netbox.net-mgmt.net:3000/bricl/beeG.git`
  * Copy `beeG.ini` to `My_Service.ini` to ensure it won't get overwritten if you `git pull` later.


* **Files** :  
  * [Windows](http://netbox.net-mgmt.net:3000/bricl/beeG/raw/branch/main/bin/beeG.exe)
  * [Linux](http://netbox.net-mgmt.net:3000/bricl/beeG/raw/branch/main/bin/beeG)
  * [.ini file](http://netbox.net-mgmt.net:3000/bricl/beeG/raw/branch/main/bin/beeG.ini)

## Usage
### Commands
* **Windows** :   
`c:\> ./bin/beeG.exe [.ini file] ["Status"] ["Value"]`

* **Linux** :   
$ `./bin/beeG [.ini file] ["Status"] ["Value"]`

### Options

  - `[ini file]` : **Mandatory** - JSON file containing parameters (eplained below), extension must be .ini
  - `["Status"]` : **Mandatory** - Bee Status. Must be either : "OK", "WA" (*Warning*), "CR" (*Critical*), "TO" (*TimeOut*) or "DS" (*Never Used*).
  - `["Value"]` : **Optional** - *Output* column content in Bee portal. If not present current date/time will be used.
<br><br>

## INI file

```json
{

    "BEE_URL"             : "https://toolbox.net-mgmt.net/BEE-update/bee-update.php",
    "AUTHKEY"             : "35165584761285",
    "REVISION"            : "1",
    "HEARTBEAT"           : "1560",
    "ZOMBEAT"             : "168",
    "TEAMID"              : "CVVS",
    "CMDB_SERVICE"        : "SIPT",
    "X_ORG_ID"            : "MY_SERVICE_NAME",
    "BATCH_DOC"           : "My Service daily backup and sync at 18:00",
    "BATCH_URL"           : "http://my_service/",
    "ALERTMAIL"           : "",
    "ALERTID1"            : "",
    "ALERTID2"            : "",
    "ALERTTEXT"           : "",
    "ALERTENV"            : "",
    "BATCH_COMMENT"       : "",
    "BEHAVIOR"            : ""
}
```

* `"BEE_URL"` : Bee server url
* `"AUTHKEY"` : You can chose any number but it needs to stay the same for next updates.
* `"REVISION"` : Should change at each contract (in fact if it is the same, all firelds are not rewritten in the DB)
* `"HEARTBEAT"` : Minutes before the entry goes in TO state.
* `"ZOMBEAT"`: Hours after an entry in TimeOut will disappear (generates a last email)
* `"TEAMID"` : *CVVS* or *BCNS*
* `"CMDB_SERVICE"` : Application code.
* `"X_ORG_ID"`: Master Key. Service name in Bee. Needs to be in uppercase and with underscores for spaces.
* `"BATCH_DOC"` : Comments for the *inlineDoc* column in Bee.
* `"BATCH_URL"` : URL to reach service (not used in Bee Portal yet).
* `"ALERTID1"` : Not in use yet.
* `"ALERTID2"` : Not in use yet.
* `"ALERTMAIL"` : Not in use yet.
* `"ALERTTEXT"` : Not in use yet.
* `"ALERTENV"` : Not in use yet.
* `"BATCH_COMMENT"` : Not in use yet.
* `"BEHAVIOR"` : Not in use yet.
