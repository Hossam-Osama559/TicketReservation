
let scanner_on=true
let result=document.getElementById("state")
let scanner_controller=document.getElementById("control_scanner")
const html5QrCode = new Html5Qrcode("reader");

function start_scan(){

    html5QrCode.start(
  { facingMode: "environment" },
  {
    fps: 10,
    qrbox: 250
  },
  onScanSuccess
).catch(err => {
  console.error("Camera start failed:", err);
});
}


function stop_scan(){

    html5QrCode.stop().then(()=>console.log("done"))
}


scanner_controller.addEventListener("click",function(e){

       

       if (scanner_on){

          scanner_on=false

          start_scan()

          scanner_controller.textContent="Stop the scanner"
       }else {

          scanner_on=true 

          stop_scan()

          scanner_controller.textContent="Start the scanner"
       }
})


function scan_result(state){

    if (state===1){

            result.textContent="this  ticket is valid"

    }
    else {

            result.textContent="this ticket is not valid"
    }


    setTimeout(()=>{

        result.textContent=""
    },1000)


}



async function onScanSuccess(decodedText, decodedResult) {
//   console.log("Scanned:", decodedText," decodes results",decodedResult);


    let req_body={

        Ticket:decodedText
    }


    let response=await fetch("/tickets/",{

        method:"PUT",

        headers:{

            "Content-Type":"application/json"
        },

        body:JSON.stringify(req_body)
    })


    let body=await response.json()


    console.log(body)

    console.log(req_body)


    if (body.State==="valid"){

        scan_result(1)
    }

    else {

        scan_result(0)
    }
}







