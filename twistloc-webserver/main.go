package main

import (

    "log"

    "net/http"

)



func triggerImageRebuild(w http.ResponseWriter, r *http.Request) {



    buildNames, ok := r.URL.Query()["buildName"]



    ua := r.Header.Get("User-Agent")



    if ua != "tekton" {

        log.Println("invalid requester")

        return

    }



    if !ok || len(buildNames[0]) < 1 {

        log.Println("Url Param 'buildName' is missing")

        return

    }



    // Query()["buildName"] will return an array of items,

    // we only want the single item.

    buildName := buildNames[0]



    log.Println("Url Param 'buildName' is: " + string(buildName))

}



func main() {



    http.HandleFunc("/rebuild-image", triggerImageRebuild)



    http.ListenAndServe(":8080", nil)

}