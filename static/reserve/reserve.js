let reserve_form=document.getElementById("reserve")
let my_tickets=document.getElementById("my_tickets")


reserve_form.addEventListener("submit",async function (e) {


    e.preventDefault()



    //sending a api to make a new session 


    // the respone will be a json response has the price and the url of the checkout page 



    let form_data=new FormData(this)

    let data=Object.fromEntries(form_data.entries())


    let response=await fetch("/tickets/",{

        method:"POST",

        headers:{

            "Content-Type":"application/json"
        },

        body:JSON.stringify(data)
    })


    if (response.status===200){


        let res_body=await response.json()


        // redirect to the checkout page
        
        
        console.log("the price of this session is ",res_body.Price)

        window.location.href=res_body.Url 


    }else {

        console.log("failed to make the session")
    }

    
    
})

function create_new_img(url){

    let img=document.createElement("img")

    img.classList.add("qrcode")

    img.src=url 

    let div=document.createElement("div")

    let a=document.createElement("a")

    a.href=url 

    a.download="ticket.png"

    let p=document.createElement("p")

    p.textContent="From x To y"

    // a.click()

    a.textContent="Download The Ticket"    
    div.appendChild(img)

    div.appendChild(p)
    div.appendChild(a)


    


    

    return div
}


function register_img(img){

    my_tickets.appendChild(img)
}

function new_ticket(ticket){


    let url=`./static/tickets/${ticket.Id}` 

    let new_img=create_new_img(url)


    register_img(new_img)

}




// this will return an array of the srcs of the qr code images 
async function fetch_my_tickets(){


    let respone=await fetch("/tickets")



    let urls=await respone.json()


    console.log(urls)


    let tickets=urls.Tickets


    tickets.forEach(ticket => {

        console.log(ticket)


        new_ticket(ticket)
        
    });



}


fetch_my_tickets()