/*

== Setup ==
export GOPATH=$(pwd)
go get github.com/mattn/go-sqlite3

# if you have trouble with libsqlite3 and go build (gccgo) use this build:
go build -gccgoflags -pthread -x cs.go


opencellid.org
    - Website
    - search in csv with LAC and CellID
    - display on Website
    - API Key : <API KEY    

Search Location
    CellLocation
        - MCC:  Mobile Country Code
        - MNC:  Mobile Network Code
        - LAC:  Location Area Code
        - cell id: Cell ID

CSV Structure:
== measurements_* CSV file format ==

mmc:int         Mobile Country Code
net:int         MNC (GSM,UMTS,LTE), SID (CDMA)
area:int        Location Area Code (GSM, UMTS), Tracking Area Code, TAC (LTE), NID (CDMA)
cell:int        CID (GSM, LTE), UTRAN (UMTS)
lon:double      longitude
lat:double      latitude
signal:int      dBm or TS, signal level
measured:int    date in timestamp format
created:int     date in timestamp format
rating:double   gps quality in meters
speed:double
direction:double
radio:string
ta:int          Timing advance, GSM and LTE
rnc:int         Radio Network Controller (UMTS)
cid:int         Cell ID short (UMTS)
psc:int         Primary scambling code (UMTS)
tac:int         Tracking area code (LTE)
pci:int         Physical cell id (LTE)
sid:int         System Identifier (CDMA)
nid:int         Network id (CDMA)
bid:int         Base station id (CDMA)


cell_towers.csv

radio:string    Network type. One of the strings GSM, UMTS, LTE or CDMA.
mcc:int         Mobile Country Code, for example 260 for Poland
net:int         For GSM, UMTS and LTE networks, this is the Mobile Network Code (MNC).
                For CDMA networks, this is the System IDentification number (SID).
area:int        Location Area Code (LAC) for GSM and UMTS networks.
                Tracking Area Code (TAC) for LTE networks.
                Network IDenfitication number (NID) for CDMA networks. ocation Area Code (LAC) for GSM and UMTS networks.
                Tracking Area Code (TAC) for LTE networks.
                Network IDenfitication number (NID) for CDMA networks.
cell:int        Cell ID (CID) for GSM and LTE networks.
                UTRAN Cell ID / LCID for UMTS networks, which is the concatenation of 2 or 4 bytes of Radio Network Controller (RNC) code and 4 bytes of Cell ID.
                Base station IDentifier number (BID) for CDMA networks.
unit:int        Primary Scrambling Code (PSC) for UMTS networks.
                Physical Cell ID (PCI) for LTE networks. An empty value for GSM and CDMA networks.
lon:double
lat:double
range:int       Estimate of cell range, in meters.
samples:int     Total number of measurements assigned to the cell tower
changeable:int  Defines if coordinates of the cell tower are exact or approximate.
                changeable=1: the GPS position of the cell tower has been calculated from all available measurements
                changeable=0: the GPS position of the cell tower is precise - no measurements have been used to calculate it.
created:int
averageSignal:int   verage signal strength from all assigned measurements for the cell.
                    Either in dBm or as defined in TS 27.007 8.5 - both is accepted. verage signal strength from all assigned measurements for the cell.
                    Either in dBm or as defined in TS 27.007 8.5 - both is accepted.



== SQLITE ==

**create the table with SQLITE**
sqlite cell_towers.db
create table cell_towers(radio string, mcc int, net int, area int, cell int, uint int, lon double, lat double, range int, samples int, changeable int, created int, updated int, averageSignal int);
.separator ","
.import ../csv/cell_towers.csv cell_tower
(size 2.8GB)


Example Data to Query:
    MCC:  262
    AREA: 801    #LAC
    cell: 86355

radio|mcc|net|area|cell |unit|lon     |lat      |range|samples|changeable|created   |updated   |averageSignal
UMTS |262|2  |801 |86355|    |13.28527|52.521711|37   |7      |1         |1282569574|1300175362|-91

*/

package main
import (
    "fmt"
    "net/http"
    "html/template"
    "strconv"
    "log"
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

type Page struct {
    Title string
    Body  []byte
    Result *[]Item
    CellID int64
    LacID int64
    Search bool
    Err string
}

type Item struct{
    Lat,Lon float64
    Mcc,Area,Cell int    
}

func handler(w http.ResponseWriter, r *http.Request){
    var p = &Page{Title: "CellSearch",Err:""}
    p.CellID = 0
    p.LacID = 0
    p.Search = false   
    
    if r.Method == "POST" {
        a :=r.FormValue("cellid")
        b :=r.FormValue("lacid")
        if a != "" || b != "" {            
            p.Search = true          
            p.CellID, _ = strconv.ParseInt(a, 10, 32)           
            p.LacID, _ = strconv.ParseInt(b, 10, 32)           
            if p.CellID == 0 || p.LacID ==0 { 
                p.Err = "Not enough arguments or invalid numbers. Please use CellId AND LAC Number"
            }else{            
                // search in db, if you want to have more matches, remove LIMIT
                stmt, err := db.Prepare("SELECT mcc,area,cell,lat,lon FROM cell_towers WHERE cell=? AND area=? LIMIT 1")
                if err != nil { log.Fatal(err) }
                defer stmt.Close()
                rows,err := stmt.Query(p.CellID,p.LacID)

                if err != nil { log.Fatal(err) }
                defer rows.Close()       
                
                var result []Item
                for rows.Next(){        
                    item :=Item{}
                    err = rows.Scan(&item.Mcc,&item.Area,&item.Cell,&item.Lat, &item.Lon)
                    if err != nil { log.Fatal(err) }
                    result = append(result,item)            
                    
                }
                if result != nil{ 
                    p.Result = &result
                }                       
            }            
        }
    }
    t,_ := template.ParseFiles("tmpl/base.tpl","tmpl/content.tpl","tmpl/result.tpl")
    t.ExecuteTemplate(w, "base", p)
}

var db *sql.DB
func db_connect(fn string){
    var err error
    db, err = sql.Open("sqlite3", fn)
    if err != nil {
        log.Fatal(err)
    }    
}

func main(){
    db_connect("./db/csgo.db")    
    if err := db.Ping(); err != nil { panic(err) }
    fmt.Println("Starting local CSGO Webserver on localhost:8080")
    http.HandleFunc("/", handler)    
    http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("html/css"))))
    http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("html/js")))) 
    http.Handle("/fonts/", http.StripPrefix("/fonts/", http.FileServer(http.Dir("html/fonts"))))    
    http.Handle("/imgs/", http.StripPrefix("/imgs/", http.FileServer(http.Dir("html/imgs"))))   
    http.ListenAndServe(":8080",nil)
    defer db.Close()
}
